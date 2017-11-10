package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	readAndCreateFiles("exercises_itemized")
}

func readAndCreateFiles(directory string) {
	//create test files for all exercises in directory
	files, err := ioutil.ReadDir(fmt.Sprintf("./%v", directory))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		split := strings.Split(f.Name(), ".")
		testName := fmt.Sprintf("%v_test.go", split[0])
		f, err := os.Create(fmt.Sprintf("./exercises_test/%v", testName))
		check(err)
		defer f.Close()
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
