package goExercises

import (
	"strings"
)

type Node struct {
	val  int
	next *Node
}

//given a linked list, delete nodes with the the value of x
//optional create a function that builds a linked list with an input array of integers
//if you don't want to do this, you can just look in the test file for this problem and grab one
func deleteANode(head *Node, delete int) *Node {
	lastNode := &Node{}
	if head.val == delete {
		head = head.next
	}

	currentNode := head.next
	for currentNode != nil {
		if currentNode.val == delete {
			lastNode.next = currentNode.next
		}
		lastNode = currentNode
		currentNode = currentNode.next
	}
	return head
}

//write code to partition a linked list around a value x, such that all nodes less than
//x come before all nodes greater than or equal to x
func partitionLL(head *Node, partition int) *Node {
	beforeStart := &Node{}
	beforeEnd := &Node{}
	afterStart := &Node{}
	afterEnd := &Node{}

	//partition
	currentNode := head
	for currentNode != nil {
		next := currentNode.next
		currentNode.next = nil
		if currentNode.val < partition {
			//insert node into end of before list
			if beforeStart.val == 0 {
				beforeStart = currentNode
				beforeEnd = beforeStart
			} else {
				beforeEnd.next = currentNode
				beforeEnd = currentNode
			}
		} else {
			//insert node into end of after list
			if afterStart.val == 0 {
				afterStart = currentNode
				afterEnd = afterStart
			} else {
				afterEnd.next = currentNode
				afterEnd = currentNode
			}
		}
		currentNode = next
	}
	if beforeStart.val == 0 {
		return afterStart
	}

	//merge the two lists
	beforeEnd.next = afterStart
	return beforeStart
}

//Given two strings, write a method to decide if one is a permutation of the other
func checkPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	charMap1 := map[string]int{}
	for _, ch := range strings.Split(s1, "") {
		if count, ok := charMap1[ch]; ok {
			charMap1[ch] = count + 1
		} else {
			charMap1[ch] = 1
		}
	}

	charMap2 := map[string]int{}
	for _, ch := range strings.Split(s2, "") {
		if count, ok := charMap2[ch]; ok {
			charMap2[ch] = count + 1
		} else {
			charMap2[ch] = 1
		}
	}

	for k, v := range charMap1 {
		count, _ := charMap2[k]
		if count != v {
			return false
		}
	}
	return true
}

//return the third greatest el in arr
func thirdGreatest(arr []int) int {
	first := 0
	second := 0
	third := 0

	for _, el := range arr {
		if el > first {
			third = second
			second = first
			first = el
		} else if el > second {
			third = second
			second = el
		} else if el > third {
			third = el
		}
	}
	return third
}

//return the nth prime
func nthPrime(nth int) int {
	primeCount := 0
	checkIfPrime := 2
	for primeCount < nth {
		if isPrime(checkIfPrime) {
			primeCount++
		}
		checkIfPrime++
	}
	return checkIfPrime - 1
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
