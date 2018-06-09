package goExercises

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

//return a string where every char in sentence is shifted by shift
func caesarCipher(shift int, sentence string) string {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	cipher := make([]byte, len(sentence))
	for i, el := range strings.Split(sentence, "") {
		currIdx := strings.Index(alpha, el)
		if currIdx == -1 {
			cipher[i] = ' '
			continue
		}
		cipher[i] = alpha[(currIdx+shift)%26]
	}
	return string(cipher)
}

//Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
func isUnique(word string) bool {
	alphaMap := make(map[string]bool)
	for _, el := range strings.Split(word, "") {
		if _, ok := alphaMap[el]; ok {
			return false
		} else {
			alphaMap[el] = true
		}
	}
	return true
}

//return the nth prime
func nthPrime(nth int) int {
	n := 0
	checkIfPrime := 2
	for n < nth {
		if isPrime(checkIfPrime) {
			n++
		}
		checkIfPrime++
	}
	return checkIfPrime - 1
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//add dashes around an odd number, except on first or last el. Only allow one dash
//ex: 303 => 3-0-3
//ex: 333 => 3-3-3
//ex: 3223 => 3-22-3
func dasherizeNumber(n int) string {
	intString := []string{}
	for n > 0 {
		digit := n % 10
		if digit%2 == 0 {
			intString = append(intString, strconv.Itoa(digit))
		} else {
			intString = append(intString, "-")
			intString = append(intString, strconv.Itoa(digit))
			intString = append(intString, "-")
		}
		n = n / 10
	}

	//remove first dash and last dash
	if intString[0] == "-" {
		intString = intString[1:]
	}

	if intString[len(intString)-1] == "-" {
		intString = intString[:len(intString)-1]
	}

	//remove double dashes
	r, err := regexp.Compile("--")
	if err != nil {
		log.Fatal("Regexp didn't compile")
	}
	stringed := strings.Join(intString, "")

	stringed = r.ReplaceAllString(stringed, "-")
	return reverse(stringed)
}

func reverse(sentence string) string {
	r := []rune(sentence)
	for i, j := 0, len(sentence)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
