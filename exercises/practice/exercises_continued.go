package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func recursiveCircus(filepath string) string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	diskMap := make(map[string]bool)
	splitByLine := strings.Split(string(content), "\n")
	for i, line := range splitByLine {
		fmt.Printf("Idx %v: %v \n", i, line)
		splitLine := strings.Split(line, " ")
		fmt.Printf("Split: %v, len: %v \n", splitLine, len(splitLine))
		//remove from map
		if len(splitLine) > 2 {
			for _, el := range splitLine[3:] {
				newEntry := strings.Replace(el, ",", "", -1)
				diskMap = checkIfInMap(newEntry, diskMap)
			}
		}
		entry := splitLine[0]
		diskMap = checkIfInMap(entry, diskMap)
	}

	rootOfDisk := ""
	for k := range diskMap {
		rootOfDisk = k
	}
	return rootOfDisk
}

func checkIfInMap(el string, myMap map[string]bool) map[string]bool {
	if _, ok := myMap[el]; ok {
		delete(myMap, el)
	} else {
		myMap[el] = true
	}
	return myMap
}

type Node struct {
	name     string
	children []string
	val      int
}

func buildCircusTree(filepath string) map[string]*Node {
	nodeTable := make(map[string]*Node)
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	splitByLine := strings.Split(string(content), "\n")

	//build map and nodes
	for _, line := range splitByLine {
		// fmt.Printf("Idx %v: %v \n", i, line)
		splitLine := strings.Split(line, " ")
		// fmt.Printf("Split: %v, len: %v \n", splitLine, len(splitLine))
		//populate map
		entry := splitLine[0]
		newNode := Node{
			name:     entry,
			val:      convertToInt(splitLine[1]),
			children: []string{},
		}
		if len(splitLine) > 2 {
			for _, el := range splitLine[3:] {
				newEntry := strings.Replace(el, ",", "", -1)
				newNode.children = append(newNode.children, newEntry)
			}
		}
		// fmt.Printf("New node: %v \n", newNode)
		nodeTable[entry] = &newNode
		// fmt.Printf("nodeTable: %v \n", nodeTable)
	}
	return nodeTable
}

func convertToInt(s string) int {
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	integerS, _ := strconv.Atoi(s)
	return integerS
}

func recursiveCircusDifference(nodeMap map[string]*Node, nodeName string) int {
	// fmt.Printf("NodeMap: %v \n", nodeMap)

	childrenOfNode := nodeMap[nodeName].children
	valOfNode := nodeMap[nodeName].val
	childVals := make(map[string]int)
	// fmt.Printf("Node: %v \n", nodeName)

	if len(childrenOfNode) > 0 {
		for i, child := range childrenOfNode {
			// fmt.Printf("ChildNode: %v, Val: %v \n", child, nodeMap[child].val)
			newVal := recursiveCircusDifference(nodeMap, child)
			if i != 0 {
				if val := childVals[nodeName]; val != newVal {
					fmt.Printf("Child vals: %v from child %v, new entry: %v \n", childVals, child, newVal)
					fmt.Printf("%v 's value is %v \n", child, nodeMap[child].val)
					fmt.Printf("There is a difference between child nodes of %v \n", nodeName)
				}
			}
			childVals[nodeName] = newVal
		}
		return valOfNode + (childVals[nodeName] * len(childrenOfNode))
	}
	return valOfNode

}

func recursiveCircusTree(filepath string) int {
	nodeMap := buildCircusTree(filepath)
	diff := recursiveCircusDifference(nodeMap, "eugwuhl")
	return diff
}

func register(filepath string) (int, int) {
	registerMap := make(map[string]int)
	maxVal := 0
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	splitByLine := strings.Split(string(content), "\n")
	for _, line := range splitByLine {
		splitLine := strings.Split(line, " ")

		cmdRegister := splitLine[0]
		primaryCommand := splitLine[1]
		cmdVal := splitLine[2]
		condRegister := splitLine[4]
		condCommand := splitLine[5]
		condVal := splitLine[6]

		//check if condRegister exists
		if _, ok := registerMap[condRegister]; ok {
			//check condition
			if checkCondition(condRegister, condCommand, condVal, registerMap) {
				maxVal = performCommand(cmdRegister, primaryCommand, cmdVal, registerMap, maxVal)
			} else {
				continue
			}
		} else {
			//init condRegister
			registerMap[condRegister] = 0
			if checkCondition(condRegister, condCommand, condVal, registerMap) {
				maxVal = performCommand(cmdRegister, primaryCommand, cmdVal, registerMap, maxVal)
			} else {
				continue
			}
		}
	}

	currMax := 0
	for _, value := range registerMap {
		if value > currMax {
			currMax = value
		}
	}
	return currMax, maxVal
}

func checkCondition(register string, command string, val string, myMap map[string]int) bool {
	registerVal := 0
	if _, ok := myMap[register]; ok {
		registerVal = myMap[register]
	}
	newVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	switch command {
	case ">":
		return registerVal > newVal
	case "<":
		return registerVal < newVal
	case ">=":
		return registerVal >= newVal
	case "==":
		return registerVal == newVal
	case "<=":
		return registerVal <= newVal
	case "!=":
		return registerVal != newVal
	}
	return false
}

