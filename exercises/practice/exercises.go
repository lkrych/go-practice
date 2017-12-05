package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
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

func twoSum(arr []int, find int) []int {
	idxs := []int{}
	for i, val1 := range arr {
		for j, val2 := range arr[i+1:] {
			if val1+val2 == find {
				//bc you are iterating with modified slice
				//you need to take into account correct j index
				idxs = append(idxs, i, j+i+1)
			}
		}
	}
	return idxs
}

func twoSumOptimized(arr []int, find int) []int {
	sums := []int{}
	idxs := map[int]int{}
	//fill map
	for i, val := range arr {
		idxs[val] = i
	}
	//search for correct factors to negate the key
	for key, val := range idxs {
		search := find - key
		if searchIdx, exists := idxs[search]; exists {
			sums = append(sums, val, searchIdx)
			return sums
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

func longestPalindrome(sentence string) string {
	startIdx := 0
	longest := ""
	for startIdx < len(sentence)-1 {
		for endIdx := len(sentence) - 1; endIdx > startIdx; endIdx-- {
			subString := sentence[startIdx : endIdx+1]
			if subString == reverse(subString) && len(subString) > len(longest) {
				longest = subString
			}
		}
		startIdx++
	}
	return longest
}

func greatestCommonFactor(n1 int, n2 int) int {
	var smaller int
	var larger int
	if n1 > n2 {
		smaller = n2
		larger = n1
	} else {
		smaller = n1
		larger = n2
	}

	for i := 1; i <= smaller; i++ {
		if smaller%i == 0 {
			factor := smaller / i
			if larger%factor == 0 {
				return factor
			}
		}
	}
	//if no factors are found, the GCF is 1
	return 1
}

func caesarCipher(shift int, sentence string) string {
	shifted := []string{}
	for _, ascii := range []byte(sentence) {
		if ascii == 32 {
			shifted = append(shifted, " ")
		} else {
			shifted = append(shifted, shiftChar(shift, ascii))
		}
	}
	return strings.Join(shifted, "")
}

func shiftChar(shift int, ascii byte) string {
	newChar := int(ascii) + (shift % 26)
	if newChar > 122 {
		newChar = 96 + (newChar % 122)
	}
	return string(newChar)
}

func numRepeats(word string) int {
	lettersMap := map[string]int{}
	for _, ch := range strings.Split(word, "") {
		if ch == " " {
			continue
		}
		if lettersMap[ch] >= 1 {
			lettersMap[ch]++
		} else {
			lettersMap[ch] = 1
		}
	}
	count := 0
	for _, v := range lettersMap {
		if v > 1 {
			count++
		}
	}
	return count
}

func baseConverter(n int, base int) string {
	base16 := map[int]string{
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f"}
	answer := []string{}
	for n > 0 {
		answer = append(answer, strconv.Itoa(n%base))
		n = n / base
	}
	if base == 16 {
		for i, el := range answer {
			intEl, err := strconv.Atoi(el)
			if err != nil {
				fmt.Println("There was an error!")
			}
			if hex, exists := base16[intEl]; exists {
				answer[i] = hex
			}
		}
	}
	return reverse(strings.Join(answer, ""))
}

func binarySearch(arr []int, toFind int) int {
	if len(arr) < 1 {
		return -1
	}
	midIdx := len(arr) / 2
	if toFind == arr[midIdx] {
		return midIdx
	} else if toFind > arr[midIdx] {
		search := binarySearch(arr[midIdx+1:], toFind)
		if search != -1 {
			return search + midIdx + 1
		}
		return search
	} else {
		return binarySearch(arr[0:midIdx], toFind)
	}

}

func bubbleSort(arr []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}
	return arr
}

func digitalRoot(n int) int {
	if n < 10 {
		return n
	}
	return digitalRoot(n%10 + n/10)
}

func mergeSort(arr []int) []int {
	//make copy of array so you don't mutate input
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	if len(tmp) <= 1 {
		return tmp
	}
	mid := len(tmp) / 2
	left := mergeSort(tmp[0:mid])
	right := mergeSort(tmp[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	sorted := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			sorted = append(sorted, left[0])
			left = left[1:]
		} else {
			sorted = append(sorted, right[0])
			right = right[1:]
		}
	}
	sorted = append(sorted, left...)
	sorted = append(sorted, right...)
	return sorted
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	left := []int{}
	right := []int{}

	for _, val := range arr[1:] {
		if val >= pivot {
			right = append(right, val)
		} else {
			left = append(left, val)
		}
	}
	answer := []int{}
	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)
	answer = append(answer, sortedLeft...)
	answer = append(answer, pivot)
	answer = append(answer, sortedRight...)
	return answer
}

func exponent(base int, exp int) int {
	if exp == 1 {
		return base
	}
	if exp < 0 {
		return 1 / (base * exponent(base, (exp*-1)-1))
	}
	return base * exponent(base, exp-1)
}

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func partitionEvenOdd(arr []int) []int {
	partition := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			arr[i], arr[partition] = arr[partition], arr[i]
			partition++
		}
	}
	front := quickSort(arr[:partition])
	back := quickSort(arr[partition:])
	return append(front, back...)
}

