package main

import (
	"math"
	"strconv"
	"strings"
)

func reverse(sentence string) string {
	chars := strings.Split(sentence, "")
	idx1 := 0
	idx2 := len(sentence) - 1
	for idx1 < idx2 {
		chars[idx1], chars[idx2] = chars[idx2], chars[idx1]
		idx1++
		idx2--
	}
	return strings.Join(chars, "")
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func longestWord(sentence string) string {
	split := strings.Split(sentence, " ")
	longest := ""
	for _, word := range split {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

func sumNums(n int) int {
	if n == 1 {
		return 1
	}
	return n + sumNums(n-1)
}

func timeConversion(timeInMinutes int) string {
	minutes := timeInMinutes % 60
	hours := timeInMinutes / 60
	var time string
	if minutes < 10 {
		time = strconv.Itoa(hours) + ":0" + strconv.Itoa(minutes)
	} else {
		time = strconv.Itoa(hours) + ":" + strconv.Itoa(minutes)
	}
	return time
}

func countVowels(sentence string) int {
	split := strings.Split(sentence, "")
	vowels := "aeiou"
	count := 0
	for _, chars := range split {
		count += strings.Count(vowels, chars)
	}

	return count
}

func isPalindrome(word string) bool {
	idx1 := 0
	idx2 := len(word) - 1
	for idx1 < idx2 {
		if word[idx1] != word[idx2] {
			return false
		}
		idx1++
		idx2--
	}
	return true
}

func nearbyAZ(word string) bool {
	split := strings.Split(word, "")
	foundA := -2000
	for i, char := range split {
		if char == "a" {
			foundA = i
		}
		if char == "z" {
			if i-foundA <= 3 {
				return true
			}
		}
	}
	return false
}

func twoSum(arr []int) []int {
	idxs := []int{}
	for i, val1 := range arr {
		for j, val2 := range arr[i+1:] {
			if val1+val2 == 0 {
				//bc you are iterating with modified slice
				//you need to take into account correct j index
				idxs = append(idxs, i, j+i+1)
			}
		}
	}
	return idxs
}

func twoSumOptimized(arr []int) []int {
	sums := []int{}
	idxs := map[int]int{}
	//fill map
	for i, val := range arr {
		idxs[val] = i
	}
	//search for correct factors to negate the key
	//add conditional to make sure values haven't been found
	found := map[int]bool{}
	for key, val := range idxs {
		search := 0 - key
		if searchIdx, exists := idxs[search]; exists && !found[searchIdx] {
			sums = append(sums, val, searchIdx)
			found[val] = true
			found[searchIdx] = true
		}
	}
	return sums
}

func isPowerOfTwo(n float64) bool {
	power := float64(0)
	value := math.Exp2(power)
	for value <= n {
		if value == n {
			return true
		}
		power++
		value = math.Exp2(power)
	}
	return false
}

func thirdGreatest(arr []int) int {
	var first int
	var second int
	var third int

	for _, el := range arr {
		if el >= first {
			third = second
			second = first
			first = el
			continue
		}
		if el >= second {
			third = second
			second = el
			continue
		}
		if el >= third {
			third = el
		}
	}
	return third
}

func dasherizeNumber(n int) string {
	dasherized := []string{}
	toString := strconv.Itoa(n)
	for i, el := range strings.Split(toString, "") {

		if int, _ := strconv.Atoi(el); int%2 != 0 {
			if i == 0 {
				dasherized = append(dasherized, el, "-")
			} else if i == len(toString)-1 {
				dasherized = append(dasherized, "-", el)
			} else {
				dasherized = append(dasherized, "-", el, "-")
			}
		} else {
			dasherized = append(dasherized, el)
		}
	}
	joined := strings.Join(dasherized, "")
	return strings.Replace(joined, "--", "-", -1)
}

func mostCommonLetter(sentence string) []string {
	lettersMap := map[string]int{}
	for _, char := range strings.Split(sentence, "") {
		if char == " " {
			continue
		}
		if lettersMap[char] >= 1 {
			lettersMap[char]++
		} else {
			lettersMap[char] = 1
		}
	}
	count := 0
	mostCommon := []string{}
	for key, val := range lettersMap {
		if val > count {
			count = val
			mostCommon = []string{key, strconv.Itoa(val)}
		}
	}
	return mostCommon
}

func capitalizeWords(sentence string) string {
	return strings.Title(sentence)
}

func scrambleString(word string, idxArray []int) string {
	split := strings.Split(word, "")
	scrambled := []string{}
	for _, el := range idxArray {
		scrambled = append(scrambled, split[el])
	}
	return strings.Join(scrambled, "")
}

func isPrime(n int) bool {
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nthPrime(nth int) int {
	primeCount := 0
	n := 2
	for primeCount < nth {
		if isPrime(n) {
			primeCount++
		}
		n++
	}
	return n - 1
}
