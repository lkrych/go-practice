package goExercises

import (
	"fmt"
	"testing"
)

func TestCreateALinkedList(t *testing.T) {
	input1 := []int{1, 2, 3, 4, 5}
	input2 := []int{5, 6, 7, 8, 9, 10}
	input3 := []int{100, 101, 102}

	list1 := createALinkedList(input1)
	list2 := createALinkedList(input2)
	list3 := createALinkedList(input3)

	len1 := countLengthOfLL(list1)
	len2 := countLengthOfLL(list2)
	len3 := countLengthOfLL(list3)

	if len1 != 5 {
		t.Errorf("Expected createALinkedList to make a list of length 5 not %v", len1)
	}
	if !findValueInLL(list1, 2) {
		t.Errorf("Expected createALinkedList to contain 2")
	}
	if !findValueInLL(list1, 5) {
		t.Errorf("Expected createALinkedList to contain 5")
	}
	if findValueInLL(list1, 10) {
		t.Errorf("Expected createALinkedList to not contain 10 ")
	}

	if len2 != 6 {
		t.Errorf("Expected createALinkedList to make a list of length 6 not %v", len2)
	}
	if !findValueInLL(list2, 7) {
		t.Errorf("Expected createALinkedList to contain 7")
	}
	if !findValueInLL(list2, 9) {
		t.Errorf("Expected createALinkedList to contain 9")
	}
	if findValueInLL(list2, 101) {
		t.Errorf("Expected createALinkedList to not contain 101 ")
	}

	if len3 != 3 {
		t.Errorf("Expected createALinkedList to make a list of length 3 not %v", len3)
	}
	if !findValueInLL(list3, 101) {
		t.Errorf("Expected createALinkedList to contain 101")
	}
	if !findValueInLL(list3, 102) {
		t.Errorf("Expected createALinkedList to contain 102")
	}
	if findValueInLL(list1, 10) {
		t.Errorf("Expected createALinkedList to not contain 10 ")
	}

	//testing delete

	list4 := createALinkedList(input1)
	list5 := createALinkedList(input2)
	list6 := createALinkedList(input3)

	d4 := deleteANode(list4, 3)
	d5 := deleteANode(list5, 7)
	d6 := deleteANode(list6, 100)

	dlen4 := countLengthOfLL(d4)
	dlen5 := countLengthOfLL(d5)
	dlen6 := countLengthOfLL(d6)

	if dlen4 != 4 {
		t.Errorf("Expected deleteANode to make a list of length 4 not %v", dlen4)
		fmt.Println("Printing d4")
		printLL(d4)
	}
	if findValueInLL(d4, 3) {
		t.Errorf("Expected deleteANode(3) to delete nodes with the value of 3")
		fmt.Println("Printing d4")
		printLL(d4)
	}

	if dlen5 != 5 {
		t.Errorf("Expected deleteANode to make a list of length 5 not %v", dlen5)
		fmt.Println("Printing d5")
		printLL(d5)
	}
	if findValueInLL(d5, 7) {
		t.Errorf("Expected deleteANode(7) to delete nodes with the value of 7")
		fmt.Println("Printing d5")
		printLL(d5)
	}

	if dlen6 != 2 {
		t.Errorf("Expected deleteANode to make a list of length 2 not %v", dlen6)
		fmt.Println("Printing d6")
		printLL(d6)
	}
	if findValueInLL(d6, 100) {
		t.Errorf("Expected deleteANode(100) to delete nodes with the value of 100")
		fmt.Println("Printing d6")
		printLL(d6)
	}

}

//if you created your own function, comment this out!
// func createALinkedList(input []int) *Node {
// 	head := &Node{
// 		val:  input[0],
// 		next: nil,
// 	}
// 	currentNode := head
// 	for i := 1; i < len(input); i++ {
// 		nextNode := &Node{
// 			val:  input[i],
// 			next: nil,
// 		}
// 		currentNode.next = nextNode
// 		currentNode = nextNode
// 	}
// 	return head
// }

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

// func printLL(head *Node) {
// 	currentNode := head

// 	for currentNode.next != nil {
// 		fmt.Printf("  %v  ->", currentNode)
// 		currentNode = currentNode.next
// 	}
// 	fmt.Printf("  %v  \n", currentNode)
// }

func TestPartitionALinkedList(t *testing.T) {
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

func TestCheckPermutation(t *testing.T) {
	test1 := checkPermutation("blah", "halb")
	test2 := checkPermutation("baleen", "neelab")
	test3 := checkPermutation("cactus", "flower")
	test4 := checkPermutation("ballsy", "heptagon")
	test5 := checkPermutation("winky face", "kin wafeyc")

	if test1 != true {
		t.Errorf("Expected checkPermutation to detect that blah and halb are permutations of eachother")
	}

	if test2 != true {
		t.Errorf("Expected checkPermutation to detect that baleen and neelab are permutations of eachother")
	}

	if test3 != false {
		t.Errorf("Expected checkPermutation to detect that cactus and flower are not permutations of eachother")
	}

	if test4 != false {
		t.Errorf("Expected checkPermutation to detect that ballsy and heptagon are not permutations of eachother")
	}

	if test5 != true {
		t.Errorf("Expected checkPermutation to detect that winky face and kin wafeyc are permutations of eachother")
	}
}

func TestThirdGreatest(t *testing.T) {
	test1 := thirdGreatest([]int{5, 3, 7})
	test2 := thirdGreatest([]int{5, 3, 4, 7})
	test3 := thirdGreatest([]int{2, 3, 4, 7})

	if test1 != 3 {
		t.Errorf("Expected thirdGreatest of [5, 3, 7] to be 3, not %v", test1)
	}

	if test2 != 4 {
		t.Errorf("Expected thirdGreatest of [5, 3, 4, 7] to be 4, not %v", test2)
	}

	if test3 != 3 {
		t.Errorf("Expected thirdGreatest of [2, 3, 4, 7] to be 3, not %v", test3)
	}
}

func TestNthPrime(t *testing.T) {
	test1 := nthPrime(3)
	test2 := nthPrime(4)
	test3 := nthPrime(15)
	test4 := nthPrime(28)
	test5 := nthPrime(77)

	if test1 != 5 {
		t.Errorf("Expected nthPrime of 3 to be 5, not %v", test1)
	}

	if test2 != 7 {
		t.Errorf("Expected nthPrime of 4 to be 7, not %v", test2)
	}

	if test3 != 47 {
		t.Errorf("Expected nthPrime of 15 to be 47, not %v", test3)
	}

	if test4 != 107 {
		t.Errorf("Expected nthPrime of 28 to be 107, not %v", test4)
	}

	if test5 != 389 {
		t.Errorf("Expected nthPrime of 77 to be 389, not %v", test5)
	}
}
