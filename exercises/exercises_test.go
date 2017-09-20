package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	word := reverse("candelabra")
	sentence := reverse("reverse this sentence")

	if word != "arbalednac" {
		t.Errorf("Expected reverse to be able to reverse a string")
	}

	if sentence != "ecnetnes siht esrever" {
		t.Errorf("Expected reverse to be able to reverse a sentence, but got %v", sentence)
	}

}

func TestFactorial(t *testing.T) {
	five := factorial(5)
	ten := factorial(10)
	one := factorial(1)

	if five != 120 {
		t.Errorf("Expected factorial of 5 to equal 120, not %v", five)
	}

	if ten != 3628800 {
		t.Errorf("Expected factorial of 10 to equal 3628800, not %v", ten)
	}

	if one != 1 {
		t.Errorf("Expected factorial of 1 to equal 1, not %v", one)
	}
}

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
