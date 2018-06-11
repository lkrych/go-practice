package goExercises
 
// func binarySearch(arr []int, toFind int) int {
// 	if len(arr) < 1 {
// 		return -1
// 	}
// 	midIdx := len(arr) / 2
// 	if toFind == arr[midIdx] {
// 		return midIdx
// 	} else if toFind > arr[midIdx] {
// 		search := binarySearch(arr[midIdx+1:], toFind)
// 		if search != -1 {
// 			return search + midIdx + 1
// 		}
// 		return search
// 	} else {
// 		return binarySearch(arr[0:midIdx], toFind)
// 	}

// }
 
// import (
// 	"strconv"
// 	"strings"
// )

// func dasherizeNumber(n int) string {
// 	dasherized := []string{}
// 	toString := strconv.Itoa(n)
// 	for i, el := range strings.Split(toString, "") {

// 		if int, _ := strconv.Atoi(el); int%2 != 0 {
// 			if i == 0 {
// 				dasherized = append(dasherized, el, "-")
// 			} else if i == len(toString)-1 {
// 				dasherized = append(dasherized, "-", el)
// 			} else {
// 				dasherized = append(dasherized, "-", el, "-")
// 			}
// 		} else {
// 			dasherized = append(dasherized, el)
// 		}
// 	}
// 	joined := strings.Join(dasherized, "")
// 	return strings.Replace(joined, "--", "-", -1)
// }
 
// import "strings"

// func countVowels(sentence string) int {
// 	split := strings.Split(sentence, "")
// 	vowels := "aeiou"
// 	count := 0
// 	for _, chars := range split {
// 		count += strings.Count(vowels, chars)
// 	}

// 	return count
// }
 
// //Given a string, find the length of the longest substring without repeating characters.
// func longestSubString(sentence string) int {

// }
 
