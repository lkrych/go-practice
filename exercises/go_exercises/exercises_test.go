package goExercises

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLongestWord(t *testing.T) {
	test1 := longestWord("short longest")
	test2 := longestWord("one")
	test3 := longestWord("abc def abcde")

	if test1 != "longest" {
		t.Errorf("Expected the longest word in test1 to be longest, not %v", test1)
	}

	if test2 != "one" {
		t.Errorf("Expected the longest word in test2 to be one, not %v", test2)
	}

	if test3 != "abcde" {
		t.Errorf("Expected the longest word in test3 to be abcde, not %v", test3)
	}
}

func TestLongestPalindrome(t *testing.T) {
	test1 := longestPalindrome("abcbd")
	test2 := longestPalindrome("abba")
	test3 := longestPalindrome("abcbdeffe")

	if test1 != "bcb" {
		t.Errorf("Expected longestPalindrome of abcbd to be bcb, not %v", test1)
	}

	if test2 != "abba" {
		t.Errorf("Expected longestPalindrome of abba to be abba, not %v", test2)
	}

	if test3 != "effe" {
		t.Errorf("Expected longestPalindrome of abcbdeffe to be effe, not %v", test3)
	}
}

func TestSumNums(t *testing.T) {
	test1 := sumNums(1)
	test2 := sumNums(2)
	test3 := sumNums(5)

	if test1 != 1 {
		t.Errorf("Expected sumNums(1) to equal 1, not %v", test1)
	}

	if test2 != 3 {
		t.Errorf("Expected sumNums(2) to equal 3, not %v", test2)
	}

	if test3 != 15 {
		t.Errorf("Expected sumNums(5) to equal 15, not %v", test3)
	}
}

func TestMostCommonLetter(t *testing.T) {
	test1 := mostCommonLetter("abca")
	ans1 := []string{
		"a", "2",
	}
	test2 := mostCommonLetter("abbab")
	ans2 := []string{
		"b", "3",
	}
	test3 := mostCommonLetter("guernsey was here")
	ans3 := []string{
		"e", "4",
	}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected mostCommonLetter of abca to be a with a count of 2, not %v", test1)
	}

	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected mostCommonLetter of abbab to be b with a count of 3, not %v", test2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected mostCommonLetter of guernsey was here to be e with a count of 4, not %v", test3)
	}
}

func TestBuyStockTwice(t *testing.T) {
	test1 := []int{12, 11, 13, 9, 12, 8, 14, 13, 15}
	test2 := []int{3, 2, 0, 0, 2, 0, 1, 5}
	test3 := []int{5, 3, 1, 5, 2, 0, 1}
	test4 := []int{5, 3, 2, 4, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	test6 := []int{7, 6, 5, 4, 3, 2}
	ans1 := buyStockTwice(test1)
	ans2 := buyStockTwice(test2)
	ans3 := buyStockTwice(test3)
	ans4 := buyStockTwice(test4)
	ans5 := buyStockTwice(test5)
	ans6 := buyStockTwice(test6)
	if ans1 != 10 {
		t.Errorf("Expected array %v to be 10, not %v", test1, ans1)
	}
	if ans2 != 5 {
		t.Errorf("Expected array %v to be 5, not %v", test2, ans2)
	}

	if ans3 != 4 {
		t.Errorf("Expected array %v to be 4, not %v", test3, ans3)
	}

	if ans4 != 2 {
		t.Errorf("Expected array %v to be 2, not %v", test4, ans4)
	}

	if ans5 != 6 {
		t.Errorf("Expected array %v to be 6, not %v", test5, ans5)
	}

	if ans6 != 0 {
		t.Errorf("Expected array %v to be 0, not %v", test6, ans6)
	}
}
