package main

import "fmt"

func main() {

	fmt.Println("Looping through string:")
	s := "Hello, 世界, สวัสดี"
	for k, v := range s {
		fmt.Println(k, string(v))
	}

	fmt.Println("\nLooping through []rune:")
	r := []rune(s)
	for k, v := range r {
		fmt.Println(k, string(v))
	}

	fmt.Println("\nLooping through []byte:")
	b := []byte(s)
	for k, v := range b {
		fmt.Println(k, string(v))
	}

}
