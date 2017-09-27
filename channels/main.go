package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go makeRequest(link, c)
	}

	//range function grabs link from channel
	for l := range c {
		go func(link string) {
			time.Sleep(3 * time.Second)
			makeRequest(link, c)
		}(l)
	}
}

func makeRequest(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is responding to http requests")
	c <- link
}
