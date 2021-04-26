package main

import "fmt"

func main() {
	// Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		// i = i + 1
		i++
	}

	// Short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("%3d FizzBuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%3d Fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%3d Buzz\n", i)
		} else {
			fmt.Printf("%3d\n", i)
		}
	}
}
