package goExercises
 
import (
"fmt"
"testing"
"github.com/google/go-cmp/cmp"
)
 
import (
	"fmt"
	"testing"
)

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

//if you created your own function, comment this out!
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

func printLL(head *Node) {
	currentNode := head

	for currentNode.next != nil {
		fmt.Printf("  %v  ->", currentNode)
		currentNode = currentNode.next
	}
	fmt.Printf("  %v  \n", currentNode)
}

 
func TestScrambleString(t *testing.T) {
	test1 := scrambleString("abcd", []int{3, 1, 2, 0})
	test2 := scrambleString("markov", []int{5, 3, 1, 4, 2, 0})

	if test1 != "dbca" {
		t.Errorf("Expected scrambleString of 'abcd' with inputs [3,1,2,0] to be 'dcba', not %v", test1)
	}

	if test2 != "vkaorm" {
		t.Errorf("Expected scrambleString of 'markov' with inputs [5,3,1,4,2,0] to be 'vkaorm', not %v", test2)
	}
}
 
func TestPartitionEvenOdd(t *testing.T) {
	test1 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3})
	test2 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3, 1, 6, 10})
	test3 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3, 3, 3, 3})
	test4 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3, 100, 102, 6, 1, 10})
	test5 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3, 11, 10, 6, 1})
	ans1 := []int{2, 4, 8, 3, 5, 7, 9}
	ans2 := []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9}
	ans3 := []int{2, 4, 8, 3, 3, 3, 3, 5, 7, 9}
	ans4 := []int{2, 4, 6, 8, 10, 100, 102, 1, 3, 5, 7, 9}
	ans5 := []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9, 11}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans1, test1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans2, test2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans3, test3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans4, test4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans5, test5)
	}
}

 
func TestMultiplyArbitrary(t *testing.T) {
	test1 := multiplyArb([]int{1, 9, 3, 7, 0, 7, 7, 2, 1}, []int{-7, 6, 1, 8, 3, 8, 2, 5, 7, 2, 8, 7})
	test2 := multiplyArb([]int{1, 2, 9, 8, 2, 3, 4, 1, 3, 5, 0}, []int{4, 5})
	test3 := multiplyArb([]int{1, 2, 5, 3, 2, 4, 6, 7, 8, 2}, []int{-4, 0, 0, 2, 3})
	test4 := multiplyArb([]int{5, 3, 6, 3, 7, 9, 5, 0}, []int{3, 5, 2})
	test5 := multiplyArb([]int{1, 2}, []int{5})
	test6 := multiplyArb([]int{1, 2, 3}, []int{9, 8, 7})
	ans1 := []int{-1, 4, 7, 5, 7, 3, 9, 5, 2, 5, 8, 9, 6, 7, 6, 4, 1, 2, 9, 2, 7}
	ans2 := []int{5, 8, 4, 2, 0, 5, 3, 6, 0, 7, 5, 0}
	ans3 := []int{-5, 0, 1, 5, 8, 6, 9, 5, 9, 5, 5, 9, 8, 6}
	ans4 := []int{1, 8, 8, 8, 0, 5, 5, 8, 4, 0, 0}
	ans5 := []int{6, 0}
	ans6 := []int{1, 2, 1, 4, 0, 1}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans1, test1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans2, test2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans3, test3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans4, test4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans5, test5)
	}

	if !cmp.Equal(test6, ans6) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans6, test6)
	}
}

 
func TestIsPowerOfTwo(t *testing.T) {
	test1 := isPowerOfTwo(1)
	test2 := isPowerOfTwo(16)
	test3 := isPowerOfTwo(64)
	test4 := isPowerOfTwo(78)
	test5 := isPowerOfTwo(0)

	if test1 != true {
		t.Errorf("Expected isPowerOfTwo of 1 to be true, not %v", test1)
	}

	if test2 != true {
		t.Errorf("Expected isPowerOfTwo of 16 to be true, not %v", test2)
	}

	if test3 != true {
		t.Errorf("Expected isPowerOfTwo of 64 to be true, not %v", test3)
	}

	if test4 != false {
		t.Errorf("Expected isPowerOfTwo of 78 to be false, not %v", test4)
	}

	if test5 != false {
		t.Errorf("Expected isPowerOfTwo of 0 to be false, not %v", test5)
	}
}

 
