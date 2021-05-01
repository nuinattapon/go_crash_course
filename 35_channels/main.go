package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	n := 75

	for i := 0; i < n; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < n; i++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
		// delay := rand.Intn(40)
		// time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	close(jobs)
}

// Fast fib
func fib(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// Slow fib
// func fib(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	return fib(n-1) + fib(n-2)
// }
