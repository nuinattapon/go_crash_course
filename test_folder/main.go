package main

import (
	"fmt"
	"strconv"
)

var (
	name    string = "Nui"
	surname string = "Sub"
	i       int64  = 10
)

func main() {
	// var i int = 42
	fmt.Printf("Type of i is %T and its value is %v\n", i, i)

	i := 42
	f := 45.67

	fmt.Printf("Type of i is %T and its value is %v\n", i, i)

	fmt.Printf("Type of f is %T and its value is %v\n", f, f)

	str := "1234"
	if j, err := strconv.ParseInt(str, 10, 64); err == nil {
		fmt.Printf("Type of j is %T and its value is %v\n", j, j)
	} else {
		fmt.Println(err)
	}

	str2 := "123.4567"
	if k, err := strconv.ParseFloat(str2, 64); err == nil {
		fmt.Printf("Type of k is %T and its value is %v\n", k, k)
	} else {
		fmt.Println(err)
	}

	l := float64(i)
	fmt.Printf("Type of l is %T and its value is %v\n", l, l)

	var m int64
	fmt.Printf("Type of m is %T and its value is %v\n", m, m)

	o := 1 + 2i
	fmt.Printf("Type of o is %T and its value is %v\n", o, o)

	s := "Hello, 世界, สวัสดี"
	fmt.Printf("Type of s is %T and its value is %v, its len is %d\n", s, s, len(s))
	b := []byte(s)
	fmt.Printf("Type of r is %T and its value is %v, its len is %d\n", b, b, len(b))

	r := []rune(s)
	fmt.Printf("Type of r is %T and its value is %v, its len is %d\n", r, r, len(r))

	s2 := string(r)
	fmt.Printf("Type of s2 is %T and its value is %v\n", s2, s2)

	r2 := '世'
	fmt.Printf("Type of r2 is %T and its value is %v\n", r2, r2)

	r3 := 'A'
	fmt.Printf("Type of r3 is %T and its value is %v\n", r3, r3)
}
