package goExercises

import (
	"fmt"
	"testing"
)

func TestCreatingAStack(t *testing.T) {
	s := &stack{
		data: []int{1, 2, 3, 4, 5, 6},
	}
	testPop := s.pop()
	testPop2 := s.pop()

	if len(s.data) != 4 {
		t.Errorf("Expected pop to return elements from the top of the stack: %v, popped: %v", s.data, testPop)
	}
	if s.peek() != 4 {
		t.Errorf("Expected peek to return the element at the top of the stack, not %v", s.peek())
	}
	if len(s.data) != 4 {
		t.Errorf("Expected peek to not delete elements from the top of the stack")
	}
	if s.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the stack has %v elements", len(s.data))
	}

	s.push(testPop)
	if len(s.data) != 5 {
		t.Errorf("Expected pop to return elements from the top of the stack: %v, popped: %v", s.data, testPop2)
	}
	if s.peek() != 6 {
		t.Errorf("Expected peek to return the element at the top of the stack, not %v", s.peek())
	}
	if len(s.data) != 5 {
		t.Errorf("Expected peek to not delete elements from the top of the stack")
	}
	if s.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the stack has %v elements\n", len(s.data))
	}

	s.push(testPop2)
	if len(s.data) != 6 {
		t.Errorf("Expected push to return elements to the top of the stack")
	}
	if s.peek() != 5 {
		t.Errorf("Expected peek to return the element at the top of the stack, not %v", s.peek())
	}
	if len(s.data) != 6 {
		t.Errorf("Expected peek to not delete elements from the top of the stack")
	}
	if s.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the stack has %v elements\n", len(s.data))
	}

	for i := 0; i < 6; i++ {
		s.pop()
	}
	if !s.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the stack should be empty")
	}
}

func TestCreateALinkedList(t *testing.T) {
	input1 := []int{10, 5, 8, 1, 3, 2, 4}
	input2 := []int{4, 7, 2, 3, 1}
	input3 := []int{105, 783, 98, 77, 33, 102}

	list1 := createALinkedList(input1)
	list2 := createALinkedList(input2)
	list3 := createALinkedList(input3)

	partition1 := partitionLL(list1, 5)
	partition2 := partitionLL(list2, 3)
	partition3 := partitionLL(list3, 105)

	if findIdxOfVal(partition1, 4) > 5 {
		t.Errorf("Expected partition to adjust the position of nodes")
		printLL(partition1)
	}

	if findIdxOfVal(partition1, 10) < 4 {
		t.Errorf("Expected partition to adjust the position of nodes")
		printLL(partition1)
	}

	if findIdxOfVal(partition2, 7) < 3 {
		t.Errorf("Expected partition to adjust the position of nodes")
		printLL(partition2)
	}

	if findIdxOfVal(partition1, 2) > 3 {
		t.Errorf("Expected partition to adjust the position of nodes")
		printLL(partition2)
	}

	if findIdxOfVal(partition3, 102) > 5 {
		t.Errorf("Expected partition to adjust the position of nodes")
		printLL(partition3)
	}

	if findIdxOfVal(partition3, 783) < 3 {
		t.Errorf("Expected partition to adjust the position of nodes")
		printLL(partition3)
	}

}

func findIdxOfVal(head *Node, search int) int {
	idx := 0
	for head.val != search {
		idx++
		head = head.next
	}
	return idx
}

func createALinkedList(input []int) *Node {
	head := &Node{
		val:  input[0],
		next: nil,
	}
	currentNode := head
	for i := 1; i < len(input); i++ {
		nextNode := &Node{
			val:  input[i],
			next: nil,
		}
		currentNode.next = nextNode
		currentNode = nextNode
	}
	return head
}

func printLL(head *Node) {
	currentNode := head

	for currentNode.next != nil {
		fmt.Printf("  %v  ->", currentNode)
		currentNode = currentNode.next
	}
	fmt.Printf("  %v  \n", currentNode)
}

func TestRemoveDuplicatesLL(t *testing.T) {
	input1 := []int{1, 2, 3, 3, 4, 4, 5}
	input2 := []int{5, 8, 6, 7, 8, 9, 7, 10}
	input3 := []int{100, 102, 101, 102, 101, 102}

	list1 := createALinkedList(input1)
	list2 := createALinkedList(input2)
	list3 := createALinkedList(input3)

	list1 = deleteDupsLL(list1)
	list2 = deleteDupsLL(list2)
	list3 = deleteDupsLL(list3)

	len1 := countLengthOfLL(list1)
	len2 := countLengthOfLL(list2)
	len3 := countLengthOfLL(list3)

	if len1 != 5 {
		t.Errorf("Expected deleteDups to make a list of length 5 not %v", len1)
		fmt.Println("Printing list1")
		printLL(list1)
	}

	if len2 != 6 {
		t.Errorf("Expected deleteDups to make a list of length 6 not %v", len2)
		fmt.Println("Printing list2")
		printLL(list2)
	}

	if len3 != 3 {
		t.Errorf("Expected deleteDups to make a list of length 3 not %v", len3)
		fmt.Println("Printing list3")
		printLL(list3)
	}
}

func countLengthOfLL(head *Node) int {
	count := 0
	currentNode := head
	for currentNode.next != nil {
		count++
		currentNode = currentNode.next
	}
	count++ //b/c you won't count the last one
	return count
}

func findValueInLL(head *Node, findVal int) bool {
	found := false
	currentNode := head
	for currentNode.next != nil {
		if currentNode.val == findVal {
			found = true
		}
		currentNode = currentNode.next
	}
	if currentNode.val == findVal { //catch the last node
		found = true
	}
	return found
}
