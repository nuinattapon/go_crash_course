package main

import "fmt"

func main() {
	// a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// b := a[:]
	// c := a[3:]
	// d := a[:6]
	// e := a[3:6]

	// fmt.Printf("%+v\n", a)
	// fmt.Printf("%+v\n", b)
	// fmt.Printf("%+v\n", c)
	// fmt.Printf("%+v\n", d)
	// fmt.Printf("%+v\n", e)

	// a := make([]int, 3, 100) // This is sutable for bigger slice - preallocate 100 slots
	a := []int{} // This is suitable for small size slice
	fmt.Printf("a: %+v, length: %+v, capacity: %+v\n", a, len(a), cap(a))
	a = append(a, 1)
	fmt.Printf("a: %+v, length: %+v, capacity: %+v\n", a, len(a), cap(a))
	a = append(a, 2, 3, 4)
	fmt.Printf("a: %+v, length: %+v, capacity: %+v\n", a, len(a), cap(a))
	a = append(a, []int{5, 6, 7, 8}...)
	fmt.Printf("a: %+v, length: %+v, capacity: %+v\n", a, len(a), cap(a))
	b := a[1:]
	fmt.Printf("b: %+v, length: %+v, capacity: %+v\n", b, len(b), cap(b))
	c := a[:len(a)-1]
	fmt.Printf("c: %+v, length: %+v, capacity: %+v\n", c, len(c), cap(c))
	// c = append(c, 8)
	// fmt.Printf("c: %+v, length: %+v, capacity: %+v\n", c, len(c), cap(c))
	fmt.Printf("a: %+v, length: %+v, capacity: %+v\n", a, len(a), cap(a))

	// This kind of code is discouraging
	// d := append(a[:2], a[3:]...)
	// fmt.Printf("d: %+v, length: %+v, capacity: %+v\n", d, len(d), cap(d))
	// fmt.Printf("a: %+v, length: %+v, capacity: %+v\n", a, len(a), cap(a))

}
