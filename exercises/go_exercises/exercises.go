package goExercises

import (
	"strconv"
	"strings"
)

//return the longest word in sentence
func longestWord(sentence string) string {
	longestWord := ""
	for _, word := range strings.Split(sentence, " ") {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}
	return longestWord
}

//return the longest palindrome sequence in sentence
func longestPalindrome(sentence string) string {
	longestPal := ""
	for i := 0; i < len(sentence)-1; i++ {
		for j := len(sentence) - 1; j > i; j-- {
			word := sentence[i : j+1]
			if isPalindrome(word) {
				if len(word) > len(longestPal) {
					longestPal = word
				}
			}

		}
	}
	return longestPal
}

func isPalindrome(word string) bool {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}

//return sum of n + n-1 + ... 1
func sumNums(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + sumNums(n-1)
	}
}

//return the most common character in a string with its count
func mostCommonLetter(sentence string) []string {
	charMap := map[string]int{}
	maxCount := 0
	maxChar := ""
	for _, el := range strings.Split(sentence, "") {
		if el == " " {
			continue
		}

		if count, ok := charMap[el]; ok {
			if count+1 > maxCount {
				maxCount = count + 1
				maxChar = el
			}
			charMap[el] = count + 1
		} else {
			charMap[el] = 1
		}
	}

	return []string{maxChar, strconv.Itoa(maxCount)}
}

//given an arr representing stock prices, return the maximum profit you could retrieve from buying and selling the stock twice
//ex: [310, 315, 275, 295, 260, 270, 290, 230, 255, 250] => 50
func buyStockTwice(arr []int) int {
	maxProfits := make([]int, len(arr))
	maxProfit := 0
	minBuy := 1000000
	for i := 0; i < len(arr); i++ {
		if arr[i] < minBuy {
			minBuy = arr[i]
		}
		if (arr[i] - minBuy) > maxProfit {
			maxProfit = arr[i] - minBuy
		}
		maxProfits[i] = maxProfit
	}
	high := -1
	maxProfit = 0
	totalMaxProfit := 0
	for j := len(arr) - 1; j > 0; j-- {
		if arr[j] > high {
			high = arr[j]
		}
		if (high - arr[j]) > maxProfit {
			maxProfit = high - arr[j]
		}
		maxProfits[j] = maxProfit + maxProfits[j-1]
		if maxProfits[j] > totalMaxProfit {
			totalMaxProfit = maxProfits[j]
		}

	}
	return totalMaxProfit

}
