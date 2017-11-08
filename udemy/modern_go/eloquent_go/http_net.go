package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	//client
	resp, err := http.Get("http://example.com/")
	defer resp.Body.Close() //make sure you close the response!
	body, e := ioutil.ReadAll(resp.Body)
}
