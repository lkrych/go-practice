package goExercises

import "log"
 
//In O(n) time, and using O(1) space, remove duplicates from a sorted array
func removeDups(arr []int) []int {
	writeIdx := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[writeIdx -1] {
			arr[writeIdx] = arr[i]
			writeIdx++
		}
	}
	return arr[:writeIdx]
}

type Node struct {
	val int
	next *Node
}
 
//implement an algorithm to delete a node (any node but the first and last node)
// of a singly-linked list, given only access to that list
func deleteMiddleNode(n *Node) {
	if n == nil || n.next == nil {
		log.Fatal("This is not a middle node!")
	}
	n.val = n.next.val
	n.next = n.next.next
}
 
