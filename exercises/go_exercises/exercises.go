package goExercises

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

//add dashes around an odd number, except on first or last el. Only allow one dash
//ex: 303 => 3-0-3
//ex: 333 => 3-3-3
//ex: 3223 => 3-22-3
func dasherizeNumber(n int) string {
	stringArr := []string{}
	for n > 0 {
		el := n % 10
		if el%2 == 0 {
			stringArr = append(stringArr, strconv.Itoa(el))
		} else {
			stringArr = append(stringArr, "-")
			stringArr = append(stringArr, strconv.Itoa(el))
			stringArr = append(stringArr, "-")
		}
		n = n / 10
	}

	//strip front and last
	if stringArr[0] == "-" {
		stringArr = stringArr[1:]
	}

	if stringArr[len(stringArr)-1] == "-" {
		stringArr = stringArr[:len(stringArr)-1]
	}

	stringed := strings.Join(stringArr, "")

	//remove double
	r, err := regexp.Compile("--")
	if err != nil {
		log.Fatal("Couldn't compile regexp")
	}

	stringed = r.ReplaceAllString(stringed, "-")
	return reverse(stringed)
}

// calculate the checksum of the multi-dimensional array
// For each row, determine the difference between the largest value and the smallest value;
// the checksum is the sum of all of these differences.
func checkSum(multiArr [][]int) int {
	checkSum := 0
	for _, row := range multiArr {
		colMax := 0
		colMin := 100000
		for _, col := range row {
			if col > colMax {
				colMax = col
			}

			if col < colMin {
				colMin = col
			}
		}
		checkSum += (colMax - colMin)
	}
	return checkSum
}

func checkSumAdvent() int {
	//read in the file ../practice/checksum_input and find the checksum for it
	return 0
}

//return num of vowels in sentence
func countVowels(sentence string) int {
	r, err := regexp.Compile("[aeiou]")
	if err != nil {
		log.Fatal("Failed to compile regexp")
	}
	chars := strings.Split(sentence, "")
	vowelCount := 0
	for _, char := range chars {
		if r.MatchString(char) {
			vowelCount++
		}
	}
	return vowelCount
}

//take sentence and reverse it, do not use library method
func reverse(sentence string) string {
	r := []rune(sentence)
	for i, j := 0, len(sentence)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
