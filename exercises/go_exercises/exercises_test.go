package goExercises

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuyStock(t *testing.T) {
	test1 := []int{3, 3, 1, 0, 2, 0, 1}
	test2 := []int{3, 2, 0, 0, 2, 0, 1, 5}
	test3 := []int{5, 3, 1, 5, 2, 0, 1}
	test4 := []int{5, 3, 2, 4, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	test6 := []int{7, 6, 5, 4, 3, 2}
	if buyStock(test1) != 2 {
		t.Errorf("Expected array %v to be 2 not %v", test1, buyStock(test1))
	}
	if buyStock(test2) != 5 {
		t.Errorf("Expected array %v to be 5 not %v", test2, buyStock(test2))
	}

	if buyStock(test3) != 4 {
		t.Errorf("Expected array %v to be 4 not %v", test3, buyStock(test3))
	}

	if buyStock(test4) != 2 {
		t.Errorf("Expected array %v to be 2 not %v", test4, buyStock(test4))
	}

	if buyStock(test5) != 6 {
		t.Errorf("Expected array %v to be 6 not %v", test5, buyStock(test5))
	}

	if buyStock(test6) != 0 {
		t.Errorf("Expected array %v to be 0 not %v", test6, buyStock(test6))
	}
}

func TestBinarySearch(t *testing.T) {
	test1 := binarySearch([]int{0, 1, 2, 3, 4, 5}, 3)
	test2 := binarySearch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9)
	test3 := binarySearch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)
	test4 := binarySearch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
	test5 := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0)

	if test1 != 3 {
		t.Errorf("Expected binarySearch to find the index of 3 at 4, not %v", test1)
	}
	if test2 != 9 {
		t.Errorf("Expected binarySearch to find the index of 9 at 10, not %v", test2)
	}

	if test3 != 0 {
		t.Errorf("Expected binarySearch to find the index of 0 at 0, not %v", test3)
	}

	if test4 != -1 {
		t.Errorf("Expected binarySearch to find the index of 10 at -1, not %v", test4)
	}

	if test5 != -1 {
		t.Errorf("Expected binarySearch to find the index of 0 at -1, not %v", test5)
	}
}

func TestIsPrime(t *testing.T) {
	test1 := isPrime(2)
	test2 := isPrime(3)
	test3 := isPrime(4)
	test4 := isPrime(9)
	test5 := isPrime(227)

	if test1 != true {
		t.Errorf("Expected isPrime of 2 to be true, not false")
	}

	if test2 != true {
		t.Errorf("Expected isPrime of 3 to be true, not false")
	}

	if test3 != false {
		t.Errorf("Expected isPrime of 4 to be false, not true")
	}

	if test4 != false {
		t.Errorf("Expected isPrime of 9 to be false, not true")
	}

	if test5 != true {
		t.Errorf("Expected isPrime of 227 to be true, not false")
	}
}

func testFibonacci(t *testing.T) {
	t1, err1 := fibonacci(3)
	t2, err2 := fibonacci(6)
	t3, err3 := fibonacci(7)
	t4, err4 := fibonacci(9)
	t5, err5 := fibonacci(15)
	_, err6 := fibonacci(-1)
	_, err7 := fibonacci(0)
	ans1 := []int{1, 1, 2}
	ans2 := []int{1, 1, 2, 3, 5, 8}
	ans3 := []int{1, 1, 2, 3, 5, 8, 13}
	ans4 := []int{1, 1, 2, 3, 5, 8, 13, 21, 34}
	ans5 := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	errorMessage := "n must be greater than 0"

	if !cmp.Equal(t1, ans1) {
		t.Errorf("Expected the fibonacci sequence of %v to be %v not %v", 3, ans1, t1)
	}

	if err1 != nil {
		t.Errorf("There was not supposed to be an error for t1")
	}

	if !cmp.Equal(t2, ans2) {
		t.Errorf("Expected the fibonacci sequence of %v to be %v not %v", 6, ans2, t2)
	}

	if err2 != nil {
		t.Errorf("There was not supposed to be an error for t2")
	}

	if !cmp.Equal(t3, ans3) {
		t.Errorf("Expected the fibonacci sequence of %v to be %v not %v", 7, ans3, t3)
	}

	if err3 != nil {
		t.Errorf("There was not supposed to be an error for t3")
	}

	if !cmp.Equal(t4, ans4) {
		t.Errorf("Expected the fibonacci sequence of %v to be %v not %v", 9, ans4, t4)
	}

	if err4 != nil {
		t.Errorf("There was not supposed to be an error for t4")
	}

	if !cmp.Equal(t5, ans5) {
		t.Errorf("Expected the fibonacci sequence of %v to be %v not %v", 15, ans5, t5)
	}
	if err5 != nil {
		t.Errorf("There was not supposed to be an error for t5")
	}

	if err6.Error() != errorMessage {
		t.Errorf("Expected the function to throw an error for input less than 1")
	}

	if err7.Error() != errorMessage {
		t.Errorf("Expected the function to throw an error for input less than 1")
	}
}

func TestSumLists(t *testing.T) {
	i1 := []int{1, 2, 3}
	i2 := []int{5, 8, 6}
	sum1 := []int{6, 0, 0, 1}

	i3 := []int{3, 5}
	i4 := []int{1, 8}
	sum2 := []int{4, 3, 1}

	i5 := []int{6, 4, 5}
	i6 := []int{1, 1, 1}
	sum3 := []int{7, 5, 6}

	list1 := createALinkedList(i1)
	list2 := createALinkedList(i2)
	answer1 := listToArray(createALinkedList(sum1))

	list3 := createALinkedList(i3)
	list4 := createALinkedList(i4)
	answer2 := listToArray(createALinkedList(sum2))

	list5 := createALinkedList(i5)
	list6 := createALinkedList(i6)
	answer3 := listToArray(createALinkedList(sum3))

	returned1 := listToArray(sumLists(list1, list2))
	returned2 := listToArray(sumLists(list3, list4))
	returned3 := listToArray(sumLists(list5, list6))

	if !cmp.Equal(answer1, returned1) {
		t.Errorf("Expected sumLists of %v and %v to be %v not %v", list1, list2, answer1, returned1)

	}

	if !cmp.Equal(answer2, returned2) {
		t.Errorf("Expected sumLists of %v and %v to be %v not %v", list3, list4, answer2, returned2)
	}

	if !cmp.Equal(answer3, returned3) {
		t.Errorf("Expected sumLists of %v and %v to be %v not %v", list5, list6, answer3, returned3)
	}
}

func listToArray(head *Node) []int {
	vals := []int{}
	for head != nil {
		vals = append(vals, head.val)
		head = head.next
	}
	return vals
}

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
