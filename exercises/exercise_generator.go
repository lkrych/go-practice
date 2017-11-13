package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//ask for number of Exercises
	numExercises := askForExercises()
	//create quiz folder with exercises and tests
	makeExercises(numExercises)
}

//use stdinput to ask for a number of Exercises to practice
func askForExercises() int {
	numExercises := 0
	for { //iterate until the number is between bounds
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter number of exercises that you want to practice (min: 1, max: 10): ")
		numExercisesString, err := reader.ReadString('\n')
		numExercises, err = strconv.Atoi(strings.Split(numExercisesString, "\n")[0])

		if err != nil {
			formatPrint("That isn't a valid number!")
		}

		if numExercises < 1 || numExercises > 10 {
			formatPrint("You need to pick an integer between 1 and 10")
		} else {
			break
		}
	}
	return numExercises
}

func makeExercises(numExercises int) {

	//create directory and files
	createFiles()

	//read exercises
	files, err := ioutil.ReadDir("./exercises_itemized")
	checkErr(err)

	//grab random indices corresponding to files in files slice
	randIdxs := randomIndices(numExercises, len(files))

	//write files to output
	exerciseFile, err := os.OpenFile("./go_exercises/exercises.go", os.O_RDWR, 0666)
	checkErr(err)
	defer exerciseFile.Close()
	testFile, err := os.OpenFile("./go_exercises/exercises_test.go", os.O_RDWR, 0666)
	checkErr(err)
	defer testFile.Close()

	// add exercises and tests to file
	for _, idx := range randIdxs {
		fileName := files[idx].Name()

		//read exercise file
		exercise, err := ioutil.ReadFile(fmt.Sprintf("./exercises_itemized/%v", fileName))
		checkErr(err)

		// read exercise test file
		split := strings.Split(fileName, ".")
		testName := fmt.Sprintf("%v_test.go", split[0])
		test, err := ioutil.ReadFile(fmt.Sprintf("./exercises_test/%v", testName))
		checkErr(err)

		//check to make sure files are being read correctly
		fmt.Println("Exercises and Tests for", fileName)

		_, err = exerciseFile.Write(exercise)
		checkErr(err)
		_, err = exerciseFile.Write([]byte("\n"))
		checkErr(err)
		_, err = testFile.Write(test)
		checkErr(err)
		_, err = testFile.Write([]byte("\n"))
		checkErr(err)
	}
}

func createFiles() {
	os.Mkdir("./go_exercises", os.ModePerm)
	execs, err := os.Create("./go_exercises/exercises.go")
	checkErr(err)
	execs.Close()
	tests, err := os.Create("./go_exercises/exercises_test.go")
	checkErr(err)
	tests.Close()
}

func randomIndices(n int, arrLen int) []int {
	//a map is a more efficient way to enforce unique idxs
	m := make(map[int]bool)
	//seed
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	idxCount := 0
	for idxCount < n {
		//generate a random int
		newIdx := r1.Intn(arrLen)
		//check if that int has been found
		if _, ok := m[newIdx]; ok {
			//if it has, continue looping
			continue
		}
		//otherwise add it to the map and increase the idx count
		m[newIdx] = true
		idxCount++
	}
	//create arr of idxs
	keys := make([]int, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func addExercisesToFiles() {

}

// print funcs for formatting
func printSpace() {
	fmt.Printf("\n")
}

func printBorder() {
	fmt.Printf("==================================================================== \n")
}

func formatPrint(printThis string) {
	printBorder()
	fmt.Println(printThis)
	printBorder()
	printSpace()
}

//
func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}
