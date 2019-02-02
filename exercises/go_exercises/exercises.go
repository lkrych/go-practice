package goExercises

import ("strings"; "strconv")

//return the most common character in a string with its count
func mostCommonLetter(sentence string) []string {
	count := make(map[string]int)
	highestCount := 0
	var mostCommonChar string
	for _, char := range strings.Split(sentence, ""){
		if _, found := count[char]; found {
			count[char] = count[char] + 1
			if count[char] + 1 > highestCount {
				highestCount = count[char]
				mostCommonChar = char
			}
		} else {
			count[char] = 1
		}
	}
	countString := strconv.Itoa(highestCount)
	return []string{mostCommonChar, countString}
}
 
//you have two numbers represented by a linked list, where each node contains a single digit
//the digits are stored in reverse order, such that 1's digit is at the head of the list
//write a function that adds the two numbers and returns the sum as a linked list
//example:
//(7 -> 1 -> 6) + (5 -> 9 -> 2) = (2 -> 1 -> 9)
func sumLists(h1, h2 *Node) *Node {
	head, carry := addTwoNodes(h1,h2, 0)
	currentNode := head
	for h1.next != nil && h2.next != nil {
		h1 = h1.next
		h2 = h2.next
		nextNode, newCarry := addTwoNodes(h1, h2, carry)
		carry = newCarry
		currentNode.next = nextNode
		currentNode = nextNode
	}
	if carry != 0 {
		currentNode.next = &Node{
			val: carry,
			next: nil,
		}
	}
	return head
}

func addTwoNodes(n1, n2 *Node, carry int) (*Node, int) {
	//addTwoNodes return new node and remainder
	newVal := n1.val + n2.val + carry
	newCarry := newVal / 10
	newNode :=  &Node{
		val: newVal % 10,
		next: nil,
	}

	return newNode, newCarry
}

type Node struct {
	val int
	next *Node
}
 
