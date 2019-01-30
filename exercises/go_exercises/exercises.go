package goExercises

import (
	"fmt"
	"strings"
)

//return sorted arr using mergesort
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// mergeSort both halfs of array and then merge
	midIdx := len(arr) / 2
	left := mergeSort(arr[:midIdx])
	right := mergeSort(arr[midIdx:])
	return merge(left, right)
}

func merge(arr1, arr2 []int) []int {
	//iterate and compare each element to create sorted array
	sorted := []int{}
	for len(arr1) > 0 && len(arr2) > 0 {
		// if either array is empty
		// push contents of other array into sorted
		// and skip to next loop cycle which will exit
		el1 := arr1[0]
		el2 := arr2[0]
		if el1 >= el2 {
			sorted = append(sorted, el2)
			arr2 = arr2[1:]
		} else {
			sorted = append(sorted, el1)
			arr1 = arr1[1:]
		}
	}
	sorted = append(sorted, arr1...)
	sorted = append(sorted, arr2...)
	return sorted
}

//Given a string, find the length of the longest substring without repeating characters.
func longestSubString(sentence string) int {
	charMap := make(map[string]int)
	longestSub := 0
	currentSub := 0
	fmt.Println("sentence is", sentence)
	for idx, ch := range strings.Split(sentence, "") {
		fmt.Println("Current Char is", ch)
		fmt.Println("LongestSub is", longestSub)
		fmt.Println("CurrentSub is", currentSub)
		if prevIdx, ok := charMap[ch]; ok {
			currentSub = idx - prevIdx
		} else {
			currentSub++
		}
		charMap[ch] = idx
		if currentSub > longestSub {
			longestSub = currentSub
		}

	}
	return longestSub
}
