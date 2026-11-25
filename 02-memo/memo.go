package main

import (
	"bufio"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var memos Memos
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
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
