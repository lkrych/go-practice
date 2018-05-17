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

func knotHash(input []int, knot []int, currIdx int, skipSize int) ([]int, int, int) {
	for _, el := range input {
		knot = reverseAndWrap(knot, el, currIdx%len(knot))
		currIdx += el + skipSize
		skipSize++
		fmt.Printf("The currentIndex is %v \n", currIdx)
		fmt.Printf("The current skip size is %v \n", skipSize)
		fmt.Printf("####################### \n")
		fmt.Printf("The knot is %v \n", knot)
		fmt.Printf("####################### \n")
	}
	return knot, currIdx, skipSize
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
	knotInput := make([]int, 256)
	currIdx := 0
	skipSize := 0
	knot, _, _ := knotHash(intArr, knotInput, currIdx, skipSize)
	return knot[0] * knot[1]
}

func knotHashAdventAdvanced(filepath string) string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	content = append(content, 17, 31, 73, 47, 23)
	intArr := make([]int, len(content))
	for i, element := range content {
		intArr[i] = int(element)
	}
	knotInput := make([]int, 256)
	for i := 0; i < len(knotInput); i++ {
		knotInput[i] = i
	}
	knot := computeSparseHash(intArr, knotInput)
	dense := computeDenseHash(knot)
	return convertHex(dense)
}

func computeSparseHash(inputArr []int, knotInput []int) []int {
	currPos := 0
	skipSize := 0
	for i := 0; i < 64; i++ {
		knotInput, currPos, skipSize = knotHash(inputArr, knotInput, currPos, skipSize)
	}
	return knotInput
}

func computeDenseHash(knot []int) []int {
	hexArr := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i, el := range knot {
		group := i / 16
		mod := i % 16
		if mod == 0 {
			hexArr[group] = el
		} else {
			hexArr[group] = hexArr[group] ^ el
		}
	}
	return hexArr
}

func convertHex(denseHash []int) string {
	hexStr := ""
	for _, el := range denseHash {
		hexStr += fmt.Sprintf("%.02x", el)
	}
	return hexStr
}

func hexMaze(mazeInput []string) int {
	dirMap := make(map[string]int)
	farthestDistance := 0
	currentDistance := 0
	for _, direction := range mazeInput {
		dirMap = addDistance(direction, dirMap)
		currentDistance, dirMap = cleanUp(dirMap)
		if currentDistance > farthestDistance {
			farthestDistance = currentDistance
		}
	}
	fmt.Printf("The farthest distance gone by the wandering child process is %v \n", farthestDistance)
	return currentDistance
}

func addDistance(dir string, dirMap map[string]int) map[string]int {
	switch dir {
	case "n":
		if val, ok := dirMap["s"]; ok {
			if val > 0 {
				dirMap["s"]--
				break
			}
		} else if val, ok := dirMap["se"]; ok {
			if val > 0 {
				dirMap["se"]--
				dirMap["e"]++
				break
			}
		} else if val, ok := dirMap["sw"]; ok {
			if val > 0 {
				dirMap["sw"]--
				dirMap["w"]++
				break
			}
		}
		dirMap["n"]++
	case "s":
		if val, ok := dirMap["n"]; ok {
			if val > 0 {
				dirMap["n"]--
				break
			}
		} else if val, ok := dirMap["ne"]; ok {
			if val > 0 {
				dirMap["ne"]--
				dirMap["e"]++
				break
			}
		} else if val, ok := dirMap["nw"]; ok {
			if val > 0 {
				dirMap["nw"]--
				dirMap["w"]++
				break
			}
		}
		dirMap["s"]++

	case "nw":
		if val, ok := dirMap["ne"]; ok {
			if val > 0 {
				dirMap["ne"]--
				dirMap["n"]++
				break
			}
		}

		if val, ok := dirMap["se"]; ok {
			if val > 0 {
				dirMap["se"]--
				break
			}
		}

		dirMap["nw"]++

	case "ne":
		if val, ok := dirMap["nw"]; ok {
			if val > 0 {
				dirMap["nw"]--
				dirMap["n"]++
				break
			}
		}

		if val, ok := dirMap["sw"]; ok {
			if val > 0 {
				dirMap["sw"]--
				break
			}
		}

		dirMap["ne"]++

	case "sw":
		if val, ok := dirMap["se"]; ok {
			if val > 0 {
				dirMap["se"]--
				dirMap["s"]++
				break
			}
		}

		if val, ok := dirMap["ne"]; ok {
			if val > 0 {
				dirMap["ne"]--
				break
			}
		}

		dirMap["sw"]++
	case "se":
		if val, ok := dirMap["sw"]; ok {
			if val > 0 {
				dirMap["sw"]--
				dirMap["s"]++
				break
			}
		}

		if val, ok := dirMap["nw"]; ok {
			if val > 0 {
				dirMap["nw"]--
				break
			}
		}

		dirMap["se"]++
	}
	fmt.Printf("Direction map: %v \n", dirMap)
	return dirMap
}

