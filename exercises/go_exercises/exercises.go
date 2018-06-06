package goExercises

import (
	"sort"
	"strings"
)

//take in an array of digits encoding a decimal, d and update the array to rep d + 1
//ex: 1.29 => 1.30 ~ [1,2,9] => [1,3,0]
//ex 1.48372 => 1.48373
//TAKE AWAY:  This algorithm should work in a language that has finite-precision arithmetic
//Arrays can be used to break up numbers of arbitrary size, this allows us to handle very big numbers!
func incrementArb(arr []int) []int {
	carry := 0
forLoop:
	for j := len(arr) - 1; j > 0; j-- {
		sum := arr[j] + 1 + carry
		carry = sum % 10
		if sum/10 > 0 {
			arr[j] = sum % 10
			continue
		} else {
			arr[j] = sum
			break forLoop
		}
	}
	return arr
}

//given an arr representing stock prices, return the maximum profit you could retrieve from buying and selling the stock twice
//ex: [310, 315, 275, 295, 260, 270, 290, 230, 255, 250] => 50
func buyStockTwice(arr []int) int {
	maxProfits := make([]int, len(arr))
	low := 1000000
	diff := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < low {
			low = arr[i]
		}
		if (arr[i] - low) > diff {
			diff = arr[i] - low
		}
		maxProfits[i] = diff
	}
	totalMaxProfit := 0
	high := -1
	maxProfit := 0
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

//given an arr representing stock prices, return the maximum profit you could retrieve from buying and selling the stock
//ex: [310, 315, 275, 295, 260, 270, 290, 230, 255, 250] => 30
func buyStock(arr []int) int {
	low := 1000000
	diff := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < low {
			low = arr[i]
		}
		if (arr[i] - low) > diff {
			diff = arr[i] - low
		}
	}
	return diff
}

//Given two strings, write a method to decide if one is a permutation of the other
func checkPermutation(s1 string, s2 string) bool {
	s1Array := strings.Split(s1, "")
	s2Array := strings.Split(s2, "")
	sort.Strings(s1Array)
	sort.Strings(s2Array)
	if strings.Join(s1Array, "") == strings.Join(s2Array, "") {
		return true
	} else {
		return false
	}
}
