package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	// A channel with buffer of 10 
	ch := make(chan int, 10)

	rand.Seed(time.Now().UnixNano())

	for j := 0; j < 10; j++ {
		// 2 wait group are created
		wg.Add(2)
		go func(ch <-chan int) {
			i := <-ch
			fmt.Printf("Receiving %2d\n", i)
			wg.Done()
		}(ch)
		go func(ch chan<- int) {
			i := rand.Intn(100)
			fmt.Printf("Sending   %2d\n", i)
			ch <- i
			wg.Done()
		}(ch)
	}
	wg.Wait()

}
