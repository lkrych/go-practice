package goExercises
 
import "testing"

func TestPalindromePermutation(t *testing.T) {
	test1 := palindromePermutation("tact coa")
	test2 := palindromePermutation("aaa")
	test3 := palindromePermutation("abab")
	test4 := palindromePermutation("cadak")
	test5 := palindromePermutation("abcde")

	if test1 != true {
		t.Errorf("Expected palindromePermutation of tact coa to be true, not %v", test1)
	}

	if test2 != true {
		t.Errorf("Expected palindromePermutation of aaa to be true, not %v", test2)
	}

	if test3 != true {
		t.Errorf("Expected palindromePermutation of abab to be true, not %v", test3)
	}

	if test4 != false {
		t.Errorf("Expected palindromePermutation of cadac to be false, not %v", test4)
	}

	if test5 != false {
		t.Errorf("Expected palindromePermutation of abcde to be false, not %v", test5)
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
 
