package goExercises

import (
	"fmt"
	"strconv"
	"strings"
)

//take in an array of digits encoding a decimal, d and update the array to rep d + 1
//ex: 1.29 => 1.30 ~ [1,2,9] => [1,3,0]
//ex 1.48372 => 1.48373
//TAKE AWAY:  This algorithm should work in a language that has finite-precision arithmetic
//Arrays can be used to break up numbers of arbitrary size, this allows us to handle very big numbers!
func incrementArb(arr []int) []int {
	carry := 1
	idx := len(arr) - 1
	for carry > 0 {
		sum := carry + arr[idx]
		arr[idx] = sum % 10
		carry = sum / 10
		idx--
	}
	return arr
}

//return n!
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

//take timeinminutes and convert it into a clock representation
//ex: 15 => 0:15
//ex: 150 => 2:30
//ex: 360 => 6:00
func timeConversion(timeInMinutes int) string {
	mins := timeInMinutes % 60
	hrs := timeInMinutes / 60
	return fmt.Sprintf("%v:%02d", hrs, mins)
}

//implement a method to perform basic string compression using the counts of repeated chars
//ex: aabcccccaaa = a2b1c5a3
func stringCompression(in string) string {
	if checkStringCompresionLength(in) > len(in) {
		return in
	}
	compressed := []string{}
	consecutive := 0
	for i := 0; i < len(in); i++ {
		consecutive++
		char := in[i]
		if i+1 >= len(in) || char != in[i+1] {
			compressed = append(compressed, string(char), strconv.Itoa(consecutive))
			consecutive = 0
		}
	}

	return strings.Join(compressed, "")
}

func checkStringCompresionLength(in string) int {
	compressedCount := 0
	consecutive := 0
	for i := 0; i < len(in)-1; i++ {
		consecutive++
		char := in[i]
		if char != in[i+1] {
			compressedCount = compressedCount + (1 + len(strconv.Itoa(consecutive)))
			consecutive = 0
		}
	}
	return compressedCount
}
