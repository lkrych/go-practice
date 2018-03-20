package goExercises

import (
	"strings"
)

//return sorted arr using quicksort
func quickSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	pivot := arr[0]
	leftArr := []int{}
	rightArr := []int{}

	for i := 1; i < len(arr); i++ {
		if arr[i] <= pivot {
			leftArr = append(leftArr, arr[i])
		} else {
			rightArr = append(rightArr, arr[i])
		}
	}
	leftSorted := quickSort(leftArr)
	rightSorted := quickSort(rightArr)

	leftSorted = append(leftSorted, pivot)
	return append(leftSorted, rightSorted...)
}

//return true if the letter a is within two spaces of the letter z
func nearbyAZ(word string) bool {
	found := 0
	foundA := false
	chars := strings.Split(word, "")
	for i := 0; i < len(chars); i++ {
		if chars[i] == "a" {
			found = 3
			continue
		}
		if found > 0 {
			if chars[i] == "z" {
				foundA = true
			}
			found--
		}
	}
	return foundA
}

//return the GCF of the two inputs
func greatestCommonFactor(n1 int, n2 int) int {
	var greater, smaller int
	if n1 >= n2 {
		greater = n1
		smaller = n2
	} else {
		greater = n2
		smaller = n1
	}

	for i := smaller; i > 0; i-- {
		if greater%i == 0 && smaller%i == 0 {
			return i
		}
	}
	return 1
}

//In O(n) time, and using O(1) space, remove duplicates from an array
func removeDups(arr []int) []int {
	nextUnduplicatedIdx := 1
	//O(n) time complexity
	for i := 1; i < len(arr); i++ {
		//if the adjacent elements are not equal, assign the element at i to our nextUnduplicatedIdx
		if arr[nextUnduplicatedIdx-1] != arr[i] {
			arr[nextUnduplicatedIdx] = arr[i]
			nextUnduplicatedIdx++
		}
	}
	return arr[:nextUnduplicatedIdx]
}

func removeDupsAlt(arr []int) []int {
	h := make(map[int]bool)

	//O(n) time complexity
	for i := 0; i < len(arr); i++ {
		h[arr[i]] = true
	}
	//O(n) space complexity
	keys := make([]int, len(h))

	i := 0
	//O(n) time complexity
	for k := range h {
		keys[i] = k
		i++
	}
	//O(nlogn) time complexity
	return quickSort(keys)

}
