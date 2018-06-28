package goExercises

// import "math"

// func isPowerOfTwo(n float64) bool {
// 	power := float64(0)
// 	value := math.Exp2(power)
// 	for value <= n {
// 		if value == n {
// 			return true
// 		}
// 		power++
// 		value = math.Exp2(power)
// 	}
// 	return false
// }

// func partitionEvenOdd(arr []int) []int {
// 	partition := 0
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i]%2 == 0 {
// 			arr[i], arr[partition] = arr[partition], arr[i]
// 			partition++
// 		}
// 	}
// 	front := quickSort(arr[:partition])
// 	back := quickSort(arr[partition:])
// 	return append(front, back...)
// }

// func quickSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	pivot := arr[0]
// 	left := []int{}
// 	right := []int{}

// 	for _, val := range arr[1:] {
// 		if val >= pivot {
// 			right = append(right, val)
// 		} else {
// 			left = append(left, val)
// 		}
// 	}
// 	answer := []int{}
// 	sortedLeft := quickSort(left)
// 	sortedRight := quickSort(right)
// 	answer = append(answer, sortedLeft...)
// 	answer = append(answer, pivot)
// 	answer = append(answer, sortedRight...)
// 	return answer
// }

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"strings"
// )

// func passPhrase(passphrase string) bool {
// 	split := strings.Split(passphrase, " ")
// 	wordMap := make(map[string]int)
// 	for _, word := range split {
// 		if _, ok := wordMap[word]; ok {
// 			fmt.Printf("%v is the repeated word in %v \n", word, passphrase)
// 			return false
// 		}
// 		wordMap[word] = 1
// 	}
// 	return true
// }

// func passPhraseAdvent() int {
// 	phraseCount := 0
// 	content, err := ioutil.ReadFile("./advent_input/passphrase_input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	splitByLine := strings.Split(string(content), "\n")
// 	for _, line := range splitByLine {
// 		if passPhrase(line) {
// 			phraseCount++
// 		}
// 	}
// 	return phraseCount
// }

// func passPhraseAnagram(passphrase string) bool {
// 	split := strings.Split(passphrase, " ")
// 	wordMap := make(map[string]int)
// 	for _, word := range split {
// 		sorted := SortString(word)
// 		if _, ok := wordMap[sorted]; ok {
// 			return false
// 		}
// 		wordMap[sorted] = 1
// 	}
// 	return true
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

// func sumNums(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	return n + sumNums(n-1)
// }
