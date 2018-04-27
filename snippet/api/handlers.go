package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
)

var feedMap = map[string]string{
	"the weeds":           "http://feeds.feedburner.com/voxtheweeds",
	"worldly":             "http://feeds.feedburner.com/voxworldly",
	"the ezra klein show": "http://feeds.feedburner.com/voxworldly",
}

var searchTermHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the searchTerm handler")
	params := mux.Vars(r)

	fmt.Println("The params are ", params)

	fp := gofeed.NewParser()
	feedURL, _ := feedMap[strings.ToLower(params["keyword"])]
	feed, err := fp.ParseURL(feedURL)
	if err != nil {
		log.Fatalln("Cannot find feed for ", params["keyword"])
		log.Fatalln("The error is ", err)
	}
	fmt.Println(feed.Title)
	fmt.Println(feed.Author)
	fmt.Println(feed.Description)
	fmt.Println(feed.Link)

	//marshal feed into json
	feedJSON, err := json.Marshal(feed)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Sending Feed Results")
	w.Write(feedJSON)
	printStars()
})

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}

//separates requests in log
func printStars() {
	fmt.Println()
	fmt.Println("**************************************")
	fmt.Println()
}
