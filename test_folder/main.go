package main

import (
	"fmt"
)

type Message struct {
	id      int
	message string
}

func main() {
	fmt.Printf("Hello, playground\n")
	a := Message{1, "Hi"}

	fmt.Printf("%d %s", a.id, a.message)
}
