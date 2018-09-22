package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Filename string `short:"f" long:"file"  description:"a .vm file or a directory with .vm files."`
}

func main() {

	//interpret the CL stdin
	flags.Parse(&opts)
	fileName := opts.Filename

	// input could either be a .vm file or a directory
	// manage this using a slice of file pointers
	vmFiles := findVMFiles(fileName)
	filePointers := []*os.File{}

	fmt.Println("There are ", len(vmFiles), " .vm files")
	for _, vmFile := range vmFiles {
		fp, err := os.Open(vmFile)
		if err != nil {
			fmt.Println("Err! The file cannot open")
			os.Exit(3)
		}
		filePointers = append(filePointers, fp)
	}

	//setup assembly file
	asmFile := ""
	if strings.HasSuffix(fileName, ".vm") {
		asmFile = strings.TrimSuffix(fileName, ".vm") + ".asm"
	} else {
		asmFile = fileName + ".asm"
	}
	f, err := os.Create(asmFile)
	checkErr(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	for _, fp := range filePointers {
		scanner := bufio.NewScanner(fp)

		// construct a parser to handle the input
		// construct a codewriter to handle the output
		//move through the lines in the input
	}
}

// findVMFiles takes in the CL input and returns an array of all
// the .vm files that need to be parsed into an .asm file
func findVMFiles(fileName string) []string {
	vmFiles := []string{}

	fi, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Usage: VMTranslator foo.vm or ./foo which is directory with .vm files  ")
		fmt.Println("The program is terminating 💀")
		os.Exit(2)
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		// recursively add .vm files to slice
		files, err := ioutil.ReadDir(fileName)
		if err != nil {
			fmt.Println(err)
			return vmFiles
		}
		for _, file := range files {
			//make recursive call
			vmFile := findVMFiles(fileName + "/" + file.Name())
			vmFiles = append(vmFiles, vmFile...)
		}

	case mode.IsRegular():
		if strings.HasSuffix(fileName, ".vm") {
			vmFiles = append(vmFiles, fileName)
		}
	}

	return vmFiles
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
