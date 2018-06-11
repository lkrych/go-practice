package goExercises

import (
	"testing"
)

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

func TestDasherizeNumber(t *testing.T) {
	test1 := dasherizeNumber(203)
	test2 := dasherizeNumber(303)
	test3 := dasherizeNumber(333)
	test4 := dasherizeNumber(3223)

	if test1 != "20-3" {
		t.Errorf("Expected dasherizeNumber of 203  to be 20-3, not %v", test1)
	}

	if test2 != "3-0-3" {
		t.Errorf("Expected dasherizeNumber of 303  to be 3-0-3, not %v", test2)
	}

	if test3 != "3-3-3" {
		t.Errorf("Expected dasherizeNumber of 333 to be 3-3-3, not %v", test3)
	}

	if test4 != "3-22-3" {
		t.Errorf("Expected dasherizeNumber of 3223  to be 3-22-3, not %v", test4)
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

func TestLongestSubString(t *testing.T) {
	test1 := longestSubString("abcabcbb")
	test2 := longestSubString("bbbbb")
	test3 := longestSubString("pwwkew")

	if test1 != 3 {
		t.Errorf("Expected longestSubString of abcabcbb to be 3, not %v", test1)
	}

	if test2 != 1 {
		t.Errorf("Expected longestSubString of abba to be 1, not %v", test2)
	}

	if test3 != 3 {
		t.Errorf("Expected longestSubString of pwwkew to be 3, not %v", test3)
	}
}
