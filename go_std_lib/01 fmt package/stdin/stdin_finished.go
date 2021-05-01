package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What is your name: ")

	if s, err := reader.ReadString('\n'); err == nil {
		fmt.Printf("Hi, %s!\n", strings.TrimSpace(s))
	} else {
		fmt.Println(err)
	}
}
