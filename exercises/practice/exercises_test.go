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

func TestScrambleString(t *testing.T) {
	test1 := scrambleString("abcd", []int{3, 1, 2, 0})
	test2 := scrambleString("markov", []int{5, 3, 1, 4, 2, 0})

	if test1 != "dbca" {
		t.Errorf("Expected scrambleString of 'abcd' with inputs [3,1,2,0] to be 'dcba', not %v", test1)
	}

	if test2 != "vkaorm" {
		t.Errorf("Expected scrambleString of 'markov' with inputs [5,3,1,4,2,0] to be 'vkaorm', not %v", test2)
	}
}

func TestIsPrime(t *testing.T) {
	test1 := isPrime(2)
	test2 := isPrime(3)
	test3 := isPrime(4)
	test4 := isPrime(9)
	test5 := isPrime(227)

	if test1 != true {
		t.Errorf("Expected isPrime of 2 to be true, not false")
	}

	if test2 != true {
		t.Errorf("Expected isPrime of 3 to be true, not false")
	}

	if test3 != false {
		t.Errorf("Expected isPrime of 4 to be false, not true")
	}

	if test4 != false {
		t.Errorf("Expected isPrime of 9 to be false, not true")
	}

	if test5 != true {
		t.Errorf("Expected isPrime of 227 to be true, not false")
	}
}

func TestNthPrime(t *testing.T) {
	test1 := nthPrime(3)
	test2 := nthPrime(4)
	test3 := nthPrime(15)
	test4 := nthPrime(28)
	test5 := nthPrime(77)

	if test1 != 5 {
		t.Errorf("Expected nthPrime of 3 to be 5, not %v", test1)
	}

	if test2 != 7 {
		t.Errorf("Expected nthPrime of 4 to be 7, not %v", test2)
	}

	if test3 != 47 {
		t.Errorf("Expected nthPrime of 15 to be 47, not %v", test3)
	}

	if test4 != 107 {
		t.Errorf("Expected nthPrime of 28 to be 107, not %v", test4)
	}

	if test5 != 389 {
		t.Errorf("Expected nthPrime of 77 to be 389, not %v", test5)
	}
}

func TestLongestPalindrome(t *testing.T) {
	test1 := longestPalindrome("abcbd")
	test2 := longestPalindrome("abba")
	test3 := longestPalindrome("abcbdeffe")

	if test1 != "bcb" {
		t.Errorf("Expected longestPalindrome of abcbd to be bcb, not %v", test1)
	}

	if test2 != "abba" {
		t.Errorf("Expected longestPalindrome of abba to be abba, not %v", test2)
	}

	if test3 != "effe" {
		t.Errorf("Expected longestPalindrome of abcbdeffe to be effe, not %v", test3)
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

func TestCaesarCipher(t *testing.T) {
	test1 := caesarCipher(3, "abc")
	test2 := caesarCipher(3, "abc xyz")
	test3 := caesarCipher(29, "abc xyz")

	if test1 != "def" {
		t.Errorf("Expected caesarCipher with a shift of 3 and input of abc to be def, not %v", test1)
	}

	if test2 != "def abc" {
		t.Errorf("Expected caesarCipher with a shift of 3 and input of abc xyz to be def abc, not %v", test2)
	}

	if test3 != "def abc" {
		t.Errorf("Expected caesarCipher with a shift of 29 and input of abc xyz to be def abc, not %v", test3)
	}
}

func TestNumRepeats(t *testing.T) {
	test1 := numRepeats("abdbc")
	test2 := numRepeats("aaa")
	test3 := numRepeats("abab")
	test4 := numRepeats("cadac")
	test5 := numRepeats("abcde")

	if test1 != 1 {
		t.Errorf("Expected numRepeats of abdbc to be 1, not %v", test1)
	}

	if test2 != 1 {
		t.Errorf("Expected numRepeats of aaa to be 1, not %v", test2)
	}

	if test3 != 2 {
		t.Errorf("Expected numRepeats of abab to be 2, not %v", test3)
	}

	if test4 != 2 {
		t.Errorf("Expected numRepeats of cadac to be 2, not %v", test4)
	}

	if test5 != 0 {
		t.Errorf("Expected numRepeats of abcde to be 0, not %v", test5)
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

func TestBinarySearch(t *testing.T) {
	test1 := binarySearch([]int{0, 1, 2, 3, 4, 5}, 3)
	test2 := binarySearch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9)
	test3 := binarySearch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)
	test4 := binarySearch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
	test5 := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0)

	if test1 != 3 {
		t.Errorf("Expected binarySearch to find the index of 3 at 4, not %v", test1)
	}
	if test2 != 9 {
		t.Errorf("Expected binarySearch to find the index of 9 at 10, not %v", test2)
	}

	if test3 != 0 {
		t.Errorf("Expected binarySearch to find the index of 0 at 0, not %v", test3)
	}

	if test4 != -1 {
		t.Errorf("Expected binarySearch to find the index of 10 at -1, not %v", test4)
	}

	if test5 != -1 {
		t.Errorf("Expected binarySearch to find the index of 0 at -1, not %v", test5)
	}
}

func TestBubbleSort(t *testing.T) {
	test1 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3})
	test2 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3, 1, 6, 10})
	test3 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3, 3, 3, 3})
	test4 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3, 100, 102, 6, 1, 10})
	test5 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3, 11, 10, 6, 1})
	ans1 := []int{2, 3, 4, 5, 7, 8, 9}
	ans2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans3 := []int{2, 3, 3, 3, 3, 4, 5, 7, 8, 9}
	ans4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 102}
	ans5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test1, ans1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test2, ans2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test3, ans3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test4, ans4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test5, ans5)
	}
}

