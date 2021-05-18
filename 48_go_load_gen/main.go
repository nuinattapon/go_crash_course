// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {

}
func main() {
	start := time.Now()
	url := []string{"https://go.nattapon.me", "https://go.nattapon.me/ping", "https://go.nattapon.me/version", "https://go.nattapon.me/json"}

	ch := make(chan string)
	for _, url := range url {
		go fetch(url, ch) // start a goroutine
	}
	for range url {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("Elapsed Time %5.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	// start := time.Now()
	p := message.NewPrinter(language.English)

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	// nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	// nbytes := len(body)
	// fmt.Printf("%s\n", body)
	// secs := time.Since(start).Seconds()
	// ch <- p.Sprintf("Elapsed Time %5.2fs - %5.2fKB - URL  %s", secs, float64(nbytes)/1024.0, url)
	ch <- p.Sprintf("%s", body)
}
