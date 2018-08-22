package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	File  string `short:"f" long:"file"  description:"a file to read DFA input from."`
	Input string `short:"i" long:"input" description:"a string input to match against the DFA"`
}

func main() {
	//interpret the CL stdin
	flags.Parse(&opts)
	file, err := ioutil.ReadFile(opts.File)
	checkErr(err)
	//create an array of strings that is split by the newline character
	splitByLine := strings.Split(string(file), "\n")
	DFA := createDFA(splitByLine)
	match := DFA.isMatch(opts.Input)
	if match {
		fmt.Printf("Yes, %v matches the DFA represented by %v \n", opts.Input, opts.File)
	} else {
		fmt.Printf("No, %v does not match the DFA represented by %v \n", opts.Input, opts.File)
	}

}

func createDFA(dfaInput []string) DFA {
	newDFA := DFA{}
	symbols := []string{}
	for i := 0; i < len(dfaInput); i++ {
		line := dfaInput[i]
		if i == 0 {
			newDFA.action = []bool{}
			actions := strings.Split(line, " ")
			for i := 0; i < len(actions); i++ {
				newAction, err := strconv.ParseBool(actions[i])
				checkErr(err)
				newDFA.action = append(newDFA.action, newAction)
			}
		} else if i == 1 {
			//initialize map and symbols
			symbols = strings.Split(line, "")
			for _, val := range symbols {
				if _, ok := newDFA.pathMap[val]; ok {
					continue
				} else {
					newDFA.pathMap[val] = []int{}
				}
			}
		} else {
			nodes := strings.Split(line, " ")
			for _, stringNode := range nodes {
				nextNode, err := strconv.Atoi(stringNode)
				checkErr(err)
				for _, symbol := range symbols {
					pathMapVal, _ := newDFA.pathMap[symbol]
					newDFA.pathMap[symbol] = append(pathMapVal, nextNode)
				}
			}
		}
	}
	return newDFA
}

type DFA struct {
	action  []bool
	pathMap map[string][]int
}

func (d *DFA) isMatch(stringToMatch string) bool {
	currentNode := 0
	for _, char := range strings.Split(stringToMatch, "") {
		currentNode = d.pathMap[char][currentNode]
	}
	return d.action[currentNode]
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
