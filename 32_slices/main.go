package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 3, 20)
	copy(y, x[1:])

	fmt.Println(x, len(x), cap(x))
	fmt.Println(y, len(y), cap(y))

}
