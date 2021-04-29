package main

import (
	"fmt"
	"math"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func main() {
	a := 0.1 + 1e-10
	b := 0.2
	fmt.Printf("%.10f\n", a+b)
	fmt.Printf("%.10f\n", 0.3)

	fmt.Println(almostEqual(a+b, 0.3))
}