func TestDigitalRoot(t *testing.T) {
	test1 := digitalRoot(10)
	test2 := digitalRoot(83)
	test3 := digitalRoot(123)
	test4 := digitalRoot(789)
	test5 := digitalRoot(452982)

	if test1 != 1 {
		t.Errorf("Expected digitalRoot of 10 to be 1, not %v", test1)
	}
	if test2 != 2 {
		t.Errorf("Expected digitalRoot of 83 to be 2, not %v", test2)
	}

	if test3 != 6 {
		t.Errorf("Expected digitalRoot of 123 to be 6, not %v", test3)
	}

	if test4 != 6 {
		t.Errorf("Expected digitalRoot of 789 to be 6, not %v", test4)
	}

	if test5 != 3 {
		t.Errorf("Expected digitalRoot of 452982 to be 3, not %v", test5)
	}
}

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

func TestExponent(t *testing.T) {
	test1 := exponent(11, 3)
	test2 := exponent(10, 2)
	test3 := exponent(2, 16)
	test4 := exponent(7, 6)
	test5 := exponent(9, 3)
	test6 := exponent(5, -3)

	if test1 != 1331 {
		t.Errorf("Expected exponent to raise 11 to 3 to be 1331, not %v", test1)
	}
	if test2 != 100 {
		t.Errorf("Expected exponent to raise 10 to 2 to be 100, not %v", test2)
	}

	if test3 != 65536 {
		t.Errorf("Expected exponent to raise 2 to 16 to be 65536, not %v", test3)
	}

	if test4 != 117649 {
		t.Errorf("Expected exponent to raise 7 to 6 to be 117649, not %v", test4)
	}

	if test5 != 729 {
		t.Errorf("Expected exponent to raise 9 to 3 to be 729, not %v", test5)
	}

	if test6 != 1/125 {
		t.Errorf("Expected exponent to raise 5 to -3 to be 1/125, not %v", test6)
	}
}

func TestSqrt(t *testing.T) {
	test1 := sqrt(2)
	test2 := sqrt(4)
	test3 := sqrt(9)
	test4 := sqrt(8)
	test5 := sqrt(10)

	if test1 < 1.41 || test1 >= 1.42 {
		t.Errorf("Expected sqrt of 11 to be about 1.41, not %v", test1)
	}
	if test2 != 2 {
		t.Errorf("Expected sqrt of 4 to be 2, not %v", test2)
	}

	if test3 != 3 {
		t.Errorf("Expected sqrt of 9 to be 3, not %v", test3)
	}

	if test4 < 2.82 || test4 >= 2.83 {
		t.Errorf("Expected sqrt of 8 to be about 2.82, not %v", test4)
	}

	if test5 < 3.16 || test5 >= 3.17 {
		t.Errorf("Expected sqrt of 10 to be about 3.16, not %v", test5)
	}
}

