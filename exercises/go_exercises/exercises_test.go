package goExercises

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIncrementArbitrary(t *testing.T) {
	test1 := incrementArb([]int{1, 2, 9})
	test2 := incrementArb([]int{1, 2, 9, 8, 2, 3, 4, 1, 3, 5, 0})
	test3 := incrementArb([]int{1, 2, 5, 3, 2, 4, 6, 7, 8, 2})
	test4 := incrementArb([]int{5, 3, 6, 3, 7, 9, 5, 0})
	test5 := incrementArb([]int{1, 2, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9})
	ans1 := []int{1, 3, 0}
	ans2 := []int{1, 2, 9, 8, 2, 3, 4, 1, 3, 5, 1}
	ans3 := []int{1, 2, 5, 3, 2, 4, 6, 7, 8, 3}
	ans4 := []int{5, 3, 6, 3, 7, 9, 5, 1}
	ans5 := []int{1, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected incrementArb to increment the array %v, instead you got %v", ans1, test1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected incrementArb to increment the array %v, instead you got %v", ans2, test2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected incrementArb to increment the array %v, instead you got %v", ans3, test3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected incrementArb to increment the array %v, instead you got %v", ans4, test4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected incrementArb to increment the array %v, instead you got %v", ans5, test5)
	}
}

func TestIsPalindrome(t *testing.T) {
	test1 := isPalindrome("abc")
	test2 := isPalindrome("racecar")
	test3 := isPalindrome("z")

	if test1 != false {
		t.Errorf("Expected isPalindrome of abc to be false, not %v", test1)
	}

	if test2 != true {
		t.Errorf("Expected isPalindrome of racecar to be true, not %v", test2)
	}

	if test3 != true {
		t.Errorf("Expected isPalindrome of z to be true, not %v", test3)
	}
}

func TestMergeSort(t *testing.T) {
	test1 := mergeSort([]int{7, 5, 2, 4, 8, 9, 3})
	test2 := mergeSort([]int{7, 5, 2, 4, 8, 9, 3, 1, 6, 10})
	test3 := mergeSort([]int{7, 5, 2, 4, 8, 9, 3, 3, 3, 3})
	test4 := mergeSort([]int{7, 5, 2, 4, 8, 9, 3, 100, 102, 6, 1, 10})
	test5 := mergeSort([]int{7, 5, 2, 4, 8, 9, 3, 11, 10, 6, 1})
	ans1 := []int{2, 3, 4, 5, 7, 8, 9}
	ans2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans3 := []int{2, 3, 3, 3, 3, 4, 5, 7, 8, 9}
	ans4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 102}
	ans5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected mergeSort to sort the array %v, instead you got %v", ans1, test1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected mergeSort to sort the array %v, instead you got %v", ans2, test2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected mergeSort to sort the array %v, instead you got %v", ans3, test3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected mergeSort to sort the array %v, instead you got %v", ans4, test4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected mergeSort to sort the array %v, instead you got %v", ans5, test5)
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	test1 := isPowerOfTwo(1)
	test2 := isPowerOfTwo(16)
	test3 := isPowerOfTwo(64)
	test4 := isPowerOfTwo(78)
	test5 := isPowerOfTwo(0)

	if test1 != true {
		t.Errorf("Expected isPowerOfTwo of 1 to be true, not %v", test1)
	}

	if test2 != true {
		t.Errorf("Expected isPowerOfTwo of 16 to be true, not %v", test2)
	}

	if test3 != true {
		t.Errorf("Expected isPowerOfTwo of 64 to be true, not %v", test3)
	}

	if test4 != false {
		t.Errorf("Expected isPowerOfTwo of 78 to be false, not %v", test4)
	}

	if test5 != false {
		t.Errorf("Expected isPowerOfTwo of 0 to be false, not %v", test5)
	}
}

func TestAdvancingArray(t *testing.T) {
	test1 := []int{3, 3, 1, 0, 2, 0, 1}
	test2 := []int{3, 2, 0, 0, 2, 0, 1}
	test3 := []int{5, 3, 1, 0, 2, 0, 1}
	test4 := []int{0, 3, 1, 0, 2, 0, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	if advancingArray(test1) != true {
		t.Errorf("Expected array %v to be true", test1)
	}
	if advancingArray(test2) != false {
		t.Errorf("Expected array %v to be false", test2)
	}

	if advancingArray(test3) != true {
		t.Errorf("Expected array %v to be true", test3)
	}

	if advancingArray(test4) != false {
		t.Errorf("Expected array %v to be false", test4)
	}

	if advancingArray(test5) != true {
		t.Errorf("Expected array %v to be true", test5)
	}
}
