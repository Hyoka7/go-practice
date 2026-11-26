package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Memos []string

func (s *Memos) add(memo string) error {
	if memo == "" {
		return errors.New("invalid syntax")
	}
	*s = append(*s, memo)
	return nil
}

func (s *Memos) list() {
	for i, m := range *s {
		fmt.Println(i+1, m)
	}
}

func (s *Memos) del(index string) error {
	if index == "" {
		return errors.New("invalid syntax")
	}
	atoied, err := strconv.Atoi(index)
	if err != nil {
		return fmt.Errorf("del needs positive integer value: %w", err)
	}
	memolen := len(*s)
	if atoied <= 0 || memolen < atoied {
		return errors.New("index out of range")
	}
	*s = slices.Delete(*s, atoied-1, atoied)
	return nil
}

func (s *Memos) save() error {
	jsoned, err := json.Marshal(*s)
	if err != nil {
		return fmt.Errorf("error while saving: %w", err)
	}
	errWrite := os.WriteFile("memos.json", jsoned, 0644)
	if errWrite != nil {
		return fmt.Errorf("error while writing: %w", errWrite)
	}
	return nil
}

func (s *Memos) load() error {
	readed, err := os.ReadFile("memos.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("error while reading: %w", err)
	}
	err = json.Unmarshal(readed, s)
	if err != nil {
		return fmt.Errorf("error while reading: %w", err)
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var memos Memos
	err := memos.load()
	if err != nil {
		log.Fatal(err)
	}
Loop:
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		before, after, _ := strings.Cut(text, " ")
		switch before {
		case "":
			continue Loop
		case "add":
			err := memos.add(after)
			if err != nil {
				fmt.Println(err)
			}
		case "list":
			memos.list()
		case "del":
			err := memos.del(after)
			if err != nil {
				fmt.Println(err)
			}
		case "quit":
			break Loop
		default:
			fmt.Println("Unknown Command:", before)
		}
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	err = memos.save()
	if err != nil {
		log.Fatal(err)
	}
}