func TestPartitionEvenOdd(t *testing.T) {
	test1 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3})
	test2 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3, 1, 6, 10})
	test3 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3, 3, 3, 3})
	test4 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3, 100, 102, 6, 1, 10})
	test5 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3, 11, 10, 6, 1})
	ans1 := []int{2, 4, 8, 3, 5, 7, 9}
	ans2 := []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9}
	ans3 := []int{2, 4, 8, 3, 3, 3, 3, 5, 7, 9}
	ans4 := []int{2, 4, 6, 8, 10, 100, 102, 1, 3, 5, 7, 9}
	ans5 := []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9, 11}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans1, test1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans2, test2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans3, test3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans4, test4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans5, test5)
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

func TestAdvancingArray(t *testing.T) {
	test1 := []int{3, 3, 1, 0, 2, 0, 1}
	test2 := []int{3, 2, 0, 0, 2, 0, 1}
	test3 := []int{5, 3, 1, 0, 2, 0, 1}
	test4 := []int{0, 3, 1, 0, 2, 0, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	if advancingArray(test1) != true {
		t.Errorf("Expected array %v to be true", test1)
	}
	if advancingArray(test2) != false {
		t.Errorf("Expected array %v to be false", test2)
	}

	if advancingArray(test3) != true {
		t.Errorf("Expected array %v to be true", test3)
	}

	if advancingArray(test4) != false {
		t.Errorf("Expected array %v to be false", test4)
	}

	if advancingArray(test5) != true {
		t.Errorf("Expected array %v to be true", test5)
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

func TestReverseCaptcha(t *testing.T) {
	test1 := reverseCaptcha([]int{1, 1, 2, 2})
	test2 := reverseCaptcha([]int{1, 1, 1, 1})
	test3 := reverseCaptcha([]int{1, 2, 3, 4})
	test4 := reverseCaptcha([]int{9, 1, 2, 9})
	test5 := reverseCaptcha([]int{5, 4, 2, 4, 2, 2, 7, 7, 3, 2, 5})
	test6 := adventCaptcha()

	if test1 != 3 {
		t.Errorf("Expected answer to be 3, not %v", test1)
	}
	if test2 != 4 {
		t.Errorf("Expected answer to be 4, not %v", test2)
	}

	if test3 != 0 {
		t.Errorf("Expected answer to be 0, not %v", test3)
	}

	if test4 != 9 {
		t.Errorf("Expected answer to be 9, not %v", test4)
	}

	if test5 != 14 {
		t.Errorf("Expected answer to be 14, not %v", test5)
	}

	if test6 != 1102 {
		t.Errorf("Expected answer to be 1102, not %v", test6)
	}

}

func TestCaptcha(t *testing.T) {
	adventCaptcha()
}

func TestCheckSum(t *testing.T) {
	test1 := checkSum([][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}})
	test2 := checkSum([][]int{{1, 2, 3, 4}, {1, 2, 3}, {1, 2, 3}})
	test3 := checkSum([][]int{{1, 2, 3, 10}, {1, 2, 3, 12}, {1, 2, 3, 15}})
	test4 := checkSum([][]int{{3, 22, 124}, {3, 70}, {1, 2, 5}})
	test5 := checkSum([][]int{{1, 2, 10}, {1, 11, 3}, {500, 2, 3}})
	test6 := checkSum(checkSumAdvent())
	test7 := checkSumEven(checkSumAdvent())
	test8 := checkSumEven([][]int{{5, 9, 2, 8}, {9, 4, 7, 3}, {3, 8, 6, 5}})

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

	if test6 != 53978 {
		t.Errorf("Expected answer to be 53978, not %v", test6)
	}

	if test7 != 314 {
		t.Errorf("Expected answer to be 314, not %v", test7)
	}

	if test8 != 9 {
		t.Errorf("Expected answer to be 9, not %v", test8)
	}
}

func TestSpiral(t *testing.T) {
	test1 := spiralMemory(325489)
	test2 := spiralMemory(12)
	test3 := spiralMemory(23)
	test4 := spiralMemory(1024)

	if test1 != 6 {
		t.Errorf("Expected answer to be 6, not %v", test1)
	}

	if test2 != 3 {
		t.Errorf("Expected answer to be 3, not %v", test2)
	}

	if test3 != 2 {
		t.Errorf("Expected answer to be 2, not %v", test3)
	}

	if test4 != 31 {
		t.Errorf("Expected answer to be 31, not %v", test4)
	}
}

func TestPassPhrase(t *testing.T) {
	test1 := passPhrase("doody goody woody")
	test2 := passPhrase("doody doody woody")
	test3 := passPhrase("woody goody woody")
	test4 := passPhrase("It is not I who is the Creature!")
	test5 := passPhrase("Hello, my name is Bill.")
	test6 := passPhraseAdvent()

	if test1 != true {
		t.Errorf("Expected answer to be true, not %v", test1)
	}

	if test2 != false {
		t.Errorf("Expected answer to be false, not %v", test2)
	}

	if test3 != false {
		t.Errorf("Expected answer to be false, not %v", test3)
	}

	if test4 != false {
		t.Errorf("Expected answer to be false, not %v", test4)
	}

	if test5 != true {
		t.Errorf("Expected answer to be true, not %v", test5)
	}

	if test6 != 325 {
		t.Errorf("Expected answer to be 325, not %v", test6)
	}
}

func TestPassPhraseAnagram(t *testing.T) {
	test1 := passPhraseAnagram("abcde fghij")
	test2 := passPhraseAnagram("abcde xyz ecdab")
	test3 := passPhraseAnagram("a ab abc abd abf abj")
	test4 := passPhraseAnagram("iiii oiii ooii oooi oooo")
	test5 := passPhraseAnagram("oiii ioii iioi iiio")
	test6 := passPhraseAnagramAdvent()

	if test1 != true {
		t.Errorf("Expected answer to be true, not %v", test1)
	}

	if test2 != false {
		t.Errorf("Expected answer to be false, not %v", test2)
	}

	if test3 != true {
		t.Errorf("Expected answer to be true, not %v", test3)
	}

	if test4 != true {
		t.Errorf("Expected answer to be false, not %v", test4)
	}

	if test5 != false {
		t.Errorf("Expected answer to be false, not %v", test5)
	}

	if test6 != 119 {
		t.Errorf("Expected answer to be 6, not %v", test6)
	}
}

func TestJumpMaze(t *testing.T) {
	test1 := jumpMaze([]int{0, 3, 0, 1, -3})
	test2 := jumpMaze([]int{1, 3, 0, 1, -3})
	test3 := jumpMaze([]int{0, 2, 0, 1, -3})
	test4 := jumpMazeAdvent()

	if test1 != 5 {
		t.Errorf("Expected answer to be 5, not %v", test1)
	}

	if test2 != 4 {
		t.Errorf("Expected answer to be 4, not %v", test2)
	}

	if test3 != 10 {
		t.Errorf("Expected answer to be 10, not %v", test3)
	}

	if test4 != 388611 {
		t.Errorf("Expected answer to be 388611, not %v", test4)
	}
}

func TestJumpMazeStrange(t *testing.T) {
	test1 := jumpMazeStrange([]int{0, 3, 0, 1, -3})
	test2 := jumpMazeStrange([]int{1, 3, 0, 1, -3})
	test3 := jumpMazeStrange([]int{0, 2, 0, 1, -3})
	test4 := jumpMazeStrangeAdvent()

	if test1 != 10 {
		t.Errorf("Expected answer to be 10, not %v", test1)
	}

	if test2 != 9 {
		t.Errorf("Expected answer to be 9, not %v", test2)
	}

	if test3 != 10 {
		t.Errorf("Expected answer to be 10, not %v", test3)
	}

	if test4 != 388611 {
		t.Errorf("Expected answer to be 388611, not %v", test4)
	}
}

func TestMemoryBank(t *testing.T) {
	test1 := memoryBank([]int{0, 2, 7, 0})
	test2 := memoryBank([]int{0, 1, 0, 0})
	test3 := memoryBankAdvent()

	if test1 != 5 {
		t.Errorf("Expected answer to be 5, not %v", test1)
	}

	if test2 != 4 {
		t.Errorf("Expected answer to be 4, not %v", test2)
	}

	if test3 != 11137 {
		t.Errorf("Expected answer to be 11137, not %v", test3)
	}
}
