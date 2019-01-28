package goExercises

import (
	"fmt"
	"strings"
)

//return num of vowels in sentence
func countVowels(sentence string) int {
	vowels := []string{"a", "e", "i", "o", "u"}
	vowelCount := 0
	split := strings.Split(sentence, "")
	for _, ch := range split {
		for _, vowel := range vowels {
			if ch == vowel {
				vowelCount++
			}
		}
	}
	return vowelCount
}

//given an arr representing stock prices, return the maximum profit you could retrieve from buying and selling the stock
//ex: [310, 315, 275, 295, 260, 270, 290, 230, 255, 250] => 30
func buyStock(arr []int) int {
	low := 1000000 //high number
	high := 0 //low number
	profit := 0
	fmt.Printf("Arr is %v \n", arr) 
	for _, price := range arr {
		if price < low {
			low = price
			high = price
		}
		if price > high {
			high = price
		}
		if high - low > profit {
			profit = high - low
		}
		fmt.Printf("High is %v, Low is %v, Profit is %v \n", high, low, profit) 
	}
	return profit
}

