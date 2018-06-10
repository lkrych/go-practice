package goExercises

import (
	"fmt"
	"math"
)

//take in an array of digits encoding a decimal, d and update the array to rep d + 1
//ex: 1.29 => 1.30 ~ [1,2,9] => [1,3,0]
//ex 1.48372 => 1.48373
//TAKE AWAY:  This algorithm should work in a language that has finite-precision arithmetic
//Arrays can be used to break up numbers of arbitrary size, this allows us to handle very big numbers!
func incrementArb(arr []int) []int {
	lastIdx := len(arr) - 1
	carry := (1 + arr[lastIdx]) / 10
	arr[lastIdx] = (1 + arr[lastIdx]) % 10
forLoop:
	for i := lastIdx - 1; i > 0; i-- {
		if carry > 0 {
			sum := carry + arr[i]
			carry = sum / 10
			arr[i] = sum % 10
		} else {
			break forLoop
		}
	}
	return arr
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

//return sorted arr using mergesort
func mergeSort(arr []int) []int {
	if len(arr) <= 1 { //arr is sorted!
		return arr
	}
	midIdx := len(arr) / 2
	leftMerge := mergeSort(arr[:midIdx])
	rightMerge := mergeSort(arr[midIdx:])
	return merge(leftMerge, rightMerge)
}

func merge(left, right []int) []int {
	merged := []int{}
	for len(left) > 0 && len(right) > 0 {

		if left[0] <= right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	if len(left) == 0 {
		merged = append(merged, right...)
		right = []int{}
	} else if len(right) == 0 {
		merged = append(merged, left...)
		left = []int{}
	}
	return merged
}

//return true if n is a power of two
func isPowerOfTwo(n float64) bool {
	exp := 0
	for pow(2, exp) <= n {
		if pow(2, exp) == n {
			return true
		}
		exp++
	}
	return false
}

func pow(base, exp int) float64 {
	if exp == 0 {
		return 1
	} else if exp == 1 {
		return float64(base)
	} else {
		return float64(base) * pow(base, exp-1)
	}
}

//think of this exercise like a variation of mancala.
//Starting at the zeroth index, can you navigate to the last index using the number of hops given to you at each index?
//You may use less than the number of hops given to you at each index
//return true if you can, return false if it is impossible
//ex [3,3,1,0,2,0,1] is true
// hop once from 3 to 3, hop three times to two and hop twice to 1!
//ex [3,2,0,0,2,0,1] is false
// the farthest you can hop is to index 4
func advancingArray(arr []int) bool {
	farthestTraveled := arr[0]
	for i := 0; i <= farthestTraveled; i++ {
		fmt.Println("the board is", arr)
		fmt.Println("Checking idx", i)
		if farthestTraveled == len(arr)-1 {
			fmt.Println("maneuvered to the end")
			return true
		}
		fmt.Println("Comparing farthestTraveled", farthestTraveled, "to ", arr[i]+i)

		farthestTraveled = int(math.Max(float64(farthestTraveled), float64(arr[i]+i)))
	}
	fmt.Println("didn't find the end")
	return false
}
