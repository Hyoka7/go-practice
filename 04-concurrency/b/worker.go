package main

import (
	"fmt"
	"time"
)

func worker(job, res chan int, wn int) {
	for base := range job {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("worker", wn, "is working")
		res <- base * base
	}
}

func main() {
	start := time.Now()
	jobs := make(chan int)
	results := make(chan int)
	for i := 1; i < 6; i += 1 {
		go worker(jobs, results, i)
	}
	res := 0
	go func() {
		for i := 1; i < 21; i += 1 {
			jobs <- i
		}
		close(jobs)
	}()
	for i := 1; i < 21; i += 1 {
		res += <-results
	}
	fmt.Println(res)
	since := time.Since(start)
	fmt.Println("Elapsed", since)
}
