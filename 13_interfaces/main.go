package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Rectangle) Perimeter() float64 {
	return 2*c.Width + 2*c.Height
}

func main() {
	var circle Shape = Circle{Radius: 5}
	var rectangle Shape = Rectangle{Width: 10, Height: 5}

	fmt.Printf("Circle area: %.2f, circle perimeter: %.2f\n", circle.Area(), circle.Perimeter())
	fmt.Printf("Rectangle area: %.2f, rectangle perimeter: %.2f\n", rectangle.Area(), rectangle.Perimeter())

}