func incrementArb(arr []int) []int {
	lastIdx := len(arr) - 1
	arr[lastIdx] = (arr[lastIdx] + 1) % 10
	if arr[lastIdx] == 0 {
		arr = append(incrementArb(arr[:lastIdx]), arr[lastIdx])
	}
	return arr
}

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

func advancingArray(arr []int) bool {
	farthestTraveled := arr[0]
	for i := 0; i <= farthestTraveled; i++ {
		if farthestTraveled == len(arr)-1 {
			return true
		}
		farthestTraveled = int(math.Max(float64(farthestTraveled), float64(arr[i]+i)))
	}
	return false
}

func removeDups(arr []int) []int {
	writeIdx := 1
	for i := 1; i < len(arr); i++ {
		if arr[writeIdx-1] != arr[i] {
			arr[writeIdx] = arr[i]
			writeIdx++
		}
	}
	return arr[:writeIdx]
}

func buyStock(arr []int) int {
	low := 1000000
	diff := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < low {
			low = arr[i]
		}
		if (arr[i] - low) > diff {
			diff = arr[i] - low
		}
	}
	return diff
}

//iterate forward O(n) and then backward O(n), combining the results in one array.
func buyStockTwice(arr []int) int {
	maxProfit := 0
	maxProfits := make([]int, len(arr))
	low := 1000000
	for i := 0; i < len(arr); i++ {
		if arr[i] < low {
			low = arr[i]
		}
		if (arr[i] - low) > maxProfit {
			maxProfit = arr[i] - low
		}
		maxProfits[i] = maxProfit

	}
	fmt.Printf("Before backwards iteration: %v \n", maxProfits)
	totalMaxProfit := 0
	maxProfit = 0
	high := -1
	for j := len(arr) - 1; j > 0; j-- {
		if arr[j] > high {
			high = arr[j]
		}
		if (high - arr[j]) > maxProfit {
			maxProfit = high - arr[j]
		}
		maxProfits[j] = maxProfit + maxProfits[j-1]
		if maxProfits[j] > totalMaxProfit {
			totalMaxProfit = maxProfits[j]
		}
	}
	fmt.Printf("After backwards iteration: %v \n", maxProfits)
	return totalMaxProfit
}

func reverseCaptcha(arr []int) int {

	sum := 0
	for i := 1; i < len(arr); i++ {
		if arr[i-1] == arr[i] {
			sum += arr[i]
			fmt.Printf("Repeat: %v \n", arr[i])
		}
	}
	fmt.Printf("Arr[0]: %v \n", arr[0])
	fmt.Printf("Arr[-1]: %v \n", arr[len(arr)-1])
	if arr[0] == arr[len(arr)-1] {
		sum += arr[0]
		fmt.Printf("Repeat: %v \n", arr[0])
	}
	return sum
}

