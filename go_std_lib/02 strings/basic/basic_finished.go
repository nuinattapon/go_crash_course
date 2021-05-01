package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog"
	// s := "Hello, 世界, สวัสดี"
	// Basic string operations

	// Length of string
	fmt.Printf("len(\"%s\") = %d\n", s, len(s))
	fmt.Printf("utf8.RuneCountInString(\"%s\") = %d\n", s, utf8.RuneCountInString(s))
	// iterate over each character
	for _, ch := range s {
		fmt.Print(string(ch), ",")
	}
	fmt.Println()

	// Using operators < > == !=
	fmt.Println("dog" < "cat")
	fmt.Println("dog" < "horse")
	fmt.Println("dog" == "Dog")

	// Comparing two strings
	result := strings.Compare("dog", "cat")
	fmt.Println(result)
	result = strings.Compare("dog", "dog")
	fmt.Println(result)

	// EqualFold tests using Unicode case-folding
	fmt.Println(strings.EqualFold("Hello", "hello"))

	// ToUpper, ToLower, Title
	s1 := strings.ToUpper(s)
	s2 := strings.ToLower(s)
	s3 := strings.Title(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
