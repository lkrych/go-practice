package goExercises

import "log"

//build a stack data structure
type stack struct {
	data []int
}

func (s *stack) push(el int) {
	s.data = append(s.data, el)
}

func (s *stack) pop() int {
	if s.isEmpty() {
		log.Fatal("There are no elements on the stack")
	}
	lastIdx := len(s.data) - 1
	popped := s.data[lastIdx]
	s.data = s.data[:lastIdx]
	return popped
}

func (s *stack) peek() int {
	return s.data[len(s.data)-1]
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

type Node struct {
	val  int
	next *Node
}

//write code to partition a linked list around a value x, such that all nodes less than
//x come before all nodes greater than or equal to x
func partitionLL(head *Node, partition int) *Node {
	beforeStart := &Node{}
	afterStart := &Node{}
	beforeEnd := &Node{}
	afterEnd := &Node{}
	currentNode := head
	for currentNode != nil {
		next := currentNode.next
		currentNode.next = nil
		if currentNode.val < partition {
			if beforeStart.val == 0 {
				beforeStart = currentNode
				beforeEnd = beforeStart
			} else {
				beforeEnd.next = currentNode
				beforeEnd = currentNode
			}
		} else { // if less than partiation
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
	beforeEnd.next = afterStart
	return beforeStart
}

//write code that removes duplicates from an unsorted linked list
func deleteDupsLL(head *Node) *Node {
	intMap := map[int]bool{}
	currentNode := head
	prev := &Node{}
	for currentNode != nil {
		if _, ok := intMap[currentNode.val]; ok {
			prev.next = currentNode.next
		} else {
			intMap[currentNode.val] = true
			prev = currentNode
		}
		currentNode = currentNode.next
	}
	return head
}
