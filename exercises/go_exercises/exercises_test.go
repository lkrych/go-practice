package goExercises

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQuickSort(t *testing.T) {
	test1 := quickSort([]int{7, 5, 2, 4, 8, 9, 3})
	test2 := quickSort([]int{7, 5, 2, 4, 8, 9, 3, 1, 6, 10})
	test3 := quickSort([]int{7, 5, 2, 4, 8, 9, 3, 3, 3, 3})
	test4 := quickSort([]int{7, 5, 2, 4, 8, 9, 3, 100, 102, 6, 1, 10})
	test5 := quickSort([]int{7, 5, 2, 4, 8, 9, 3, 11, 10, 6, 1})
	ans1 := []int{2, 3, 4, 5, 7, 8, 9}
	ans2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans3 := []int{2, 3, 3, 3, 3, 4, 5, 7, 8, 9}
	ans4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 102}
	ans5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected quickSort to sort the array %v, instead you got %v", ans1, test1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected quickSort to sort the array %v, instead you got %v", ans2, test2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected quickSort to sort the array %v, instead you got %v", ans3, test3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected quickSort to sort the array %v, instead you got %v", ans4, test4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected quickSort to sort the array %v, instead you got %v", ans5, test5)
	}
}

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
