package main

import (
	"fmt"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	// Original return clause
	// return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
	return fmt.Sprintf("Hello, my name is %s %s and I am %d", p.firstName, p.lastName, p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	firstPerson := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	// Alternative
	secondPerson := Person{"Bob", "Johnson", "New York", "m", 30}

	// fmt.Println(firstPerson.firstName)
	fmt.Printf("firstPerson = %v\n", firstPerson)
	firstPerson.age++
	fmt.Printf("firstPerson = %v\n", firstPerson)

	fmt.Println(firstPerson.greet())
	fmt.Println(secondPerson.greet())

	firstPerson.hasBirthday()
	secondPerson.hasBirthday()

	firstPerson.getMarried("Williams")
	secondPerson.getMarried("Thompson")

	fmt.Println(firstPerson.greet())
	fmt.Println(secondPerson.greet())

	aDoctor := struct{ name string }{name: "John Pertwee"}
	anotherDoctor := &aDoctor

	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor)
	fmt.Println(*anotherDoctor)
}
