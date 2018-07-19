package goExercises

import "errors"

//given an arr representing stock prices, return the maximum profit you could retrieve from buying and selling the stock
//ex: [310, 315, 275, 295, 260, 270, 290, 230, 255, 250] => 30
func buyStock(arr []int) int {
	//buy low sell high
	low := 9999999
	diff := 0
	for _, el := range arr {
		if el < low {
			low = el
		}
		if el-low > diff {
			diff = el - low
		}
	}
	return diff
}

//return index of toFind if it exists in arr, otherwise return -1
func binarySearch(arr []int, toFind int) int {
	if len(arr) < 1 {
		return -1
	}
	midIdx := len(arr) / 2
	if arr[midIdx] == toFind {
		return midIdx
	}
	if arr[midIdx] > toFind {
		return binarySearch(arr[:midIdx], toFind)
	}
	returned := binarySearch(arr[midIdx+1:], toFind)
	if returned != -1 {
		return midIdx + 1 + returned
	}
	return returned

}

// return true if n is prime
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//create a function fibonacci(n) that prints returns n fibonacci numbers in an array
func fibonacci(n int) ([]int, error) {
	if n < 1 {
		return nil, errors.New("n must be greater than 0")
	}
	fibs := []int{1, 1}
	for len(fibs) != n {
		fibs = append(fibs, fibs[len(fibs)-1]+fibs[len(fibs)-2])
	}
	return fibs, nil
}

//you have two numbers represented by a linked list, where each node contains a single digit
//the digits are stored in reverse order, such that 1's digit is at the head of the list
//write a function that adds the two numbers and returns the sum as a linked list
//example:
//(7 -> 1 -> 6) + (5 -> 9 -> 2) = (2 -> 1 -> 9)
type Node struct {
	val  int
	next *Node
}

func sumLists(h1, h2 *Node) *Node {
	remainder := 0
	sum := 0
	head := &Node{}
	currentNode := head
	for h1 != nil || h2 != nil {
		sum = h1.val + h2.val + remainder
		remainder = sum / 10
		currentNode.val = sum % 10

		newNode := &Node{}
		//move up the lists
		h1 = h1.next
		h2 = h2.next

		if h1 != nil || h2 != nil {
			currentNode.next = newNode
			currentNode = newNode
		}
	}
	if h1 != nil {
		sum = h1.val + remainder
		remainder = sum / 10
		currentNode.val = sum % 10
	}

	if h2 != nil {
		sum = h2.val + remainder
		remainder = sum / 10
		currentNode.val = sum % 10
	}

	if remainder > 0 {
		currentNode.next = &Node{
			val: remainder,
		}
	}

	return head
}
