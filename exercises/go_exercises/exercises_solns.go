package goExercises

// type Node struct {
// 	val  int
// 	next *Node
// }

// //if you created your own function, comment this out!
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

// func deleteANode(head *Node, deleteVal int) *Node {
// 	if head.val == deleteVal { //in case the head is the value we want to delete!
// 		head = head.next
// 	}
// 	currentNode := head
// 	for currentNode.next != nil {
// 		if currentNode.next.val == deleteVal {
// 			currentNode.next = currentNode.next.next
// 		}
// 		currentNode = currentNode.next
// 	}

// 	return head
// }

// func partitionLL(head *Node, partition int) *Node {
// 	beforeStart := &Node{}
// 	beforeEnd := &Node{}
// 	afterStart := &Node{}
// 	afterEnd := &Node{}

// 	//partition
// 	currentNode := head
// 	for currentNode != nil {
// 		next := currentNode.next
// 		currentNode.next = nil
// 		if currentNode.val < partition {
// 			//insert node into end of before list
// 			if beforeStart.val == 0 {
// 				beforeStart = currentNode
// 				beforeEnd = beforeStart
// 			} else {
// 				beforeEnd.next = currentNode
// 				beforeEnd = currentNode
// 			}
// 		} else {
// 			//insert node into end of after list
// 			if afterStart.val == 0 {
// 				afterStart = currentNode
// 				afterEnd = afterStart
// 			} else {
// 				afterEnd.next = currentNode
// 				afterEnd = currentNode
// 			}
// 		}
// 		currentNode = next
// 	}
// 	if beforeStart.val == 0 {
// 		return afterStart
// 	}

// 	//merge the two lists
// 	beforeEnd.next = afterStart
// 	return beforeStart
// }

// //Given two strings, write a method to decide if one is a permutation of the other
// func checkPermutation(s1 string, s2 string) bool {
// 	//sort and compare
// }

// func thirdGreatest(arr []int) int {
// 	var first int
// 	var second int
// 	var third int

// 	for _, el := range arr {
// 		if el >= first {
// 			third = second
// 			second = first
// 			first = el
// 			continue
// 		}
// 		if el >= second {
// 			third = second
// 			second = el
// 			continue
// 		}
// 		if el >= third {
// 			third = el
// 		}
// 	}
// 	return third
// }

// func isPrime(n int) bool {
// 	for i := 2; i < n/2+1; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func nthPrime(nth int) int {
// 	primeCount := 0
// 	n := 2
// 	for primeCount < nth {
// 		if isPrime(n) {
// 			primeCount++
// 		}
// 		n++
// 	}
// 	return n - 1
// }
