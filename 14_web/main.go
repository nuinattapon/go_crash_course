package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Hello, 世界, สวัสดี</h1>\n")
	fmt.Fprintf(w, "<h2>URI: %s <h2>\n", r.URL.RequestURI())
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>About</h1>")
}

func text(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Hello, 世界, สวัสดี\n")
	fmt.Fprintln(w, "URI:", r.URL.RequestURI())

}

func jsonHandler(w http.ResponseWriter, r *http.Request) {

	type Message struct {
		Id      int    `json:"id"`
		Message string `json:"message"`
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	message := Message{Id: 1, Message: "Hello, 世界, สวัสดี"}
	json.NewEncoder(w).Encode(message)

}

func main() {
	http.HandleFunc("/", index)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/about", about)
	http.HandleFunc("/text", text)
	http.HandleFunc("/json", jsonHandler)

	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
