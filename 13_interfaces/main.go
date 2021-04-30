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
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

type Triangle struct {
	Width, Height float64
	// Rectangle
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

func (t Triangle) area() float64 {
	return t.Width * t.Height * 0.5
}

func main() {
	var circle Shape = Circle{Radius: 5}
	var rectangle Shape = Rectangle{Width: 10, Height: 5}
	var triangle Shape = Triangle{Width: 10, Height: 5}

	fmt.Printf("Circle Area: %.2f\n", circle.area())
	fmt.Printf("Rectangle Area: %.2f\n", rectangle.area())
	fmt.Printf("Triangle is %+v and its area is %.2f\n", triangle, triangle.area())

}
