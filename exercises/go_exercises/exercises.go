package goExercises

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// calculate the checksum of the multi-dimensional array
// For each row, determine the difference between the largest value and the smallest value;
// the checksum is the sum of all of these differences.
func checkSum(multiArr [][]int) int {
	checkSum := 0
	smallest := 20000
	largest := 0
	for _, arr := range multiArr {
		smallest = 20000
		largest = 0
		for _, el := range arr {
			if el > largest {
				largest = el
			}
			if el < smallest {
				smallest = el
			}
		}
		checkSum += (largest - smallest)
		//add difference between max and min to partial
	}
	return checkSum
}

func checkSumAdvent() int {
	//read in the file ../practice/checksum_input and find the checksum for it
	content, err := ioutil.ReadFile("../practice/advent_input/checksum_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	arrays := strings.Split(string(content), "\n")
	multiArr := make([][]int, len(arrays))

	for idx, arr := range arrays {
		splitArr := strings.Split(arr, "\t")
		intArr := make([]int, len(splitArr))
		for i, el := range splitArr {
			if el == "" {
				continue
			}
			intEl, err := strconv.Atoi(el)
			if err != nil {
				log.Fatal(err)
			}
			intArr[i] = intEl
		}
		multiArr[idx] = intArr
	}
	return checkSum(multiArr)
}

//return if the word is a palindrome
func isPalindrome(word string) bool {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}

//return index of toFind if it exists in arr, otherwise return -1
func binarySearch(arr []int, toFind int) int {
	//base case
	if len(arr) == 0 {
		return -1
	}

	midIdx := len(arr) / 2

	if arr[midIdx] == toFind {
		return midIdx
	} else if arr[midIdx] > toFind {
		return binarySearch(arr[:midIdx], toFind)
	} else {
		foundIdx := binarySearch(arr[midIdx+1:], toFind)
		if foundIdx == -1 {
			return -1
		} else {
			return foundIdx + (midIdx + 1)
		}
	}
}
