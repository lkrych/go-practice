package goExercises
 
// import (
// 	"strconv"
// 	"strings"
// )

// func mostCommonLetter(sentence string) []string {
// 	lettersMap := map[string]int{}
// 	for _, char := range strings.Split(sentence, "") {
// 		if char == " " {
// 			continue
// 		}
// 		if lettersMap[char] >= 1 {
// 			lettersMap[char]++
// 		} else {
// 			lettersMap[char] = 1
// 		}
// 	}
// 	count := 0
// 	mostCommon := []string{}
// 	for key, val := range lettersMap {
// 		if val > count {
// 			count = val
// 			mostCommon = []string{key, strconv.Itoa(val)}
// 		}
// 	}
// 	return mostCommon
// }
 
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

 
