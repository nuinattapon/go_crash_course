package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

type Triangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (t Triangle) area() float64 {
	return t.width * t.height * 0.5
}

func main() {
	var circle Shape = Circle{x: 0, y: 0, radius: 5}
	var rectangle Shape = Rectangle{width: 10, height: 5}
	var triangle Shape = Triangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %.2f\n", circle.area())
	fmt.Printf("Rectangle Area: %.2f\n", rectangle.area())
	fmt.Printf("Triangle Area: %.2f\n", triangle.area())

}
