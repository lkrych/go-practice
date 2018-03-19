package goExercises

import "fmt"

//return sorted arr using quicksort
func quickSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	pivot := arr[0]
	lessThanPivot := []int{}
	greaterThanPivot := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] <= pivot {
			lessThanPivot = append(lessThanPivot, arr[i])
		} else {
			greaterThanPivot = append(greaterThanPivot, arr[i])
		}
	}
	sortedLess := quickSort(lessThanPivot)
	sortedGreater := quickSort(greaterThanPivot)

	sortedLess = append(sortedLess, pivot)
	return append(sortedLess, sortedGreater...)
}

//In O(n) time, and using O(1) space, remove duplicates from an array
func removeDups(arr []int) []int {
	checkIdx := 0
	checkedAll := false
	for !checkedAll {

		if checkIdx == len(arr)-1 {
			checkedAll = true
			continue
		}
		if arr[checkIdx] == arr[checkIdx+1] {
			arr = append(arr[:checkIdx], arr[checkIdx+1:]...)
		} else {
			checkIdx++
		}
	}
	return arr
}

//take timeinminutes and convert it into a clock representation
//ex: 15 => 0:15
//ex: 150 => 2:30
//ex: 360 => 6:00
func timeConversion(timeInMinutes int) string {
	minutes := timeInMinutes % 60
	hours := timeInMinutes / 60

	return fmt.Sprintf("%d:%02d", hours, minutes)
}

//take in an array of digits encoding a decimal, d and update the array to rep d + 1
//ex: 1.29 => 1.30 ~ [1,2,9] => [1,3,0]
//ex 1.48372 => 1.48373
//TAKE AWAY:  This algorithm should work in a language that has finite-precision arithmetic
//Arrays can be used to break up numbers of arbitrary size, this allows us to handle very big numbers!
func incrementArb(arr []int) []int {
	last := arr[len(arr)-1]
	last++
	if last == 10 {
		newArr := incrementArb(arr[:len(arr)-1])
		newArr = append(newArr, 0)
		return newArr
	} else {
		arr[len(arr)-1] = last
		return arr
	}

}
