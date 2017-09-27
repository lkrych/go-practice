package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestIsPalindrome(t *testing.T) {
	test1 := isPalindrome("abc")
	test2 := isPalindrome("racecar")
	test3 := isPalindrome("z")

	if test1 != false {
		t.Errorf("Expected isPalindrome of abc to be false, not %v", test1)
	}

	if test2 != true {
		t.Errorf("Expected isPalindrome of racecar to be true, not %v", test2)
	}

	if test3 != true {
		t.Errorf("Expected isPalindrome of z to be true, not %v", test3)
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
		t.Errorf("Expected nearbyAZ of baz to be false, not %v", test1)
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

func equalityOfSlices(a []int, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			fmt.Println("INSIDE EQUALITY", v, "!=", b[i])
			return false
		}
	}
	return true
}

func TestTwoSum(t *testing.T) {
	test1 := twoSum([]int{1, 3, 5, -3})
	ans1 := []int{1, 3}
	test2 := twoSum([]int{1, 3, 5})
	ans2 := []int{}
	test3 := twoSum([]int{1, 2, 3, 4, 5, 6, -4})
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
	test1 := twoSumOptimized([]int{1, 3, 5, -3})
	ans1 := []int{1, 3}
	test2 := twoSumOptimized([]int{1, 3, 5})
	ans2 := []int{}
	test3 := twoSumOptimized([]int{1, 2, 3, 4, 5, 6, -4})
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
			12, 161, 109, 169, 21, 24, 36, 65, 50, 2, 101, 159, 148, 54, 192, 88, 47, 11, 142, 43, 70, 182})
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
			12, 161, 109, 169, 21, 24, 36, 65, 50, 2, 101, 159, 148, 54, 192, 88, 47, 11, 142, 43, 70, 182})
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

func TestThirdGreatest(t *testing.T) {
	test1 := thirdGreatest([]int{5, 3, 7})
	test2 := thirdGreatest([]int{5, 3, 4, 7})
	test3 := thirdGreatest([]int{2, 3, 4, 7})

	if test1 != 3 {
		t.Errorf("Expected thirdGreatest of [5, 3, 7] to be 3, not %v", test1)
	}

	if test2 != 4 {
		t.Errorf("Expected thirdGreatest of [5, 3, 4, 7] to be 4, not %v", test2)
	}

	if test3 != 3 {
		t.Errorf("Expected thirdGreatest of [2, 3, 4, 7] to be 3, not %v", test3)
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

func TestCapitalizeWords(t *testing.T) {
	test1 := capitalizeWords("this is a sentence")
	test2 := capitalizeWords("mike bloomfield")
	test3 := capitalizeWords("lawrence of arabia")

	if test1 != "This Is A Sentence" {
		t.Errorf("Expected capitalizeWords of this is a sentence to be This Is A Sentence, not %v", test1)
	}

	if test2 != "Mike Bloomfield" {
		t.Errorf("Expected capitalizeWords of mike bloomfield to be Mike Bloomfield, not %v", test2)
	}

	if test3 != "Lawrence Of Arabia" {
		t.Errorf("Expected capitalizeWords of lawrence of arabia to be Lawrence Of Arabia, not %v", test3)
	}
}
