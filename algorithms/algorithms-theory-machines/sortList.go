package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"

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
	createFile()

	//read exercises
	files, err := ioutil.ReadDir(opts.File)
	checkErr(err)
}

func createFile() {
	if _, err := os.Stat("./generatedStrings.txt"); err != nil {
		//file does not exist
		//create file
		f, err := os.Create("./generatedStrings.txt")
		checkErr(err)
		defer f.Close()
		strings := createStringInput()
		f.Write(strings)
	}
	return
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func createStringInput() []byte {
	strings := []byte{}
	for i := 0; i < 500; i++ {
		newWord := generateNewWord(rand.Intn(12))
		newWord = append(newWord, '\n')
		strings = append(strings, newWord)
	}
	return strings
}

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func generateNewWord(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}
