package goExercises

import "fmt"
 
//given an arr representing stock prices, return the maximum profit you could retrieve from buying and selling the stock
//ex: [310, 315, 275, 295, 260, 270, 290, 230, 255, 250] => 30
func buyStock(arr []int) int {
	low := 100000
	high := 0
	biggestProfit := 0

	for _, price := range arr {
		if price < low {
			//reset the current window
			low = price
			high = low
		}
		if price > high {
			high = price
		}

		if high - low > biggestProfit {
			biggestProfit = high - low
		}
	}
	fmt.Println("Biggest profit is", biggestProfit)
	return biggestProfit
}
 
//if the sum of digits in n is less than ten, return that sum, otherwise try to find the digital root of that sum
func digitalRoot(n int) int {
	digRoot := (n/10) + (n%10)
	if digRoot < 10 {
		return digRoot
	}
	return digitalRoot(digRoot)
}
 
