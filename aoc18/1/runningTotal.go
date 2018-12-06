package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./aocday1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	runningTotal := 0
	for scanner.Scan() {
		line := scanner.Text()
		operator := line[0]
		operand := line[1:]
		convertedOp, err := strconv.Atoi(operand)
		if err != nil {
			log.Fatalln("There was an error converting the input into an integer")
		}
		if operator == '+' {
			runningTotal = runningTotal + convertedOp
		} else if operator == '-' {
			runningTotal = runningTotal - convertedOp
		}
	}
	fmt.Println("The running total is", runningTotal)

}
