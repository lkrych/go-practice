package goExercises

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

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

func TestNumRepeats(t *testing.T) {
	test1 := numRepeats("abdbc")
	test2 := numRepeats("aaa")
	test3 := numRepeats("abab")
	test4 := numRepeats("cadac")
	test5 := numRepeats("abcde")

	if test1 != 1 {
		t.Errorf("Expected numRepeats of abdbc to be 1, not %v", test1)
	}

	if test2 != 1 {
		t.Errorf("Expected numRepeats of aaa to be 1, not %v", test2)
	}

	if test3 != 2 {
		t.Errorf("Expected numRepeats of abab to be 2, not %v", test3)
	}

	if test4 != 2 {
		t.Errorf("Expected numRepeats of cadac to be 2, not %v", test4)
	}

	if test5 != 0 {
		t.Errorf("Expected numRepeats of abcde to be 0, not %v", test5)
	}
}
