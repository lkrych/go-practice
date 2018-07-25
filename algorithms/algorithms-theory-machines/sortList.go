package main

import (
	"io/ioutil"

	flags "github.com/jessevdk/go-flags"
)

//Write a filter that reads strings from standard input
//and prints them to standard output, in sorted order with all duplicate strings removed.

var opts struct {
	File string `short:"f" long:"file" default:"generatedStrings.txt" description:"a file to read strings from."`
}

func main() {
	flags.Parse(&opts)

	//check if file exists, if it doesn't use func to create a file
	createFiles()

	//read exercises
	files, err := ioutil.ReadDir("./exercises_itemized")
	checkErr(err)
}
