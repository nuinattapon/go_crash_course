package main

import (
	"encoding/json"
	"fmt"
)

type Circle struct {
	Radius int `json:"radius"`
	Border int `json:"border"`
}

func main() {
	x := 20
	f := 123.45

	// Basic formatting
	fmt.Printf("%d\n", x)
	fmt.Printf("%x\n", x)

	// Booleans can be printed as "true" or "false"
	fmt.Printf("%t\n", x > 10)

	// floating point numbers
	fmt.Printf("%f\n", f)
	fmt.Printf("%e\n", f)

	// Using explicit argument indexes
	fmt.Printf("%[2]d %[1]d\n", 52, 40)
	// Argument indexes can be used to print values repeatedly
	fmt.Printf("%d %#[1]o %#[1]x\n", 52)
	fmt.Printf("%d %[1]d %[1]d\n", 52)

	// Print a value in default format
	c := Circle{
		Radius: 20,
		Border: 5,
	}
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%T\n", c)

	if b, err := json.Marshal(c); err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Println(string(b))
	}

	// Sprintf is the same as Printf, but returns a string
	s := fmt.Sprintf("%[2]d %[1]d\n", 52, 40)
	fmt.Println(s)
}
