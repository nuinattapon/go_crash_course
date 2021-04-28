package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("temp.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
	fmt.Printf("File content length: %d", len(content))

}
