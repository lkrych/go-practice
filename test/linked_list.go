package main

import (
	"fmt"
	"log"
)

func main() {
}

type Node struct {
	val  int
	next *Node
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

func deleteANode(head *Node, deleteVal int) *Node {
	if head.val == deleteVal { //in case the head is the value we want to delete!
		head = head.next
	}
	currentNode := head
	for currentNode.next != nil {
		if currentNode.next.val == deleteVal {
			currentNode.next = currentNode.next.next
		}
		currentNode = currentNode.next
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

func deleteDupsLL(head *Node) *Node {
	intMap := map[int]bool{}
	prev := &Node{}
	currentNode := head
	for currentNode.next != nil {
		if _, ok := intMap[currentNode.val]; ok {
			prev.next = currentNode.next
		} else {
			intMap[currentNode.val] = true
			prev = currentNode
		}
		currentNode = currentNode.next
	}
	if _, ok := intMap[currentNode.val]; ok {
		prev.next = currentNode.next
	}
	return head
}

// find the kth to last element of a singly-linked list recursively
// func findKthElement(head *Node, k int) (*Node, int) {
// 	if head.next == nil {
// 		return head, 1
// 	}
// 	var node, ith = findKthElement(head.next, k)
// 	ith++
// 	if ith == k {
// 		return head, ith
// 	}
// 	return node, ith
// }

func findKthElement(head *Node, k int) *Node {
	p1 := head
	p2 := head

	//move p1 k nodes into the list
	for i := 0; i < k; i++ {
		if p1 == nil {
			return nil
		}
		p1 = p1.next
	}

	//move them at the same pace. when p1 hits the end, p2 will be at the right ele
	for p1 != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p2
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

func deleteMiddleNode(n *Node) {
	if n == nil || n.next == nil {
		log.Fatal("this is not a middle node!")
	}
	n.val = n.next.val
	n.next = n.next.next
}

//given two numbers represented by linked lists, where each node contains a single digit, add them together
func sumLists(l1 *Node, l2 *Node) *Node {
	carry := 0
	sum := 0

	head := &Node{
		val:  0,
		next: nil,
	}
	currentNode := head

forLoop:
	for l1 != nil || l2 != nil {
		intermediate := 0
		switch {
		case l1 != nil && l2 != nil:
			intermediate = (l1.val + l2.val + carry)
			sum = intermediate % 10
			carry = intermediate / 10
			l1 = l1.next
			l2 = l2.next

		case l1 != nil:
			intermediate = (l1.val + carry)
			sum = intermediate % 10
			carry = intermediate / 10
			l1 = l1.next
		case l2 != nil:
			intermediate = (l2.val + carry)
			sum = intermediate % 10
			carry = intermediate / 10
			l2 = l2.next
		default:
			intermediate = carry
			sum = intermediate % 10
			carry = intermediate / 10
		}

		currentNode.val = sum
		nextNode := &Node{
			val:  0,
			next: nil,
		}
		//breaking logic
		switch {
		case l1 == nil && l2 == nil && carry == 0:
			break forLoop
		case l1 == nil && l2 == nil && carry != 0:
			currentNode.next = nextNode
			currentNode = nextNode
			currentNode.val = carry
			break forLoop
		}

		currentNode.next = nextNode
		currentNode = nextNode
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
