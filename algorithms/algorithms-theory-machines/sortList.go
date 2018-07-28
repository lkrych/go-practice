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
	File string `short:"f" long:"file" default:"generatedStrings.txt" description:"a file to read strings from."`
}

func main() {
	flags.Parse(&opts)

	//check if file exists, if it doesn't use func to create a file
	rand.Seed(time.Now().UTC().UnixNano())
	createFile()

	//read exercises
	file, err := ioutil.ReadFile(opts.File)
	checkErr(err)
	splitByNewline := strings.Split(string(file), "\n")
	//deduplicate words by throwing into hash
	deDup := map[string]bool{}
	for _, word := range splitByNewline {
		deDup[word] = true
	}
	//get an array of keys
	keys := make([]string, len(deDup))

	i := 0
	for k := range deDup {
		keys[i] = k
		i++
	}
	//sort the keys
	sorted := quickSort(keys)
	for _, word := range sorted {
		fmt.Println(word)
	}
}

func createFile() {
	if _, err := os.Stat("./generatedStrings.txt"); err != nil {
		//file does not exist
		//create file
		f, err := os.Create("./generatedStrings.txt")
		checkErr(err)
		defer f.Close()

		for i := 0; i < 500; i++ {
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
