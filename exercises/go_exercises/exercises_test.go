package goExercises

import (
	"testing"
)

func TestCheckSum(t *testing.T) {
	test1 := checkSum([][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}})
	test2 := checkSum([][]int{{1, 2, 3, 4}, {1, 2, 3}, {1, 2, 3}})
	test3 := checkSum([][]int{{1, 2, 3, 10}, {1, 2, 3, 12}, {1, 2, 3, 15}})
	test4 := checkSum([][]int{{3, 22, 124}, {3, 70}, {1, 2, 5}})
	test5 := checkSum([][]int{{1, 2, 10}, {1, 11, 3}, {500, 2, 3}})
	test6 := checkSumAdvent()

	if test1 != 6 {
		t.Errorf("Expected answer to be 6, not %v", test1)
	}
	if test2 != 7 {
		t.Errorf("Expected answer to be 7, not %v", test2)
	}

	if test3 != 34 {
		t.Errorf("Expected answer to be 34, not %v", test3)
	}

	if test4 != 192 {
		t.Errorf("Expected answer to be 192, not %v", test4)
	}

	if test5 != 517 {
		t.Errorf("Expected answer to be 517, not %v", test5)
	}

	if test6 != 53978 {
		t.Errorf("Expected answer to be 53978, not %v", test6)
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
