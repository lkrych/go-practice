package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./aocday1.txt")
	if err != nil {
		log.Fatal(err)
	}

	freqLog := map[int]bool{}
	runningTotal := 0
	foundRepFreq := false
	splitByLine := strings.Split(string(file), "\n")

	for foundRepFreq != true {
		fmt.Println("Restarting with runningTotal", runningTotal)
		for i := 0; i < len(splitByLine)-1; i++ {
			line := splitByLine[i]
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
			if _, ok := freqLog[runningTotal]; ok {
				fmt.Println("The first repeated Frequency is", runningTotal)
				foundRepFreq = true
				break
			}
			freqLog[runningTotal] = true
			fmt.Println("The running total is,", runningTotal)
		}
	}

}
