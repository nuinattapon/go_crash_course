package main

import (
	"fmt"
)

func main() {
	// Arrays
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	fruitArr := []string{"Apple", "Orange"}
	fruitArr[1] = "Banana"
	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

	a := []int{1, 2, 3, 4}
	b := a
	b[1] = 5
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(a)
	fmt.Println(b)

}