func hexMazeAdvent(filepath string) int {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	strArray := strings.Split(string(content), ",")

	fmt.Printf("The converted array is %v \n", strArray)
	return hexMaze(strArray)
}

func cleanUp(dirMap map[string]int) (int, map[string]int) {
	//Combine S and NE/NW
	if valSouth, ok := dirMap["s"]; ok {
		if valSouth > 0 {
			if valNorth, ok := dirMap["ne"]; ok {
				if valNorth > 0 {
					if valSouth-valNorth > 0 {
						dirMap["s"] = valSouth - valNorth
						dirMap["ne"] = 0
						dirMap["se"] = dirMap["se"] + valNorth
					} else {
						dirMap["se"] = dirMap["se"] + valSouth
						dirMap["s"] = 0
						dirMap["ne"] = valNorth - valSouth
					}
				}
			}

			if valNorth, ok := dirMap["nw"]; ok {
				if valNorth > 0 {
					if valSouth-valNorth > 0 {
						dirMap["s"] = valSouth - valNorth
						dirMap["nw"] = 0
						dirMap["sw"] = dirMap["sw"] + valNorth
					} else {
						dirMap["sw"] = dirMap["sw"] + valSouth
						dirMap["s"] = 0
						dirMap["nw"] = valNorth - valSouth
					}
				}
			}
		}
	}

	if valNorth, ok := dirMap["n"]; ok {
		if valNorth > 0 {
			if valSouth, ok := dirMap["se"]; ok {
				if valSouth > 0 {
					if valNorth-valSouth > 0 {
						dirMap["n"] = valNorth - valSouth
						dirMap["se"] = 0
						dirMap["ne"] = dirMap["ne"] + valSouth
					} else {
						dirMap["ne"] = dirMap["ne"] + valNorth
						dirMap["n"] = 0
						dirMap["se"] = valSouth - valNorth
					}
				}
			}
			//combine N and SE/SW
			if valSouth, ok := dirMap["sw"]; ok {
				if valSouth > 0 {
					if valNorth-valSouth > 0 {
						dirMap["n"] = valNorth - valSouth
						dirMap["sw"] = 0
						dirMap["nw"] = dirMap["nw"] + valSouth
					} else {
						dirMap["nw"] = dirMap["nw"] + valNorth
						dirMap["n"] = 0
						dirMap["sw"] = valSouth - valNorth
					}
				}
			}
		}
	}

	if valNorth, ok := dirMap["n"]; ok {
		if valNorth > 0 {
			if valSouth, ok := dirMap["s"]; ok {
				if valNorth-valSouth > 0 {
					dirMap["n"] = valNorth - valSouth
					dirMap["s"] = 0
				} else {
					dirMap["s"] = valSouth - valNorth
					dirMap["n"] = 0
				}
			}
		}
	}
	fmt.Printf("Cleaned up Map: %v \n", dirMap)
	totalDistance := 0
	for _, distance := range dirMap {
		totalDistance += distance
	}
	return totalDistance, dirMap

}

func digitalPlumber(filepath string) int {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	nodeTable := buildNodeTree(content)
	preventLoops := make(map[string]int)
	count := searchTree(nodeTable, "0", preventLoops)

	return count
}

