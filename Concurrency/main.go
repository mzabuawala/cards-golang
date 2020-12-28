package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.fb.com",
		"http://www.yahoo.com",
		"http://www.amazon.com",
		"http://www.stackoverflow.com",
		"http://www.torrentz2.com",
	}
	ch := make(chan string)
	for _, url := range urls {
		go checkURLStatus(url, ch)
	}

	// Receving a message from a Channel is a blocking call
	for l := range ch {
		go func(ln string) {
			time.Sleep(time.Second * 5)
			checkURLStatus(ln, ch)
		}(l)
	}
}

func checkURLStatus(url string, ch chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("DOWN\t:", url)
		ch <- url
	} else {
		fmt.Println("UP\t:", url)
		ch <- url
	}
}
