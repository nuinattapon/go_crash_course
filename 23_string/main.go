package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "Hello, 世界, สวัสดี"

	fmt.Printf("len(\"%s\") = %d\n", s, len(s))
	fmt.Printf("utf8.RuneCountInString(\"%s\") = %d\n", s, utf8.RuneCountInString(s))

	fmt.Println("\nLooping through s:")
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
