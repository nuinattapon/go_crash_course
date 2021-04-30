package main

import (
	"fmt"
)

func main() {

	// a:= 42
	// b := &a

	var a int = 42
	var b *int = &a

	a = 24
	fmt.Println(&a, b)
	fmt.Println(a, *b)

	*b = 15
	fmt.Println(&a, b)
	fmt.Println(a, *b)

	d := [3]int{1, 2, 3}
	e := &d[0]
	f := &d[1]
	fmt.Printf("%v %p %p\n", d, e, f)

	var ms *myStruct

	ms = &myStruct{foo: 42}
	fmt.Printf("ms = %+v\n", ms)

	ms2 := new(myStruct)
	(*ms2).foo = 10
	// Go compiler helps us to dereference the pointer for us
	ms2.foo = 12
	fmt.Printf("ms2 = %+v\n", ms2)
}

type myStruct struct {
	foo int64
}
