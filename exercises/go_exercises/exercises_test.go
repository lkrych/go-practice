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

func TestNearbyAZ(t *testing.T) {
	test1 := nearbyAZ("baz")
	test2 := nearbyAZ("abz")
	test3 := nearbyAZ("abcz")
	test4 := nearbyAZ("a")
	test5 := nearbyAZ("z")
	test6 := nearbyAZ("za")

	if test1 != true {
		t.Errorf("Expected nearbyAZ of baz to be true, not %v", test1)
	}

	if test2 != true {
		t.Errorf("Expected nearbyAZ of abz to be true, not %v", test2)
	}

	if test3 != true {
		t.Errorf("Expected nearbyAZ of abcz to be true, not %v", test3)
	}

	if test4 != false {
		t.Errorf("Expected nearbyAZ of a to be false, not %v", test4)
	}

	if test5 != false {
		t.Errorf("Expected nearbyAZ of z to be false, not %v", test5)
	}

	if test6 != false {
		t.Errorf("Expected nearbyAZ of za to be false, not %v", test6)
	}
}

func TestGreatestCommonFactor(t *testing.T) {
	test1 := greatestCommonFactor(3, 9)
	test2 := greatestCommonFactor(16, 24)
	test3 := greatestCommonFactor(7, 38)

	if test1 != 3 {
		t.Errorf("Expected greatestCommonFactor of 3 and 9 to be 3, not %v", test1)
	}

	if test2 != 8 {
		t.Errorf("Expected greatestCommonFactor of 16 and 24 to be 8, not %v", test2)
	}

	if test3 != 1 {
		t.Errorf("Expected greatestCommonFactor of 7 and 38 to be 1, not %v", test3)
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
