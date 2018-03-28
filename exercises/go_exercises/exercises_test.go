package goExercises
 
import (
"fmt"
"testing"
"github.com/google/go-cmp/cmp"
)
 
func TestMergeSort(t *testing.T) {
	test1 := mergeSort([]int{7, 5, 2, 4, 8, 9, 3})
	test2 := mergeSort([]int{7, 5, 2, 4, 8, 9, 3, 1, 6, 10})
	test3 := mergeSort([]int{7, 5, 2, 4, 8, 9, 3, 3, 3, 3})
	test4 := mergeSort([]int{7, 5, 2, 4, 8, 9, 3, 100, 102, 6, 1, 10})
	test5 := mergeSort([]int{7, 5, 2, 4, 8, 9, 3, 11, 10, 6, 1})
	ans1 := []int{2, 3, 4, 5, 7, 8, 9}
	ans2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans3 := []int{2, 3, 3, 3, 3, 4, 5, 7, 8, 9}
	ans4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 102}
	ans5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected mergeSort to sort the array %v, instead you got %v", ans1, test1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected mergeSort to sort the array %v, instead you got %v", ans2, test2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected mergeSort to sort the array %v, instead you got %v", ans3, test3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected mergeSort to sort the array %v, instead you got %v", ans4, test4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected mergeSort to sort the array %v, instead you got %v", ans5, test5)
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
 
