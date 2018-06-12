package goExercises

import (
	"testing"
)

func TestStringCompression(t *testing.T) {
	test1 := stringCompression("abcd")
	test2 := stringCompression("aaabbbbccccccdddd")
	test3 := stringCompression("aabbbbcccccddd")

	if test1 != "abcd" {
		t.Errorf("Expected stringCompression of 'abcd' to be 'abcd', not %v", test1)
	}

	if test2 != "a3b4c6d4" {
		t.Errorf("Expected stringCompression of 'aaabbbbccccccdddd' to be 'a3b4c6d4', not %v", test2)
	}

	if test3 != "a2b4c5d3" {
		t.Errorf("Expected stringCompression of 'aabbbbcccccddd' to be 'a2b4c5d3', not %v", test3)
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

func TestPassPhrase(t *testing.T) {
	test1 := passPhrase("doody goody woody")
	test2 := passPhrase("doody doody woody")
	test3 := passPhrase("woody goody woody")
	test4 := passPhrase("It is not I who is the Creature!")
	test5 := passPhrase("Hello, my name is Bill.")
	test6 := passPhraseAdvent()

	if test1 != true {
		t.Errorf("Expected answer to be true, not %v", test1)
	}

	if test2 != false {
		t.Errorf("Expected answer to be false, not %v", test2)
	}

	if test3 != false {
		t.Errorf("Expected answer to be false, not %v", test3)
	}

	if test4 != false {
		t.Errorf("Expected answer to be false, not %v", test4)
	}

	if test5 != true {
		t.Errorf("Expected answer to be true, not %v", test5)
	}

	if test6 != 325 {
		t.Errorf("Expected answer to be 325, not %v", test6)
	}
}

// func TestPassPhraseAnagram(t *testing.T) {
// 	test1 := passPhraseAnagram("abcde fghij")
// 	test2 := passPhraseAnagram("abcde xyz ecdab")
// 	test3 := passPhraseAnagram("a ab abc abd abf abj")
// 	test4 := passPhraseAnagram("iiii oiii ooii oooi oooo")
// 	test5 := passPhraseAnagram("oiii ioii iioi iiio")
// 	test6 := passPhraseAnagramAdvent()

// 	if test1 != true {
// 		t.Errorf("Expected answer to be true, not %v", test1)
// 	}

// 	if test2 != false {
// 		t.Errorf("Expected answer to be false, not %v", test2)
// 	}

// 	if test3 != true {
// 		t.Errorf("Expected answer to be true, not %v", test3)
// 	}

// 	if test4 != true {
// 		t.Errorf("Expected answer to be false, not %v", test4)
// 	}

// 	if test5 != false {
// 		t.Errorf("Expected answer to be false, not %v", test5)
// 	}

// 	if test6 != 6 {
// 		t.Errorf("Expected answer to be 6, not %v", test6)
// 	}
// }
