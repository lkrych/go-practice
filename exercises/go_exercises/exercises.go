package goExercises

//given a linked list, delete nodes with the the value of x
//optional create a function that builds a linked list with an input array of integers
//if you don't want to do this, you can just look in the test file for this problem and grab one

type Node struct {
	val  int
	next *Node
}

func deleteANode(h *Node) *Node {

}

func createLinkedList(input []int) *Node {
	head := &Node{
		val:  input[0],
		next: nil,
	}
	currentNode := head
	for _, el := range input[1:]
}

//return n!
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

//Given a string, find the length of the longest substring without repeating characters.
func longestSubString(sentence string) int {
	//iterate through each character in string
	//checking to see if that character exists in the currentLongest subString
	//if it does, start building the longest subString from the next index
	//otherwise, keep adding to it.
}

//convert base ten int n, into hexadecimal or binary representation
func baseConverter(n int, base int) string {
	// use the % and / operator to slowly build a number
}

//take in an array of digits encoding a decimal, d and update the array to rep d + 1
//ex: 1.29 => 1.30 ~ [1,2,9] => [1,3,0]
//ex 1.48372 => 1.48373
//TAKE AWAY:  This algorithm should work in a language that has finite-precision arithmetic
//Arrays can be used to break up numbers of arbitrary size, this allows us to handle very big numbers!
func incrementArb(arr []int) []int {

}
