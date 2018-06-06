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

func TestBuyStockTwice(t *testing.T) {
	test1 := []int{12, 11, 13, 9, 12, 8, 14, 13, 15}
	test2 := []int{3, 2, 0, 0, 2, 0, 1, 5}
	test3 := []int{5, 3, 1, 5, 2, 0, 1}
	test4 := []int{5, 3, 2, 4, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	test6 := []int{7, 6, 5, 4, 3, 2}
	if buyStockTwice(test1) != 10 {
		t.Errorf("Expected array %v to be 2", test1)
	}
	if buyStockTwice(test2) != 7 {
		t.Errorf("Expected array %v to be 5", test2)
	}

	if buyStockTwice(test3) != 5 {
		t.Errorf("Expected array %v to be 4", test3)
	}

	if buyStockTwice(test4) != 2 {
		t.Errorf("Expected array %v to be 2", test4)
	}

	if buyStockTwice(test5) != 9 {
		t.Errorf("Expected array %v to be 6", test5)
	}

	if buyStockTwice(test6) != 0 {
		t.Errorf("Expected array %v to be 0", test6)
	}
}

func TestBuyStock(t *testing.T) {
	test1 := []int{3, 3, 1, 0, 2, 0, 1}
	test2 := []int{3, 2, 0, 0, 2, 0, 1, 5}
	test3 := []int{5, 3, 1, 5, 2, 0, 1}
	test4 := []int{5, 3, 2, 4, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	test6 := []int{7, 6, 5, 4, 3, 2}
	if buyStock(test1) != 2 {
		t.Errorf("Expected array %v to be 2", test1)
	}
	if buyStock(test2) != 5 {
		t.Errorf("Expected array %v to be 5", test2)
	}

	if buyStock(test3) != 4 {
		t.Errorf("Expected array %v to be 4", test3)
	}

	if buyStock(test4) != 2 {
		t.Errorf("Expected array %v to be 2", test4)
	}

	if buyStock(test5) != 6 {
		t.Errorf("Expected array %v to be 6", test5)
	}

	if buyStock(test6) != 0 {
		t.Errorf("Expected array %v to be 0", test6)
	}
}

func TestCheckPermutation(t *testing.T) {
	test1 := checkPermutation("blah", "halb")
	test2 := checkPermutation("baleen", "neelab")
	test3 := checkPermutation("cactus", "flower")
	test4 := checkPermutation("ballsy", "heptagon")
	test5 := checkPermutation("winky face", "kin wafeyc")

	if test1 != true {
		t.Errorf("Expected checkPermutation to detect that blah and halb are permutations of eachother")
	}

	if test2 != true {
		t.Errorf("Expected checkPermutation to detect that baleen and neelab are permutations of eachother")
	}

	if test3 != false {
		t.Errorf("Expected checkPermutation to detect that cactus and flower are not permutations of eachother")
	}

	if test4 != false {
		t.Errorf("Expected checkPermutation to detect that ballsy and heptagon are not permutations of eachother")
	}

	if test5 != true {
		t.Errorf("Expected checkPermutation to detect that winky face and kin wafeyc are permutations of eachother")
	}
}
