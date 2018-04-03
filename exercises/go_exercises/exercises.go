package goExercises

import "strings"

//return if the word is a palindrome
func isPalindrome(word string) bool {
	splitWord := strings.Split(word, "")
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if splitWord[i] != splitWord[j] {
			return false
		}
	}
	return true
}
