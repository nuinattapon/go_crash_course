package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	fmt.Printf("Threads: %+v\n",runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()

	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
	m.RUnlock()
}

func increment() {
	counter++
	wg.Done()
	m.Unlock()
}