func adventCaptcha() int {
	content, err := ioutil.ReadFile("reverse_captcha_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	strs := strings.Split(string(content), "")
	ints := make([]int, len(strs))
	for i := 0; i < len(ints); i++ {
		ints[i], _ = strconv.Atoi(strs[i])
	}
	ints = ints[:len(ints)-1]
	sum := halfwayReverseCaptcha(ints)
	return sum
}

func halfwayReverseCaptcha(arr []int) int {
	sum := 0
	halfway := len(arr) / 2
	for i := 1; i < halfway; i++ {
		if arr[i] == arr[i+halfway] {
			sum += arr[i] * 2
		}
	}

	return sum
}

func checkSum(multiArr [][]int) int {
	checksum := 0
	for idx, arr := range multiArr {
		max := -100000
		min := 100000
		for i := 0; i < len(multiArr[idx]); i++ {
			if arr[i] > max {
				max = arr[i]
			}
			if arr[i] < min {
				min = arr[i]
			}
		}
		checksum += (max - min)
	}
	return checksum
}

func checkSumAdvent() [][]int {
	content, err := ioutil.ReadFile("checksum_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	arrays := strings.Split(string(content), "\n")
	multiarr := make([][]int, len(arrays)-1)

	for i := 0; i < len(arrays)-1; i++ {
		split := strings.Split(string(arrays[i]), "	")
		intSplit := make([]int, len(split))
		for j := 0; j < len(split); j++ {
			intSplit[j], _ = strconv.Atoi(split[j])
		}
		multiarr[i] = intSplit
	}
	return multiarr
}

func checkSumEven(multiArr [][]int) int {
	checksum := 0
	for i := range multiArr {
		for j := 0; j < len(multiArr[i]); j++ {
			for k := j + 1; k < len(multiArr[i]); k++ {
				elJ := multiArr[i][j]
				elK := multiArr[i][k]

				if elJ%elK == 0 {
					checksum += elJ / elK
					continue
				}

				if elK%elJ == 0 {
					checksum += elK / elJ
					continue
				}

			}
		}
	}
	return checksum
}

func spiralMemory(n int) int {
	sum := 1
	idx := 1
	for sum < n {
		sum += 8 * idx
		idx++
	}
	perRow := (idx - 1) * 2
	fmt.Printf("Per Row: %v \n", perRow)

	for sum > n {
		sum -= perRow
	}
	distanceToN := 0

	for sum < n {
		sum++
		distanceToN++
	}
	fmt.Printf("Distance to n: %v \n", distanceToN)
	middle := perRow / 2
	if distanceToN > middle {
		return idx + (distanceToN - middle)
	}

	return idx + (middle - distanceToN) - 1

}

func passPhrase(passphrase string) bool {
	split := strings.Split(passphrase, " ")
	wordMap := make(map[string]int)
	for _, word := range split {
		if _, ok := wordMap[word]; ok {
			fmt.Printf("%v is the repeated word in %v \n", word, passphrase)
			return false
		}
		wordMap[word] = 1
	}
	return true
}

func passPhraseAdvent() int {
	phraseCount := 0
	content, err := ioutil.ReadFile("passphrase_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	splitByLine := strings.Split(string(content), "\n")
	for _, line := range splitByLine {
		if passPhrase(line) {
			phraseCount++
		}
	}
	return phraseCount
}

func passPhraseAnagram(passphrase string) bool {
	split := strings.Split(passphrase, " ")
	wordMap := make(map[string]int)
	for _, word := range split {
		sorted := SortString(word)
		if _, ok := wordMap[sorted]; ok {
			return false
		}
		wordMap[sorted] = 1
	}
	return true
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func passPhraseAnagramAdvent() int {
	phraseCount := 0
	content, err := ioutil.ReadFile("passphrase_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	splitByLine := strings.Split(string(content), "\n")
	for _, line := range splitByLine {
		if passPhraseAnagram(line) {
			phraseCount++
		}
	}
	return phraseCount
}
