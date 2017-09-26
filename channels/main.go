package main

import (
	"fmt"
	"net/http"
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
	fmt.Println(<-c)
}

func makeRequest(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "The link might be down"
		return
	}
	fmt.Println(link, "is responding to http requests")
	c <- "The link is responding to http requests"
}
