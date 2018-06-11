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

func TestTimeConversion(t *testing.T) {
	test1 := timeConversion(15)
	test2 := timeConversion(150)
	test3 := timeConversion(360)

	if test1 != "0:15" {
		t.Errorf("Expected TimeConversion of 15 to be 0:15, not %v", test1)
	}

	if test2 != "2:30" {
		t.Errorf("Expected TimeConversion of 150 to be 2:30, not %v", test2)
	}

	if test3 != "6:00" {
		t.Errorf("Expected TimeConversion of 360 to be 6:00, not %v", test3)
	}
}

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
