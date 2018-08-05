package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	flags "github.com/jessevdk/go-flags"
)

//Write a filter that reads strings from standard input
//and prints them to standard output, in sorted order with all duplicate strings removed.

var opts struct {
	File     string `short:"f" long:"file" default:"generatedStrings.txt" description:"a file to read strings from."`
	NumWords int    `short:"n" long:"num" default:"500" description:"a number of strings to create."`
}

func main() {
	//interpret the CL stdin
	flags.Parse(&opts)

	//make sure that the random number generator has a unique seed
	rand.Seed(time.Now().UTC().UnixNano())
	//check if file is default, if so, create the file if it doesn't exist
	if opts.File == "generatedStrings.txt" {
		createFile(opts.NumWords)
	}

	//read from the file
	file, err := ioutil.ReadFile(opts.File)
	checkErr(err)
	//create an array of strings that is split by the newline character
	splitByNewline := strings.Split(string(file), "\n")
	//deduplicate words by throwing them into hash
	deDup := map[string]bool{}
	for _, word := range splitByNewline {
		deDup[word] = true
	}
	//get an array of keys from the hash
	keys := make([]string, len(deDup))
	i := 0
	for k := range deDup {
		keys[i] = k
		i++
	}

	//sort the keys
	sorted := quickSort(keys)
	for _, word := range sorted {
		//print to stdout
		fmt.Println(word)
	}

}

func createFile(numWords int) {
	if _, err := os.Stat("./generatedStrings.txt"); err != nil {
		//if file does not exist
		//create file
		f, err := os.Create("./generatedStrings.txt")
		checkErr(err)
		defer f.Close()

		for i := 0; i < numWords; i++ {
			strings := createStringInput()
			f.Write(strings)
		}
	}
	return
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func createStringInput() []byte {
	newWord := generateNewWord(rand.Intn(12))
	newWord = append(newWord, '\n')
	return newWord
}

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func generateNewWord(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func quickSort(strings []string) []string {
	if len(strings) <= 1 {
		return strings
	}

	pivot := strings[0]
	left := []string{}
	right := []string{}

	for i := 1; i < len(strings); i++ {
		currString := strings[i]
		if currString > pivot {
			right = append(right, currString)
		} else {
			left = append(left, currString)
		}
	}

	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)

	sorted := append(sortedLeft, pivot)
	sorted = append(sorted, sortedRight...)

	return sorted
}
