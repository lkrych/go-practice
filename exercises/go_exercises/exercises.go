package goExercises

//take in an array of digits encoding a decimal, d and update the array to rep d + 1
//ex: 1.29 => 1.30 ~ [1,2,9] => [1,3,0]
//ex 1.48372 => 1.48373
//TAKE AWAY:  This algorithm should work in a language that has finite-precision arithmetic
//Arrays can be used to break up numbers of arbitrary size, this allows us to handle very big numbers!
func incrementArb(arr []int) []int {
	carry := 0
	for i := len(arr) - 1; i >= 0; i-- {
		currentIdx := arr[i]
		currentSum := 0
		if i == len(arr)-1 {
			currentSum = carry + currentIdx + 1
		} else {
			currentSum = carry + currentIdx
		}
		carry = currentSum / 10
		arr[i] = currentSum % 10
	}

	if carry != 0 {
		arr = append([]int{1}, arr...)
	}

	return arr

}

//given a linked list, delete nodes with the the value of x
//optional create a function that builds a linked list with an input array of integers
//if you don't want to do this, you can just look in the test file for this problem and grab one
func deleteANode(list *Node, el int) *Node {
	currentHead := list
	currentNode := list
	if currentNode.val == el {
		currentHead = currentHead.next
	}
	for currentNode.next != nil {
		pastNode := currentNode
		currentNode = currentNode.next
		if currentNode.val == el {
			pastNode.next = currentNode.next
		}
	}

	return currentHead
}
