package main

import (
	"strconv"
	"strings"
)

func reverse(sentence string) string {
	chars := strings.Split(sentence, "")
	idx1 := 0
	idx2 := len(sentence) - 1
	for idx1 < idx2 {
		chars[idx1], chars[idx2] = chars[idx2], chars[idx1]
		idx1++
		idx2--
	}
	return strings.Join(chars, "")
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func longestWord(sentence string) string {
	split := strings.Split(sentence, " ")
	longest := ""
	for _, word := range split {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

func sumNums(n int) int {
	if n == 1 {
		return 1
	}
	return n + sumNums(n-1)
}

func timeConversion(timeInMinutes int) string {
	minutes := timeInMinutes % 60
	hours := timeInMinutes / 60
	var time string
	if minutes < 10 {
		time = strconv.Itoa(hours) + ":0" + strconv.Itoa(minutes)
	} else {
		time = strconv.Itoa(hours) + ":" + strconv.Itoa(minutes)
	}
	return time
}

func countVowels(sentence string) int {
	split := strings.Split(sentence, "")
	vowels := "aeiou"
	count := 0
	for _, chars := range split {
		count += strings.Count(vowels, chars)
	}

	return count
}

func isPalindrome(word string) bool {
	idx1 := 0
	idx2 := len(word) - 1
	for idx1 < idx2 {
		if word[idx1] != word[idx2] {
			return false
		}
		idx1++
		idx2--
	}
	return true
}

func nearbyAZ(word string) bool {
	split := strings.Split(word, "")
	foundA := -2000
	for i, char := range split {
		if char == "a" {
			foundA = i
		}
		if char == "z" {
			if i-foundA <= 3 {
				return true
			}
		}
	}
	return false
}

func twoSum(arr []int) []int {
	idxs := []int{}
	for i := range arr {
		for j := range arr[i+1:] {
			if i+j == 0 {
				idxs = append(idxs, i, j)
			}
		}
	}
	return idxs
}
