package main

import "fmt"

func main() {
	a := 5
	// Using var
	var b *int = &a

	// Shorthand
	// b := &a

	fmt.Printf("Value of'a' is %d and value of 'b' is 0x%X\n", a, b)
	fmt.Printf("Type of 'a' is %T\n", a)

	fmt.Printf("Type of 'b' is %T\n", b)

	//  Use * to read val from address
	fmt.Printf("Value of '*b' is %d\n", *b)
	fmt.Printf("Value of '*&a' is %d\n", *&a)

	// Change val with pointer
	*b = 10
	fmt.Printf("\nBecause b := &a\nValue of 'a' after '*b = 10' is %d\n", a)
}
