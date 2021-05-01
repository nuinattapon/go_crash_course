package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute work across multiple functions (ten goroutines) that all read from in.
	c0 := fibonacci(in)
	c1 := fibonacci(in)
	c2 := fibonacci(in)
	c3 := fibonacci(in)
	c4 := fibonacci(in)
	c5 := fibonacci(in)
	c6 := fibonacci(in)
	c7 := fibonacci(in)
	c8 := fibonacci(in)
	c9 := fibonacci(in)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel
	var y int
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		y++
		fmt.Println(y, "\t", n)
	}
	elapsed := time.Since(start)
	fmt.Printf("Elapsed %d ms\n", elapsed.Milliseconds())

}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 15; i++ {
			for j := 1; j < 76; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fibonacci(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fib(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
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

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
