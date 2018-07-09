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

 
// func factorial(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	return n * factorial(n-1)
// }
 
// //Given a string, find the length of the longest substring without repeating characters.
// func lengthOfLongestSubstring(s string) int {
// 	bs := []byte(s)
// 	var maxLen, start int
// 	idxMap := map[byte]int{}
// 	for i := 0; i < len(bs); i++ {
// 			if _, ok := idxMap[bs[i]]; ok && start <= idxMap[bs[i]] {
// 					start = idxMap[bs[i]] + 1
// 			} else {
// 					if maxLen < i - start + 1 {
// 							maxLen = i - start + 1
// 					}
// 			}
// 			idxMap[bs[i]] = i
// 	}
// 	return maxLen
// }
 
// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func baseConverter(n int, base int) string {
// 	base16 := map[int]string{
// 		10: "a",
// 		11: "b",
// 		12: "c",
// 		13: "d",
// 		14: "e",
// 		15: "f"}
// 	answer := []string{}
// 	for n > 0 {
// 		answer = append(answer, strconv.Itoa(n%base))
// 		n = n / base
// 	}
// 	if base == 16 {
// 		for i, el := range answer {
// 			intEl, err := strconv.Atoi(el)
// 			if err != nil {
// 				fmt.Println("There was an error!")
// 			}
// 			if hex, exists := base16[intEl]; exists {
// 				answer[i] = hex
// 			}
// 		}
// 	}
// 	return reverse(strings.Join(answer, ""))
// }
 
// func incrementArb(arr []int) []int {
// 	lastIdx := len(arr) - 1
// 	arr[lastIdx] = (arr[lastIdx] + 1) % 10
// 	if arr[lastIdx] == 0 {
// 		arr = append(incrementArb(arr[:lastIdx]), arr[lastIdx])
// 	}
// 	return arr
// }
 