func buildNodeTree(content []byte) map[string]*Node {
	nodeTable := make(map[string]*Node)
	splitByLine := strings.Split(string(content), "\n")
	for _, line := range splitByLine {
		splitLine := strings.Split(line, " ")
		fmt.Printf("Split line: %v, len: %v \n", splitLine, len(splitLine))
		pipeIn := splitLine[0]
		pipeOut := splitLine[2:]
		fmt.Printf("Pipe in: %v \n", pipeIn)
		fmt.Printf("Pipe out: %v \n", pipeOut)
		//populate map
		entry := splitLine[0]
		newNode := Node{
			name:     entry,
			val:      convertToInt(splitLine[1]),
			children: []string{},
		}
		for _, el := range pipeOut {
			newNode.children = append(newNode.children, el)
		}

		// fmt.Printf("New node: %v \n", newNode)
		nodeTable[entry] = &newNode
	}
	fmt.Printf("nodeTable: %v \n", nodeTable)

	return nodeTable
}

func searchTree(nodeMap map[string]*Node, node string, preventLoopsMap map[string]int) int {
	fmt.Printf("Looking for %v in %v, the preventLoopsMap looks like %v \n", node, nodeMap, preventLoopsMap)

	count := 0

	if _, ok := preventLoopsMap[node]; ok {

		fmt.Printf("%v was found in the preventLoops map \n", node)
		return count
	}
	//add currentNode to preventLoopsMap to prevent infinite loops
	preventLoopsMap[node] = 1

	for _, child := range nodeMap[node].children {
		fmt.Printf("The child is %v \n", child)
		count += searchTree(nodeMap, child, preventLoopsMap)
		count++
		fmt.Printf("Made it through the loop with child %v \n", child)
	}
	return count
}

//Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := createNumber(l1)
	n2 := createNumber(l2)

	sum := n1 + n2

	fmt.Printf("%d + %d = %d \n", n1, n2, sum)

	return createLinkedList(int64(sum))
}

func createNumber(l *ListNode) int {
	val := l.Val
	decimalCount := 1
	for l.Next != nil { //iterate until nil
		l = l.Next
		computation := (l.Val * power(10, decimalCount))
		val += computation
		decimalCount++
	}
	return val
}

func createLinkedList(n int64) *ListNode {
	head := ListNode{
		Val:  int(n % 10),
		Next: nil,
	}

	var currentNode *ListNode
	currentNode = &head

	for n > 0 {
		n = n / 10
		nextNode := ListNode{
			Val:  int(n % 10),
			Next: nil,
		}
		if n == 0 {
			break
		}
		currentNode.Next = &nextNode
		currentNode = &nextNode
	}
	return &head
}

func printLinkedList(l *ListNode) {
	i := 1
	for l != nil { //iterate until nil
		fmt.Printf("%v: %v \n", i, l.Val)
		i++
		l = l.Next
	}
}

func power(base int, power int) int {
	val := 1
	for i := 0; i < power; i++ {
		val *= base
	}
	return val
}

func optimizedAddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	sum := 0

	head := ListNode{
		Val:  0,
		Next: nil,
	}

	var currentNode *ListNode
	currentNode = &head

forLoop:
	for l1 != nil || l2 != nil {
		intermediate := 0
		switch {
		case l1 != nil && l2 != nil:
			intermediate = (l1.Val + l2.Val + carry)
			sum = intermediate % 10
			carry = intermediate / 10
			l1 = l1.Next
			l2 = l2.Next

		case l1 != nil:
			intermediate = (l1.Val + carry)
			sum = intermediate % 10
			carry = intermediate / 10
			l1 = l1.Next
		case l2 != nil:
			intermediate = (l2.Val + carry)
			sum = intermediate % 10
			carry = intermediate / 10
			l2 = l2.Next
		default:
			intermediate = carry
			sum = intermediate % 10
			carry = intermediate / 10
		}

		fmt.Println("Sum is", sum)
		fmt.Println("Carry is", carry)
		fmt.Println("l1 is", l1)
		fmt.Println("l2 is", l2)
		currentNode.Val = sum
		nextNode := ListNode{
			Val:  0,
			Next: nil,
		}
		//breaking logic
		switch {
		case l1 == nil && l2 == nil && carry == 0:
			fmt.Println("case 1")
			break forLoop
		case l1 == nil && l2 == nil && carry != 0:
			fmt.Println("case 2")
			currentNode.Next = &nextNode
			currentNode = &nextNode
			currentNode.Val = carry
			break forLoop
		}

		currentNode.Next = &nextNode
		currentNode = &nextNode
	}

	return &head

}
