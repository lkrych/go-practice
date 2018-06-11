package goExercises

import (
	"regexp"
	"strconv"
	"strings"
)

//return index of toFind if it exists in arr, otherwise return -1
func binarySearch(arr []int, toFind int) int {
	if len(arr) == 0 {
		return -1
	}
	midIdx := len(arr) / 2
	if arr[midIdx] == toFind {
		return midIdx
	}
	if arr[midIdx] > toFind {
		return binarySearch(arr[:midIdx], toFind)
	} else {
		returned := binarySearch(arr[midIdx+1:], toFind)
		if returned == -1 {
			return returned
		}

		return 1 + midIdx + returned
	}

}

//add dashes around an odd number, except on first or last el. Only allow one dash
//ex: 303 => 3-0-3
//ex: 333 => 3-3-3
//ex: 3223 => 3-22-3
func dasherizeNumber(n int) string {
	strArr := []string{}
	for n > 0 {
		digit := n % 10
		if digit%2 == 0 {
			strArr = append(strArr, strconv.Itoa(digit))
		} else {
			strArr = append(strArr, "-")
			strArr = append(strArr, strconv.Itoa(digit))
			strArr = append(strArr, "-")
		}
		n = n / 10
	}

	//replace first and last
	if strArr[0] == "-" {
		strArr = strArr[1:]
	}

	if strArr[len(strArr)-1] == "-" {
		strArr = strArr[:len(strArr)-1]
	}
	//replace double dashes
	r := regexp.MustCompile("--")
	str := strings.Join(strArr, "")
	str = r.ReplaceAllString(str, "-")

	return reverse(str)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//return num of vowels in sentence
func countVowels(sentence string) int {
	vowelCount := 0
	r := regexp.MustCompile("[aeiou]")
	for _, char := range strings.Split(sentence, "") {
		if r.MatchString(char) {
			vowelCount++
		}
	}
	return vowelCount
}

//Given a string, find the length of the longest substring without repeating characters.
func longestSubString(sentence string) int {
	bs := []byte(sentence)
	var maxLen, start int
	idxMap := map[byte]int{}
	for i := 0; i < len(bs); i++ {
		if _, ok := idxMap[bs[i]]; ok && start <= idxMap[bs[i]] {
			start = idxMap[bs[i]] + 1
		} else {
			if maxLen < i-start+1 {
				maxLen = i - start + 1
			}
		}
		idxMap[bs[i]] = i
	}
	return maxLen

}
