package main

import "testing"

// func TestCreateALinkedList(t *testing.T) {
// 	input1 := []int{1, 2, 3, 4, 5}
// 	input2 := []int{5, 6, 7, 8, 9, 10}
// 	input3 := []int{100, 101, 102}

// 	list1 := createALinkedList(input1)
// 	list2 := createALinkedList(input2)
// 	list3 := createALinkedList(input3)

// 	len1 := countLengthOfLL(list1)
// 	len2 := countLengthOfLL(list2)
// 	len3 := countLengthOfLL(list3)

// 	if len1 != 5 {
// 		t.Errorf("Expected createALinkedList to make a list of length 5 not %v", len1)
// 	}
// 	if !findValueInLL(list1, 2) {
// 		t.Errorf("Expected createALinkedList to contain 2")
// 	}
// 	if !findValueInLL(list1, 5) {
// 		t.Errorf("Expected createALinkedList to contain 5")
// 	}
// 	if findValueInLL(list1, 10) {
// 		t.Errorf("Expected createALinkedList to not contain 10 ")
// 	}

// 	if len2 != 6 {
// 		t.Errorf("Expected createALinkedList to make a list of length 6 not %v", len2)
// 	}
// 	if !findValueInLL(list2, 7) {
// 		t.Errorf("Expected createALinkedList to contain 7")
// 	}
// 	if !findValueInLL(list2, 9) {
// 		t.Errorf("Expected createALinkedList to contain 9")
// 	}
// 	if findValueInLL(list2, 101) {
// 		t.Errorf("Expected createALinkedList to not contain 101 ")
// 	}

// 	if len3 != 3 {
// 		t.Errorf("Expected createALinkedList to make a list of length 3 not %v", len3)
// 	}
// 	if !findValueInLL(list3, 101) {
// 		t.Errorf("Expected createALinkedList to contain 101")
// 	}
// 	if !findValueInLL(list3, 102) {
// 		t.Errorf("Expected createALinkedList to contain 102")
// 	}
// 	if findValueInLL(list1, 10) {
// 		t.Errorf("Expected createALinkedList to not contain 10 ")
// 	}

// 	//testing delete

// 	list4 := createALinkedList(input1)
// 	list5 := createALinkedList(input2)
// 	list6 := createALinkedList(input3)

// 	d4 := deleteANode(list4, 3)
// 	d5 := deleteANode(list5, 7)
// 	d6 := deleteANode(list6, 100)

// 	dlen4 := countLengthOfLL(d4)
// 	dlen5 := countLengthOfLL(d5)
// 	dlen6 := countLengthOfLL(d6)

// 	if dlen4 != 4 {
// 		t.Errorf("Expected deleteANode to make a list of length 4 not %v", dlen4)
// 		fmt.Println("Printing d4")
// 		printLL(d4)
// 	}
// 	if findValueInLL(d4, 3) {
// 		t.Errorf("Expected deleteANode(3) to delete nodes with the value of 3")
// 		fmt.Println("Printing d4")
// 		printLL(d4)
// 	}

// 	if dlen5 != 5 {
// 		t.Errorf("Expected deleteANode to make a list of length 5 not %v", dlen5)
// 		fmt.Println("Printing d5")
// 		printLL(d5)
// 	}
// 	if findValueInLL(d5, 7) {
// 		t.Errorf("Expected deleteANode(7) to delete nodes with the value of 7")
// 		fmt.Println("Printing d5")
// 		printLL(d5)
// 	}

// 	if dlen6 != 2 {
// 		t.Errorf("Expected deleteANode to make a list of length 2 not %v", dlen6)
// 		fmt.Println("Printing d6")
// 		printLL(d6)
// 	}
// 	if findValueInLL(d6, 100) {
// 		t.Errorf("Expected deleteANode(100) to delete nodes with the value of 100")
// 		fmt.Println("Printing d6")
// 		printLL(d6)
// 	}

// }

// func TestDeleteDuplicatesLL(t *testing.T) {
// 	input1 := []int{1, 2, 3, 3, 4, 4, 5}
// 	input2 := []int{5, 8, 6, 7, 8, 9, 7, 10}
// 	input3 := []int{100, 102, 101, 102, 101, 102}

// 	list1 := createALinkedList(input1)
// 	list2 := createALinkedList(input2)
// 	list3 := createALinkedList(input3)

// 	list1 = deleteDupsLL(list1)
// 	list2 = deleteDupsLL(list2)
// 	list3 = deleteDupsLL(list3)

// 	len1 := countLengthOfLL(list1)
// 	len2 := countLengthOfLL(list2)
// 	len3 := countLengthOfLL(list3)

// 	if len1 != 5 {
// 		t.Errorf("Expected deleteDups to make a list of length 5 not %v", len1)
// 		fmt.Println("Printing list1")
// 		printLL(list1)
// 	}

// 	if len2 != 6 {
// 		t.Errorf("Expected deleteDups to make a list of length 6 not %v", len2)
// 		fmt.Println("Printing list2")
// 		printLL(list2)
// 	}

// 	if len3 != 3 {
// 		t.Errorf("Expected deleteDups to make a list of length 3 not %v", len3)
// 		fmt.Println("Printing list3")
// 		printLL(list3)
// 	}
// }

// func TestFindKthElement(t *testing.T) {
// 	input1 := []int{1, 2, 3, 3, 4, 4, 5}
// 	input2 := []int{5, 8, 6, 7, 8, 9, 7, 10}
// 	input3 := []int{100, 102, 101, 102, 101, 102}

