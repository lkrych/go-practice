package goExercises

import "strings"

//Return the indices of the elements in the arr that sum to find
func twoSum(arr []int, find int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == find {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//create a more optimized version of the function
func twoSumOptimized(arr []int, find int) []int {
	iMap := make(map[int]int)
	for i := 0; i < len(arr)-1; i++ {
		iMap[arr[i]] = i
	}

	for i := 0; i < len(arr)-1; i++ {
		lookFor := find - arr[i]
		if v, ok := iMap[lookFor]; ok {
			return []int{i, v}
		}
	}
	return []int{}
}

//observe the benchmark code in the tests. Try to write this code after looking at it

//take an array of ints of length word, and return a string with each character of word in the order of idxArray
func scrambleString(word string, idxArray []int) string {
	newString := []string{}
	for i := 0; i < len(idxArray); i++ {
		idx := idxArray[i]
		newString = append(newString, string(word[idx]))
	}
	return strings.Join(newString, "")
}
