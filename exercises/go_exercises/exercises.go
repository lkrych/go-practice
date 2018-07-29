package goExercises

import (
	"strings"
)

//Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
func isUnique(word string) bool {
	wordMap := make(map[string]bool)
	for _, ch := range strings.Split(word, "") {
		if ok, _ := wordMap[ch]; ok {
			return false
		} else {
			wordMap[ch] = true
		}
	}
	return true
}

//return a sorted arr using bubblesort!
func bubbleSort(arr []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				sorted = false
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}

//return the number of repeated characters in word
func numRepeats(word string) int {
	repeat := 0
	wordMap := make(map[string]int)
	for _, ch := range strings.Split(word, "") {
		if val, ok := wordMap[ch]; ok {
			if val == 1 {
				repeat++
			}
			wordMap[ch] = val + 1
		} else {
			wordMap[ch] = 1
		}
	}

	return repeat
}
