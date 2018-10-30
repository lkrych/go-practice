package goExercises

import "fmt"

//given an arr representing stock prices, return the maximum profit you could retrieve from buying and selling the stock twice
//ex: [310, 315, 275, 295, 260, 270, 290, 230, 255, 250] => 50
func buyStockTwice(arr []int) int {
	//hold the current state of the difference between days if larger than previous
	maxProfit := 0
	//make storage array for keeping track fo maxProfits, most useful for step 2
	//O(n) space
	maxProfits := make([]int, len(arr))
	//loop over arr O(n)
	currLow := 1000000 // big number
	for i := 0; i < len(arr); i++ {
		if arr[i] < currLow {
			currLow = arr[i]
		}

		currentDifference := arr[i] - currLow
		if currentDifference > maxProfit {
			maxProfit = currentDifference
		}

		maxProfits[i] = maxProfit
	}
	// we now have an array maxProfits that represents the current maxProfit per day
	fmt.Println("Max profits:", maxProfits)
	// in step 2 we do the exact same procedure in reverse.
	currHigh := -1
	maxProfit = 0
	//store the addition of the current profit with the past profit computed in step 1 in this variable
	totalProfit := 0
	//step through the array going backwards O(n)
	for j := len(arr) - 1; j > 0; j-- {
		if arr[j] > currHigh {
			currHigh = arr[j]
		}

		currentDifference := currHigh - arr[j]
		if currentDifference > maxProfit {
			maxProfit = currentDifference
		}

		maxProfits[j] = maxProfit + maxProfits[j-1]
		if maxProfits[j] > totalProfit {
			totalProfit = maxProfits[j]
		}
	}

	return totalProfit
}

//return sum of n + n-1 + ... 1
func sumNums(n int) int {
	if n == 1 {
		return 1
	} else {
		return sumNums(n-1) + n
	}
}
