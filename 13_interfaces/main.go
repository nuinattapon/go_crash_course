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

type RightTriangle struct {
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

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func (t RightTriangle) Area() float64 {
	return t.Width * t.Height / 2
}

func (t RightTriangle) Perimeter() float64 {
	return t.Width + t.Height + math.Sqrt(math.Pow(t.Width, 2)+math.Pow(t.Height, 2))
}

func main() {
	var circle Shape = Circle{Radius: 5}
	var rectangle Shape = Rectangle{Width: 3, Height: 4}
	var triangle Shape = RightTriangle{Width: 3, Height: 4}
	fmt.Printf("Circle area: %.2f,  perimeter: %.2f\n", circle.Area(), circle.Perimeter())
	fmt.Printf("Rectangle area: %.2f,  perimeter: %.2f\n", rectangle.Area(), rectangle.Perimeter())
	fmt.Printf("Right Triangle area: %.2f,  perimeter: %.2f\n", triangle.Area(), triangle.Perimeter())

}
