package goExercises

//return sum of n + n-1 + ... 1
func sumNums(n int) int {
	if n <= 1 {
		return n
	} else {
		return n + sumNums(n-1)
	}
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

//return the nth prime
func nthPrime(nth int) int {
	primeCount := 0
	checkifPrime := 2
	for primeCount < nth {
		if isPrime(checkifPrime) {
			primeCount++
		}
		checkifPrime++
	}
	return checkifPrime - 1
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//take an array of ints of length word, and return a string with each character of word in the order of idxArray
func scrambleString(word string, idxArray []int) string {
	scrambled := make([]rune, len(word))
	runes := []rune(word)
	for i, el := range idxArray {
		scrambled[i] = runes[el]
	}
	return string(scrambled)
}
