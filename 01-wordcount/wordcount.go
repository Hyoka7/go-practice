package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type MyMap map[string]int

type SortKey struct {
	word  string
	count int
}

func (M MyMap) counter(s string) {
	splitted := strings.Fields(s)
	for _, v := range splitted {
		M[v] += 1
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := make(MyMap)
	for scanner.Scan() {
		m.counter(scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	sortslice := []SortKey{}
	for word, count := range m {
		sortslice = append(sortslice, SortKey{word, count})
	}
	sort.Slice(sortslice, func(i, j int) bool {
		if sortslice[i].count != sortslice[j].count {
			return sortslice[i].count > sortslice[j].count
		}
		return sortslice[i].word < sortslice[j].word
	})
	for _, item := range sortslice {
		fmt.Printf("%s: %d\n", item.word, item.count)
	}
}
