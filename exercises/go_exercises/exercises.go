package goExercises

import "strings"

//return the nth prime
func nthPrime(nth int) int {
	primeCount := 0
	checkI := 2
	for primeCount < nth {
		if isPrime(checkI) {
			primeCount++
		}
		checkI++
	}
	return checkI -1
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}	
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}
 
//take an array of ints of length word, and return a string with each character of word in the order of idxArray
func scrambleString(word string, idxArray []int) string {
	scrambled := []string{}
	for i := 0; i < len(idxArray); i++ {
		strIdx := idxArray[i];
		scrambled = append(scrambled, string(word[strIdx]))
	}
	return strings.Join(scrambled, "")
}
 
