package goExercises

import "strings"

//return sum of n + n-1 + ... 1
func sumNums(n int) int {

	if n == 1 {
		return 1
	}

	return n + sumNums(n-1)
}
 
//return the longest word in sentence
func longestWord(sentence string) string {
	longestWord := ""
	wordCount := 0
	for _, word := range strings.Split(sentence, " ") {
		if len(word) > wordCount {
			longestWord = word
			wordCount = len(word)
		}
	}
	return longestWord
}
 
