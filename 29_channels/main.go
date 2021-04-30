package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	rand.Seed(time.Now().UnixNano())
	// A channel and 2 wait group are created
	ch := make(chan int, 10)
	wg.Add(2)
	// Receiving goroutine
	go func(ch <-chan int) {
		// * r in for range chan loop is not the index like normal for loop
		// for r := range ch {
		// 	fmt.Printf("Receiving %2d\n", r)
		// }

		// Here is an alternative to for ranch channel loop
		for {
			if r, ok := <-ch; ok {
				fmt.Printf("Receiving %2d\n", r)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	// Sending goroutine
	go func(ch chan<- int) {
		defer close(ch) // Close the channel

		for j := 0; j < 10; j++ {
			r := rand.Intn(100)
			fmt.Printf("Sending   %2d\n", r)
			ch <- r
		}
		wg.Done()
	}(ch)
	wg.Wait()

}
