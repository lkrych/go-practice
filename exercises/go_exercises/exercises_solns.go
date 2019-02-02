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
 
// //given two numbers represented by linked lists, where each node contains a single digit, add them together
// func sumLists(l1 *Node, l2 *Node) *Node {
// 	carry := 0
// 	sum := 0

// 	head := &Node{
// 		val:  0,
// 		next: nil,
// 	}
// 	currentNode := head

// forLoop:
// 	for l1 != nil || l2 != nil {
// 		intermediate := 0
// 		switch {
// 		case l1 != nil && l2 != nil:
// 			intermediate = (l1.val + l2.val + carry)
// 			sum = intermediate % 10
// 			carry = intermediate / 10
// 			l1 = l1.next
// 			l2 = l2.next

// 		case l1 != nil:
// 			intermediate = (l1.val + carry)
// 			sum = intermediate % 10
// 			carry = intermediate / 10
// 			l1 = l1.next
// 		case l2 != nil:
// 			intermediate = (l2.val + carry)
// 			sum = intermediate % 10
// 			carry = intermediate / 10
// 			l2 = l2.next
// 		default:
// 			intermediate = carry
// 			sum = intermediate % 10
// 			carry = intermediate / 10
// 		}

// 		currentNode.val = sum
// 		nextNode := &Node{
// 			val:  0,
// 			next: nil,
// 		}
// 		//breaking logic
// 		switch {
// 		case l1 == nil && l2 == nil && carry == 0:
// 			break forLoop
// 		case l1 == nil && l2 == nil && carry != 0:
// 			currentNode.next = nextNode
// 			currentNode = nextNode
// 			currentNode.val = carry
// 			break forLoop
// 		}

// 		currentNode.next = nextNode
// 		currentNode = nextNode
// 	}

// 	return head

// }

 