// 	list1 := createALinkedList(input1)
// 	list2 := createALinkedList(input2)
// 	list3 := createALinkedList(input3)

// 	list1Val := findKthElement(list1, 2)
// 	list2Val := findKthElement(list2, 5)
// 	list3Val := findKthElement(list3, 3)

// 	if list1Val.val != 4 {
// 		t.Errorf("Expected kthElement to find the kth element 4 not %v", list1Val.val)
// 		fmt.Println("Printing list1")
// 		printLL(list1)
// 	}

// 	if list2Val.val != 7 {
// 		t.Errorf("Expected kthElement to find the kth element 7 not %v", list2Val.val)
// 		fmt.Println("Printing list2")
// 		printLL(list2)
// 	}

// 	if list3Val.val != 102 {
// 		t.Errorf("Expected kthElement to find the kth element 102 not %v", list3Val.val)
// 		fmt.Println("Printing list3")
// 		printLL(list3)
// 	}
// }

// func TestDeleteMiddleNode(t *testing.T) {
// 	input1 := []int{1, 2, 3, 4, 5}
// 	input2 := []int{5, 6, 7, 8, 9, 10}
// 	input3 := []int{100, 101, 102}

// 	list1 := createALinkedList(input1)
// 	list2 := createALinkedList(input2)
// 	list3 := createALinkedList(input3)

// 	middleNode1 := findNodeInLL(list1, 3)
// 	middleNode2 := findNodeInLL(list2, 7)
// 	middleNode3 := findNodeInLL(list3, 101)

// 	if middleNode1.val != 3 {
// 		t.Errorf("Expected findNode to grab the right node!")
// 	}

// 	if middleNode2.val != 7 {
// 		t.Errorf("Expected findNode to grab the right node!")
// 	}

// 	if middleNode3.val != 101 {
// 		t.Errorf("Expected findNode to grab the right node!")
// 	}

// 	deleteMiddleNode(middleNode1)
// 	deleteMiddleNode(middleNode2)
// 	deleteMiddleNode(middleNode3)

// 	len1 := countLengthOfLL(list1)
// 	len2 := countLengthOfLL(list2)
// 	len3 := countLengthOfLL(list3)

// 	if len1 != 4 {
// 		t.Errorf("Expected deleteMiddleNode to make a list of length 4 not %v", len1)
// 		fmt.Println("Printing d4")
// 		printLL(list1)
// 	}
// 	if findValueInLL(list1, 3) {
// 		t.Errorf("Expected deleteMiddleNode to delete nodes with the value of 3")
// 		fmt.Println("Printing list1")
// 		printLL(list1)
// 	}

// 	if len2 != 5 {
// 		t.Errorf("Expected deleteMiddleNode to make a list of length 5 not %v", len2)
// 		fmt.Println("Printing list2")
// 		printLL(list2)
// 	}
// 	if findValueInLL(list2, 7) {
// 		t.Errorf("Expected deleteMiddleNode to delete nodes with the value of 7")
// 		fmt.Println("Printing list2")
// 		printLL(list2)
// 	}

// 	if len3 != 2 {
// 		t.Errorf("Expected deleteMiddleNode to make a list of length 2 not %v", len3)
// 		fmt.Println("Printing list3")
// 		printLL(list3)
// 	}
// 	if findValueInLL(list3, 101) {
// 		t.Errorf("Expected deleteMiddleNode to delete nodes with the value of 101")
// 		fmt.Println("Printing list3")
// 		printLL(list3)
// 	}

// }

// func TestSumLists(t *testing.T) {
// 	i1 := []int{1, 2, 3}
// 	i2 := []int{5, 8, 6}
// 	sum1 := []int{6, 0, 0, 1}

// 	i3 := []int{3, 5}
// 	i4 := []int{1, 8}
// 	sum2 := []int{4, 3, 1}

// 	i5 := []int{6, 4, 5}
// 	i6 := []int{1, 1, 1}
// 	sum3 := []int{7, 5, 6}

// 	list1 := createALinkedList(i1)
// 	list2 := createALinkedList(i2)
// 	answer1 := listToArray(createALinkedList(sum1))

// 	list3 := createALinkedList(i3)
// 	list4 := createALinkedList(i4)
// 	answer2 := listToArray(createALinkedList(sum2))

// 	list5 := createALinkedList(i5)
// 	list6 := createALinkedList(i6)
// 	answer3 := listToArray(createALinkedList(sum3))

// 	returned1 := listToArray(sumLists(list1, list2))
// 	returned2 := listToArray(sumLists(list3, list4))
// 	returned3 := listToArray(sumLists(list5, list6))

// 	if !cmp.Equal(answer1, returned1) {
// 		t.Errorf("Expected sumLists of %v and %v to be %v not %v", list1, list2, answer1, returned1)

// 	}

// 	if !cmp.Equal(answer2, returned2) {
// 		t.Errorf("Expected sumLists of %v and %v to be %v not %v", list3, list4, answer2, returned2)
// 	}

// 	if !cmp.Equal(answer3, returned3) {
// 		t.Errorf("Expected sumLists of %v and %v to be %v not %v", list5, list6, answer3, returned3)
// 	}
// }

// func listToArray(head *Node) []int {
// 	vals := []int{}
// 	for head != nil {
// 		vals = append(vals, head.val)
// 		head = head.next
// 	}
// 	return vals
// }

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
