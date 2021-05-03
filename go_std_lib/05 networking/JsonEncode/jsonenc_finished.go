package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Address    string   `json:"addr"`
	Age        int      `json:"age"`
	FaveColors []string `json:"favecolors,omitempty"`
	Password   string   `json:"-"`
}

func encodeExample() {
	// create some people data
	people := []person{
		{"Jane", "Doe", "123 Anywhere Street", 35, nil, "password"},
		{"John", "Public", "456 Everywhere Blvd", 29, []string{"Purple", "Yellow", "Green"}, "password"},
		{"ณัฐพล", "Public", "456 มัณฑนา Blvd", 29, []string{"Purple", "Yellow", "Green"}, "password"},
	}

	// Marshal is used to convert a data structure to JSON format
	// MarshalIndent is used to format the JSON string with indentation
	result, err := json.MarshalIndent(people, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(result))
}

func main() {
	// Encode Go data as JSON
	encodeExample()
}
