package goExercises

import (
	"math"
	"strconv"
	"strings"
)

//In O(n) time, and using O(1) space, remove duplicates from an array
func removeDups(arr []int) []int {
	nonRepeatIdx := 1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			arr[nonRepeatIdx] = arr[i+1]
			nonRepeatIdx++
		}
	}
	return arr[:nonRepeatIdx]
}

//convert base ten int n, into hexadecimal or binary representation
func baseConverter(n int, base int) string {
	stringNum := []string{}
	hexMap := map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
	}
	for n > 0 {
		if base == 2 {
			stringNum = append(stringNum, strconv.Itoa(n%2))
		} else {
			stringNum = append(stringNum, hexMap[n%16])
		}
		n = n / base
	}
	return reverse(strings.Join(stringNum, ""))
}

func reverse(s string) string {
	sArr := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		sArr[i], sArr[j] = sArr[j], sArr[i]
	}
	return string(sArr)
}

//TAKE AWAY:  This algorithm should work in a language that has finite-precision arithmetic
//Arrays can be used to break up numbers of arbitrary size, this allows us to handle very big numbers!
func multiplyArb(arr1 []int, arr2 []int) []int {
	sign := 1
	//check if numbers are negative
	if (arr1[0] * arr2[0]) > 0 {
		sign = 1
	} else {
		sign = -1
	}
	//set to absolute val
	arr1[0] = int(math.Abs(float64(arr1[0])))
	arr2[0] = int(math.Abs(float64(arr2[0])))

	total := make([]int, len(arr1)+len(arr2))
	for i := len(arr1) - 1; i >= 0; i-- {
		for j := len(arr2) - 1; j >= 0; j-- {
			total[i+j+1] += arr1[i] * arr2[j]
			total[i+j] += total[i+j+1] / 10
			total[i+j+1] %= 10
		}
	}
	zeroIdx := 0
	for {
		if total[zeroIdx] != 0 {
			break
		}
		zeroIdx++
	}
	total = total[zeroIdx:]
	total[0] = total[0] * sign

	return total
}

//return true if n is a power of two
func isPowerOfTwo(n float64) bool {
	pow := 0
	for float64(power(2, pow)) <= n {
		if float64(power(2, pow)) == n {
			return true
		}
		pow++
	}
	return false
}

func power(base int, exp int) int {
	if exp == 0 {
		return 1
	} else {
		return base * power(base, exp-1)
	}
}
