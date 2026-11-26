package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func wait(s chan string) {
	waiting := rand.IntN(2000) * int(time.Millisecond)
	time.Sleep(time.Duration(waiting))
	fmt.Println(time.Duration(waiting))
	s <- "job completed"
}

func main() {
	for range 5 {
		ch := make(chan string, 1)
		go wait(ch)
		select {
		case res := <-ch:
			fmt.Println(res)

		case <-time.After(1000 * time.Millisecond):
			fmt.Println("Timeout")
		}
	}
}
