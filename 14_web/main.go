package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var version = "1.0"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Hello, 世界, สวัสดี</h1>\n")
	fmt.Fprintf(w, "<strong>URI: %s </strong>\n", r.URL.RequestURI())
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<strong>About</strong>")
}

func textHandler(w http.ResponseWriter, r *http.Request) {
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

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "%s\n", version)
	fmt.Fprintln(w, "URI:", r.URL.RequestURI())

}

func main() {
	http.HandleFunc("/", indexHandler)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/text", textHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/version", versionHandler)

	portID := 3000
	s := strconv.Itoa(portID)

	fmt.Printf("Server Starting at port %s...", s)
	http.ListenAndServe(":"+s, nil)

}
