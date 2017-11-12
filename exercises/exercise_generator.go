package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//ask for number of questions
	numQuestions := askForQuestions()
	fmt.Printf("Congrats! You picked a valid number! %v \n", numQuestions)
	//create quiz folder with exercises and tests
}

//use stdinput to ask for a number of questions to practice
func askForQuestions() int {
	numQuestions := 0
	for { //iterate until the number is between bounds
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter number of exercises that you want to practice (min: 1, max: 10): ")
		numQuestionsString, err := reader.ReadString('\n')
		numQuestions, err = strconv.Atoi(strings.Split(numQuestionsString, "\n")[0])

		if err != nil {
			formatPrint("That isn't a valid number!")
		}

		if numQuestions < 1 || numQuestions > 10 {
			formatPrint("You need to pick an integer between 1 and 10")
		} else {
			break
		}
	}
	return numQuestions
}

// print funcs
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
