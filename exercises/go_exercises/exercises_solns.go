package goExercises

// func longestWord(sentence string) string {
// 	split := strings.Split(sentence, " ")
// 	longest := ""
// 	for _, word := range split {
// 		if len(word) > len(longest) {
// 			longest = word
// 		}
// 	}
// 	return longest
// }

// func longestPalindrome(sentence string) string {
// 	startIdx := 0
// 	longest := ""
// 	for startIdx < len(sentence)-1 {
// 		for endIdx := len(sentence) - 1; endIdx > startIdx; endIdx-- {
// 			subString := sentence[startIdx : endIdx+1]
// 			if subString == reverse(subString) && len(subString) > len(longest) {
// 				longest = subString
// 			}
// 		}
// 		startIdx++
// 	}
// 	return longest
// }

// func sumNums(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	return n + sumNums(n-1)
// }

// import (
// 	"strconv"
// 	"strings"
// )

// func mostCommonLetter(sentence string) []string {
// 	lettersMap := map[string]int{}
// 	for _, char := range strings.Split(sentence, "") {
// 		if char == " " {
// 			continue
// 		}
// 		if lettersMap[char] >= 1 {
// 			lettersMap[char]++
// 		} else {
// 			lettersMap[char] = 1
// 		}
// 	}
// 	count := 0
// 	mostCommon := []string{}
// 	for key, val := range lettersMap {
// 		if val > count {
// 			count = val
// 			mostCommon = []string{key, strconv.Itoa(val)}
// 		}
// 	}
// 	return mostCommon
// }

// import "fmt"

// //iterate forward O(n) and then backward O(n), combining the results in one array.
// func buyStockTwice(arr []int) int {
// 	maxProfit := 0
// 	maxProfits := make([]int, len(arr))
// 	low := 1000000
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] < low {
// 			low = arr[i]
// 		}
// 		if (arr[i] - low) > maxProfit {
// 			maxProfit = arr[i] - low
// 		}
// 		maxProfits[i] = maxProfit

// 	}
// 	fmt.Printf("Before backwards iteration: %v \n", maxProfits)
// 	totalMaxProfit := 0
// 	maxProfit = 0
// 	high := -1
// 	for j := len(arr) - 1; j > 0; j-- {
// 		if arr[j] > high {
// 			high = arr[j]
// 		}
// 		if (high - arr[j]) > maxProfit {
// 			maxProfit = high - arr[j]
// 		}
// 		maxProfits[j] = maxProfit + maxProfits[j-1]
// 		if maxProfits[j] > totalMaxProfit {
// 			totalMaxProfit = maxProfits[j]
// 		}
// 	}
// 	fmt.Printf("After backwards iteration: %v \n", maxProfits)
// 	return totalMaxProfit
// }
