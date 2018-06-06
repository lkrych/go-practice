package goExercises

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemoveDups(t *testing.T) {
	test1 := removeDups([]int{1, 1, 1, 2, 3, 4, 4, 4, 5})
	test2 := removeDups([]int{1, 1, 1, 2, 33, 102, 102})
	test3 := removeDups([]int{1, 1, 1, 2, 3, 3, 4, 5, 6, 7, 8, 8, 9, 9})
	test4 := removeDups([]int{1, 1, 1, 2, 3, 3, 3, 3, 3})
	test5 := removeDups([]int{1, 1, 1, 2, 3, 3, 4, 5, 5})
	ans1 := []int{1, 2, 3, 4, 5}
	ans2 := []int{1, 2, 33, 102}
	ans3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans4 := []int{1, 2, 3}
	ans5 := []int{1, 2, 3, 4, 5}
	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected array %v to be %v", test1, ans1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected array %v to be %v", test2, ans2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected array %v to be %v", test4, ans3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected array %v to be %v", test5, ans4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected array %v to be %v", test5, ans5)
	}
}

func TestBaseConverter(t *testing.T) {
	test1 := baseConverter(5, 2)
	test2 := baseConverter(1239449, 16)
	test3 := baseConverter(112, 2)
	test4 := baseConverter(112, 16)

	if test1 != "101" {
		t.Errorf("Expected baseConverter of 5 to base 2 to equal 101, not %v", test1)
	}
	if test2 != "12e999" {
		t.Errorf("Expected baseConverter of 1239449 to base 16 to be 12e999, not %v", test2)
	}

	if test3 != "1110000" {
		t.Errorf("Expected baseConverter of 112 to base 2 to be 1110000, not %v", test3)
	}

	if test4 != "70" {
		t.Errorf("Expected baseConverter of 112 to base 16 to be 70, not %v", test4)
	}
}

func TestMultiplyArbitrary(t *testing.T) {
	test1 := multiplyArb([]int{1, 9, 3, 7, 0, 7, 7, 2, 1}, []int{-7, 6, 1, 8, 3, 8, 2, 5, 7, 2, 8, 7})
	test2 := multiplyArb([]int{1, 2, 9, 8, 2, 3, 4, 1, 3, 5, 0}, []int{4, 5})
	test3 := multiplyArb([]int{1, 2, 5, 3, 2, 4, 6, 7, 8, 2}, []int{-4, 0, 0, 2, 3})
	test4 := multiplyArb([]int{5, 3, 6, 3, 7, 9, 5, 0}, []int{3, 5, 2})
	test5 := multiplyArb([]int{1, 2}, []int{5})
	test6 := multiplyArb([]int{1, 2, 3}, []int{9, 8, 7})
	ans1 := []int{-1, 4, 7, 5, 7, 3, 9, 5, 2, 5, 8, 9, 6, 7, 6, 4, 1, 2, 9, 2, 7}
	ans2 := []int{5, 8, 4, 2, 0, 5, 3, 6, 0, 7, 5, 0}
	ans3 := []int{-5, 0, 1, 5, 8, 6, 9, 5, 9, 5, 5, 9, 8, 6}
	ans4 := []int{1, 8, 8, 8, 0, 5, 5, 8, 4, 0, 0}
	ans5 := []int{6, 0}
	ans6 := []int{1, 2, 1, 4, 0, 1}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans1, test1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans2, test2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans3, test3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans4, test4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans5, test5)
	}

	if !cmp.Equal(test6, ans6) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans6, test6)
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
