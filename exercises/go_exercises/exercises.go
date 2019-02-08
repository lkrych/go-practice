package goExercises

import ("strings"; "strconv";)
//return the most common character in a string with its count
func mostCommonLetter(sentence string) []string {
	charMap := make(map[string]int)
	for _, char := range strings.Split(sentence, "") {
		if count, ok := charMap[char]; ok {
			charMap[char] = count + 1
		} else {
			charMap[char] = 1
		}
	}
	greatestCount := 0
	greatestCountChar := ""
	for char, count := range charMap {
		if count > greatestCount {
			greatestCount = count
			greatestCountChar = char
		}
	}
	iToS := strconv.Itoa(greatestCount)
	return []string{greatestCountChar, iToS}
}
 
//given a linked list, delete nodes with the the value of x
//optional create a function that builds a linked list with an input array of integers
//if you don't want to do this, you can just look in the test file for this problem and grab one

func deleteANode(list *Node, x int) *Node {
	head := list
	if head.val == x {
		head = head.next
	}
	prevNode := head
	currentNode := head.next
	for currentNode.next != nil {
		if currentNode.val == x {
			prevNode.next = currentNode.next
		}
		prevNode = currentNode
		currentNode = currentNode.next
	}
	return head
}
OB