func performCommand(register string, command string, val string, myMap map[string]int, maxVal int) int {
	registerVal := 0
	if _, ok := myMap[register]; ok {
		registerVal = myMap[register]
	}
	newVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	switch command {
	case "inc":
		myMap[register] = registerVal + newVal
	case "dec":
		myMap[register] = registerVal - newVal
	}

	if myMap[register] > maxVal {
		maxVal = myMap[register]
	}
	return maxVal
}

func streamParser(stream string) (int, int) {
	bracketCount := 0
	totalCount := 0
	garbageCount := 0
	openGarbage := false
	skipChar := false
	splitByLetter := strings.Split(stream, "")
	for i, char := range splitByLetter {
		if i > 0 {
			if skipChar {
				skipChar = false
				continue
			}
			if char == "!" {
				skipChar = true
				continue
			}
			if char == ">" {
				openGarbage = false
				continue
			}
			if openGarbage {
				garbageCount++
				continue
			}
		}
		switch char {
		case "<":
			{
				openGarbage = true
			}
		case ">":
			{
				openGarbage = false
			}
		case "{":
			{
				bracketCount++
				totalCount += bracketCount
			}
		case "}":
			{
				bracketCount--
			}
		}
	}
	return totalCount, garbageCount
}

func streamParserAdvent(filepath string) (int, int) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return streamParser(string(content))
}

func knotHashMinimal(input []int) int {
	knot := []int{0, 1, 2, 3, 4}
	skipSize := 0
	currIdx := 0
	for _, el := range input {
		knot = reverseAndWrap(knot, el, currIdx)
		currIdx += el + skipSize
		if currIdx > len(knot)-1 {
			currIdx = currIdx % len(knot)
		}
		skipSize++
		fmt.Printf("####################### \n")
		fmt.Printf("The knot is %v \n", knot)
		fmt.Printf("####################### \n")
	}
	return knot[0] * knot[1]
}

func knotHash(input []int) int {
	knot := make([]int, 256)
	for i := 0; i < len(knot); i++ {
		knot[i] = i
	}
	skipSize := 0
	currIdx := 0
	for _, el := range input {
		fmt.Printf("The currentIndex is %v \n", currIdx)
		knot = reverseAndWrap(knot, el, currIdx%len(knot))
		currIdx += el + skipSize
		skipSize++
		fmt.Printf("####################### \n")
		fmt.Printf("The knot is %v \n", knot)
		fmt.Printf("####################### \n")
	}
	return knot[0] * knot[1]
}

func reverseAndWrap(arr []int, el int, currIdx int) []int {
	newArr := []int{}
	remainder := 0
	if (currIdx + el) > len(arr) {
		remainder = ((currIdx + el) % len(arr))
	}
	// fmt.Printf("The array is %v \n", arr)
	// fmt.Printf("The current index is %v \n", currIdx)
	// fmt.Printf("The el is %v \n", el)
	// fmt.Printf("The remainder is %v \n", remainder)
	if remainder > 0 {
		// fmt.Printf("The first half of the new arr is %v \n", arr[currIdx:])
		// fmt.Printf("The second half of the new arr is %v \n", arr[:remainder])
	} else {
		// fmt.Printf("The first half of the new arr is %v \n", arr[currIdx:currIdx+el])
		// fmt.Printf("The second half of the new arr is %v \n", arr[:remainder])
	}

	if remainder > 0 {
		newArr = append(newArr, arr[currIdx:]...)
		newArr = append(newArr, arr[:remainder]...)
	} else {
		newArr = append(newArr, arr[currIdx:currIdx+el]...)
	}
	//cleave off any extra
	if len(newArr) > len(arr) {
		newArr = newArr[:len(arr)]
	}

	for i, j := 0, len(newArr)-1; i < j; i, j = i+1, j-1 {
		newArr[i], newArr[j] = newArr[j], newArr[i]
	}
	// fmt.Printf("the reversed array is %v \n", newArr)

	if remainder > 0 {
		for i := currIdx; i < len(arr); i++ {
			reversedEl := newArr[0]
			newArr = newArr[1:]
			// fmt.Printf("Reversed el: %v, newArr: %v, index i: %v \n", reversedEl, newArr, i)
			arr[i] = reversedEl
		}

		for i := 0; i < remainder+1; i++ {
			if len(newArr) == 0 {
				continue
			}
			reversedEl := newArr[0]
			newArr = newArr[1:]
			// fmt.Printf("REMAINDER- Reversed el: %v, newArr: %v, index i: %v \n", reversedEl, newArr, i)
			arr[i] = reversedEl
		}
	} else {
		for i := currIdx; i < currIdx+el; i++ {
			reversedEl := newArr[0]
			newArr = newArr[1:]
			// fmt.Printf("Reversed el: %v, newArr: %v, index i: %v \n", reversedEl, newArr, i)
			arr[i] = reversedEl
		}
	}

	return arr
}

func knotHashAdvent(filepath string) int {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(string(content), ",")
	fmt.Printf("Split_arr: %v, len: %v", split, len(split))
	intArr := make([]int, len(split))
	for i, el := range split {
		intArr[i], err = strconv.Atoi(el)
		if err != nil {
			log.Fatal(err)
		}
	}
	return knotHash(intArr)
}
