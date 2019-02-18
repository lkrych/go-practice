package goExercises
 
import (
"fmt"
"testing"
"github.com/google/go-cmp/cmp"
)
 
func TestRemoveDups(t *testing.T) {
	test1 := removeDups([]int{1, 1, 1, 2, 3, 4, 4, 4, 5})
	test2 := removeDups([]int{1, 1, 1, 2, 33, 102, 102})
	test3 := removeDups([]int{1, 1, 1, 2, 3, 3, 4, 5, 6, 7, 8, 8, 9, 9})
	test4 := removeDups([]int{1, 1, 1, 2, 3, 3, 3, 3, 3})
	test5 := removeDups([]int{1, 1, 1, 2, 3, 3, 4, 5, 5})
	ans1 := []int{1, 2, 3, 4, 5}
	ans2 := []int{1, 2, 33, 102}
	ans3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans4 := []int{1, 2, 3}
	ans5 := []int{1, 2, 3, 4, 5}
	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected array %v to be %v", test1, ans1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected array %v to be %v", test2, ans2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected array %v to be %v", test4, ans3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected array %v to be %v", test5, ans4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected array %v to be %v", test5, ans5)
	}
}
 

func TestDeleteMiddleNode(t *testing.T) {
	input1 := []int{1, 2, 3, 4, 5}
	input2 := []int{5, 6, 7, 8, 9, 10}
	input3 := []int{100, 101, 102}

	list1 := createALinkedList(input1)
	list2 := createALinkedList(input2)
	list3 := createALinkedList(input3)

	middleNode1 := findNodeInLL(list1, 3)
	middleNode2 := findNodeInLL(list2, 7)
	middleNode3 := findNodeInLL(list3, 101)

	if middleNode1.val != 3 {
		t.Errorf("Expected findNode to grab the right node!")
	}

	if middleNode2.val != 7 {
		t.Errorf("Expected findNode to grab the right node!")
	}

	if middleNode3.val != 101 {
		t.Errorf("Expected findNode to grab the right node!")
	}

	deleteMiddleNode(middleNode1)
	deleteMiddleNode(middleNode2)
	deleteMiddleNode(middleNode3)

	len1 := countLengthOfLL(list1)
	len2 := countLengthOfLL(list2)
	len3 := countLengthOfLL(list3)

	if len1 != 4 {
		t.Errorf("Expected deleteMiddleNode to make a list of length 4 not %v", len1)
		fmt.Println("Printing d4")
		printLL(list1)
	}
	if findValueInLL(list1, 3) {
		t.Errorf("Expected deleteMiddleNode to delete nodes with the value of 3")
		fmt.Println("Printing list1")
		printLL(list1)
	}

	if len2 != 5 {
		t.Errorf("Expected deleteMiddleNode to make a list of length 5 not %v", len2)
		fmt.Println("Printing list2")
		printLL(list2)
	}
	if findValueInLL(list2, 7) {
		t.Errorf("Expected deleteMiddleNode to delete nodes with the value of 7")
		fmt.Println("Printing list2")
		printLL(list2)
	}

	if len3 != 2 {
		t.Errorf("Expected deleteMiddleNode to make a list of length 2 not %v", len3)
		fmt.Println("Printing list3")
		printLL(list3)
	}
	if findValueInLL(list3, 101) {
		t.Errorf("Expected deleteMiddleNode to delete nodes with the value of 101")
		fmt.Println("Printing list3")
		printLL(list3)
	}

}

func findNodeInLL(head *Node, val int) *Node {
	for head.next != nil {
		if head.val == val {
			return head
		}
		head = head.next
	}
	return nil
}

func findValueInLL(head *Node, val int) bool {
	for head.next != nil {
		if head.val == val {
			return true
		}
		head = head.next
	}
	return false
}

func countLengthOfLL(head *Node) int {
	count := 0
	for head.next != nil {
		count++
		head = head.next
	}
	return count
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
 
