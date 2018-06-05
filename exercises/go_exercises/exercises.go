package goExercises

import (
	"regexp"
	"strconv"
	"strings"
)

//add dashes around an odd number, except on first or last el. Only allow one dash
//ex: 303 => 3-0-3
//ex: 333 => 3-3-3
//ex: 3223 => 3-22-3
func dasherizeNumber(n int) string {
	buildString := []string{}
	for n > 0 {
		val := n % 10
		if (val)%2 == 0 {
			buildString = append(buildString, strconv.Itoa(val))
		} else {
			buildString = append(buildString, "-")
			buildString = append(buildString, strconv.Itoa(val))
			buildString = append(buildString, "-")
		}
		n = n / 10
	}

	//remove from front and back
	if buildString[0] == "-" {
		buildString = buildString[1:]
	}

	if buildString[len(buildString)-1] == "-" {
		buildString = buildString[:len(buildString)-1]
	}

	//remove double dashes
	r, _ := regexp.Compile("--")

	joined := Reverse(strings.Join(buildString, ""))

	return r.ReplaceAllString(joined, "-")

}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//return num of vowels in sentence
func countVowels(sentence string) int {
	count := 0
	r, _ := regexp.Compile("[aeiou]")
	for _, el := range strings.Split(sentence, "") {
		if ok := r.MatchString(el); ok {
			count++
		}
	}
	return count
}

//return n!
func factorial(n int) int {
	if n < 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
