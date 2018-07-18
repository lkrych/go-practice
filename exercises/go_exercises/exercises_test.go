package goExercises

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreatingAQueue(t *testing.T) {
	q := &queue{
		data: []int{1, 2, 3, 4, 5, 6},
	}
	testremove := q.remove()
	testremove2 := q.remove()

	if len(q.data) != 4 {
		t.Errorf("Expected remove to return elements from the top of the queue: %v, removeped: %v", q.data, testremove)
	}
	if q.peek() != 3 {
		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
	}
	if len(q.data) != 4 {
		t.Errorf("Expected peek to not delete elements from the top of the queue")
	}
	if q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue has %v elements", len(q.data))
	}

	q.add(testremove)
	if len(q.data) != 5 {
		t.Errorf("Expected remove to return elements from the top of the queue: %v, removeped: %v", q.data, testremove2)
	}
	if q.peek() != 3 {
		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
	}
	if len(q.data) != 5 {
		t.Errorf("Expected peek to not delete elements from the top of the queue")
	}
	if q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue has %v elements\n", len(q.data))
	}

	q.add(testremove2)
	if len(q.data) != 6 {
		t.Errorf("Expected push to return elements to the top of the queue")
	}
	if q.peek() != 3 {
		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
	}
	if len(q.data) != 6 {
		t.Errorf("Expected peek to not delete elements from the top of the queue")
	}
	if q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue has %v elements\n", len(q.data))
	}

	for i := 0; i < 6; i++ {
		q.remove()
	}
	if !q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue should be empty")
	}
}

func TestBubbleSort(t *testing.T) {
	test1 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3})
	test2 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3, 1, 6, 10})
	test3 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3, 3, 3, 3})
	test4 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3, 100, 102, 6, 1, 10})
	test5 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3, 11, 10, 6, 1})
	ans1 := []int{2, 3, 4, 5, 7, 8, 9}
	ans2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans3 := []int{2, 3, 3, 3, 3, 4, 5, 7, 8, 9}
	ans4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 102}
	ans5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test1, ans1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test2, ans2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test3, ans3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test4, ans4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test5, ans5)
	}
}

func TestisUnique(t *testing.T) {
	test1 := isUnique("blah")
	test2 := isUnique("banana")
	test3 := isUnique("fike hyjrant")
	test4 := isUnique("fire hydrant")
	test5 := isUnique("Another Day")

	if test1 != true {
		t.Errorf("Expected blah to be true because all letters are unique ")
	}

	if test2 != false {
		t.Errorf("Expected banana to be false because all letters are not unique ")
	}

	if test3 != true {
		t.Errorf("Expected fike hyjrant to be true because all letters are unique ")
	}

	if test4 != false {
		t.Errorf("Expected fire hydrant to be false because all letters are not unique ")
	}

	if test5 != false {
		t.Errorf("Expected Another day to be false because all letters are not unique ")
	}
}

func TestCountVowels(t *testing.T) {
	test1 := countVowels("abcd")
	test2 := countVowels("colour")
	test3 := countVowels("cecilia nutley")

	if test1 != 1 {
		t.Errorf("Expected countVowels of abcd to be 1, not %v", test1)
	}

	if test2 != 3 {
		t.Errorf("Expected countVowels of colour to be 3, not %v", test2)
	}

	if test3 != 6 {
		t.Errorf("Expected countVowels of mom to be 6, not %v", test3)
	}
}

func TestReverse(t *testing.T) {
	test1 := stringRotation("candelabra", "elabracand")
	test2 := stringRotation("reversethissentence", "entencereversethiss")
	test3 := stringRotation("blah", "cactus")

	if test1 != true {
		t.Errorf("Expected stringRotation to be able to detect a rotation in candelabra and elabracand")
	}

	if test2 != true {
		t.Errorf("Expected stringRotation to be able to detect a rotation in reversethissentence and entencereversethiss")
	}

	if test3 != false {
		t.Errorf("Expected stringRotation to be able to detect there isn't a rotation in blah and cactus")
	}

}
