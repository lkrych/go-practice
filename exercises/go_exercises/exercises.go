package goExercises
 
//In O(n) time, and using O(1) space, remove duplicates from a sorted array
func removeDups(arr []int) []int {
	writeIdx := 1
	for i := 1; i < len(arr); i++ {
		if arr[writeIdx - 1] != arr[i] {
			// keep track of where the non-duplicated characters are with writeIdx
			arr[writeIdx] = arr[i]
			writeIdx++
		}
	}
	return arr[:writeIdx]
} 

type Node struct {
	next *Node
	val int
}
 
//find the kth to last element of a singly-linked list
func findKthElement(head *Node, k int) (*Node, int) {
	if head.next == nil {
		return head, 1
	}

	node, ith := findKthElement(head.next, k)
	ith++
	if ith == k {
		return head, ith
	}
	return node, ith
}
 
