package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type contentReplied struct {
	URL    string `json:"url"`
	Origin string `json:"origin"`
}

type person struct {
	Name       string   `json:"fullname"`
	Address    string   `json:"addr"`
	Age        int      `json:"age"`
	FaveColors []string `json:"favecolors,omitempty"`
}

func postRequestTest() {
	const httpbin = "https://httpbin.org/post"
	// create some people data
	people := []person{
		{"Jane Doe", "123 Anywhere Street", 35, nil},
		{"John Public", "456 Everywhere Blvd", 29, []string{"Purple", "Yellow", "Green"}},
	}

	// jsonStr, err := json.MarshalIndent(people, "", "\t")
	jsonStr, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(httpbin, "application/json; charset=utf-8", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)

	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	// We can access parts of the response to get information:
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Protocol:", resp.Proto)
	fmt.Println("Content length:", resp.ContentLength)
	fmt.Println("Header:")
	header := resp.Header
	for k, v := range header {
		fmt.Printf("- %[1]s : %[2]q\n", k, v)
	}
	fmt.Println("Content:")
	fmt.Printf("%s\n", content)

	message := contentReplied{}
	err = json.Unmarshal(content, &message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Struct: %#v\n", message)

}

func main() {
	// Execute a POST
	postRequestTest()
}
