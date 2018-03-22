package goExercises
 
import (
"fmt"
"testing"
"github.com/google/go-cmp/cmp"
)
 
func TestTwoSum(t *testing.T) {
	test1 := twoSum([]int{1, 3, 5, -3}, 0)
	ans1 := []int{1, 3}
	test2 := twoSum([]int{1, 3, 5}, 0)
	ans2 := []int{}
	test3 := twoSum([]int{1, 2, 3, 4, 5, 6, -4}, 0)
	ans3 := []int{3, 6}

	if !equalityOfSlices(test1, ans1) {
		t.Errorf("Expected twoSum of [1,3,5,-3] to be [1 3], not %v", test1)
	}

	if !equalityOfSlices(test2, ans2) {
		t.Errorf("Expected twoSum of [1,3,5] to be [], not %v", test2)
	}

	if !equalityOfSlices(test3, ans3) {
		t.Errorf("Expected twoSum of [1, 2, 3, 4, 5, 6, -4] to be [3, 6], not %v", test3)
	}
}

func TestTwoSumOptimized(t *testing.T) {
	test1 := twoSumOptimized([]int{1, 3, 5, -3}, 0)
	ans1 := []int{1, 3}
	test2 := twoSumOptimized([]int{1, 3, 5}, 0)
	ans2 := []int{}
	test3 := twoSumOptimized([]int{1, 2, 3, 4, 5, 6, -4}, 0)
	ans3 := []int{3, 6}

	if !equalityOfSlices(test1, ans1) {
		t.Errorf("Expected twoSum of [1,3,5,-3] to be [1 3], not %v", test1)
	}

	if !equalityOfSlices(test2, ans2) {
		t.Errorf("Expected twoSum of [1,3,5] to be [], not %v", test2)
	}

	if !equalityOfSlices(test3, ans3) {
		t.Errorf("Expected twoSum of [1, 2, 3, 4, 5, 6, -4] to be [3, 6], not %v", test3)
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		twoSum([]int{1, 2, 3, 4, 5, 6, -4, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 11, 12, 14, 14, 14, 15,
			16, 16, 176, 17, 17, 8, 4, 87, 89, 9, 5, 4, 56, 7, 89, 8, 6, 43, 3, 34, 56, 7, 8, 9, 90,
			7, 1, 37, 79, 164, 155, 32, 87, 39, 113, 15, 18, 78, 175, 140, 200, 4, 160, 97, 191, 100, 91, 20, 69, 198,
			2, 123, 134, 10, 141, 13, 12, 43, 47, 3, 177, 101, 179, 77, 182, 117, 116, 36, 103, 51, 154, 162, 128, 30,
			3, 48, 123, 134, 109, 41, 17, 159, 49, 136, 16, 130, 141, 29, 176, 2, 190, 66, 153, 157, 70, 114, 65, 173, 104, 194, 54,
			4, 91, 171, 118, 125, 158, 76, 107, 18, 73, 140, 42, 193, 127, 100, 84, 121, 60, 81, 99, 80, 150, 55, 1, 35, 23, 93,
			5, 193, 156, 102, 118, 175, 39, 124, 119, 19, 99, 160, 75, 20, 112, 37, 23, 145, 135, 146, 73, 35,
			6, 155, 56, 52, 120, 131, 160, 124, 119, 14, 196, 144, 25, 75, 76, 166, 35, 87, 26, 20, 32, 23,
			7, 156, 185, 178, 79, 27, 52, 144, 107, 78, 22, 71, 26, 31, 15, 56, 76, 112, 39, 8, 113, 93,
			8, 185, 155, 171, 178, 108, 64, 164, 53, 140, 25, 100, 133, 9, 52, 191, 46, 20, 150, 144, 39, 62, 131, 42, 119, 127, 31, 7,
			9, 91, 155, 8, 160, 107, 132, 195, 26, 20, 133, 39, 76, 100, 78, 122, 127, 38, 156, 191, 196, 115,
			10, 190, 184, 154, 49, 2, 182, 173, 170, 161, 47, 189, 101, 153, 50, 30, 109, 177, 148, 179, 16, 163, 116, 13, 90, 185,
			11, 123, 134, 163, 41, 12, 28, 130, 13, 101, 83, 77, 109, 114, 21, 82, 88, 74, 24, 94, 48, 33,
			12, 161, 109, 169, 21, 24, 36, 65, 50, 2, 101, 159, 148, 54, 192, 88, 47, 11, 142, 43, 70, 182}, 0)
	}
}

func BenchmarkTwoSumOptimized(b *testing.B) {
	for n := 0; n < b.N; n++ {
		twoSumOptimized([]int{1, 2, 3, 4, 5, 6, -4, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 11, 12, 14, 14, 14, 15,
			16, 16, 176, 17, 17, 8, 4, 87, 89, 9, 5, 4, 56, 7, 89, 8, 6, 43, 3, 34, 56, 7, 8, 9, 90,
			7, 1, 37, 79, 164, 155, 32, 87, 39, 113, 15, 18, 78, 175, 140, 200, 4, 160, 97, 191, 100, 91, 20, 69, 198,
			2, 123, 134, 10, 141, 13, 12, 43, 47, 3, 177, 101, 179, 77, 182, 117, 116, 36, 103, 51, 154, 162, 128, 30,
			3, 48, 123, 134, 109, 41, 17, 159, 49, 136, 16, 130, 141, 29, 176, 2, 190, 66, 153, 157, 70, 114, 65, 173, 104, 194, 54,
			4, 91, 171, 118, 125, 158, 76, 107, 18, 73, 140, 42, 193, 127, 100, 84, 121, 60, 81, 99, 80, 150, 55, 1, 35, 23, 93,
			5, 193, 156, 102, 118, 175, 39, 124, 119, 19, 99, 160, 75, 20, 112, 37, 23, 145, 135, 146, 73, 35,
			6, 155, 56, 52, 120, 131, 160, 124, 119, 14, 196, 144, 25, 75, 76, 166, 35, 87, 26, 20, 32, 23,
			7, 156, 185, 178, 79, 27, 52, 144, 107, 78, 22, 71, 26, 31, 15, 56, 76, 112, 39, 8, 113, 93,
			8, 185, 155, 171, 178, 108, 64, 164, 53, 140, 25, 100, 133, 9, 52, 191, 46, 20, 150, 144, 39, 62, 131, 42, 119, 127, 31, 7,
			9, 91, 155, 8, 160, 107, 132, 195, 26, 20, 133, 39, 76, 100, 78, 122, 127, 38, 156, 191, 196, 115,
			10, 190, 184, 154, 49, 2, 182, 173, 170, 161, 47, 189, 101, 153, 50, 30, 109, 177, 148, 179, 16, 163, 116, 13, 90, 185,
			11, 123, 134, 163, 41, 12, 28, 130, 13, 101, 83, 77, 109, 114, 21, 82, 88, 74, 24, 94, 48, 33,
			12, 161, 109, 169, 21, 24, 36, 65, 50, 2, 101, 159, 148, 54, 192, 88, 47, 11, 142, 43, 70, 182}, 0)
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
 
