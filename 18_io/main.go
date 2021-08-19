package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// content, err := ioutil.ReadFile("temp.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("File contents: %s", content)
	// fmt.Printf("File content length: %d", len(content))

	if content, err := ioutil.ReadFile("temp.txt"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("File contents: %s", content)
		fmt.Printf("File content length: %d", len(content))
	}

}
