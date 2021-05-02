package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	rand.Seed(time.Now().UnixNano())

	// Set up a done channel that's shared by the whole pipeline,
	// and close that channel when this pipeline exits, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{}, 4)
	defer close(done)

	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute work across multiple functions (ten goroutines) that all read from in.
	c0 := worker(done, in)
	c1 := worker(done, in)
	c2 := worker(done, in)
	c3 := worker(done, in)

	// Tell the remaining senders we're leaving.
	// go func(n int) {
	// 	time.Sleep(time.Duration(n) * time.Millisecond)
	// 	done <- struct{}{}
	// 	done <- struct{}{}
	// 	done <- struct{}{}
	// 	done <- struct{}{}
	// }(20)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel
	var y int
	for n := range merge(done, c0, c1, c2, c3) {
		y++
		fmt.Println(y, "\t", n)
	}
	elapsed := time.Since(start)
	fmt.Printf("Elapsed %d ms\n", elapsed.Milliseconds())

}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 500; i++ {
			for j := 0; j < 50; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func worker(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- fib(n):
			case <-done:
				return
			}
		}
	}()
	return out
}

func fact(n int) int {
	// Add latency
	// time.Sleep(time.Duration(20) * time.Millisecond)
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

// Loop fib
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

// Recursive fib
func fib2(n int) int {
	if n <= 1 {
		return n
	}
	return fib2(n-1) + fib2(n-2)
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()

		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
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
