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
	fmt.Printf("Threads: %+v, NumCPU: %+v\n",
		runtime.GOMAXPROCS(-1), runtime.NumCPU())
	// UPDATE 8/28/2015: Go 1.5 is set to make the default value of GOMAXPROCS
	// the same as the number of CPUs on your machine, so this shouldn't be a
	// problem anymore

	// runtime.GOMAXPROCS(runtime.NumCPU())
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
