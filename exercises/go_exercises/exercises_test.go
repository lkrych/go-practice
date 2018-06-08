package goExercises

import (
	"testing"
)

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

func TestCheckSum(t *testing.T) {
	test1 := checkSum([][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}})
	test2 := checkSum([][]int{{1, 2, 3, 4}, {1, 2, 3}, {1, 2, 3}})
	test3 := checkSum([][]int{{1, 2, 3, 10}, {1, 2, 3, 12}, {1, 2, 3, 15}})
	test4 := checkSum([][]int{{3, 22, 124}, {3, 70}, {1, 2, 5}})
	test5 := checkSum([][]int{{1, 2, 10}, {1, 11, 3}, {500, 2, 3}})
	// test6 := checkSum(checkSumAdvent())

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

	// if test6 != 53978 {
	// 	t.Errorf("Expected answer to be 53978, not %v", test6)
	// }
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
	word := reverse("candelabra")
	sentence := reverse("reverse this sentence")

	if word != "arbalednac" {
		t.Errorf("Expected reverse to be able to reverse a string")
	}

	if sentence != "ecnetnes siht esrever" {
		t.Errorf("Expected reverse to be able to reverse a sentence, but got %v", sentence)
	}

}
