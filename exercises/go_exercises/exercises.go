package goExercises

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//implement a method to perform basic string compression using the counts of repeated chars
//ex: aabcccccaaa = a2b1c5a3
func stringCompression(in string) string {
	if compressedLength(in) > len(in) { // if compressed string is greater than uncompressed string
		return in
	}
	countConsecutive := 0
	compressed := []string{}
	splitStr := strings.Split(in, "")
	for i := 0; i < len(splitStr); i++ {
		countConsecutive++

		if i+1 >= len(in) || splitStr[i] != splitStr[i+1] {
			compressed = append(compressed, splitStr[i], strconv.Itoa(countConsecutive))
			countConsecutive = 0
		}
	}

	return strings.Join(compressed, "")
}

func compressedLength(in string) int {
	countConsecutive := 0
	compressedLen := 0
	for i := 0; i < len(in); i++ {
		countConsecutive++

		if i+1 >= len(in) || in[i] != in[i+1] {
			compressedLen = compressedLen + (1 + (len(string(countConsecutive))))
			countConsecutive = 0
		}
	}
	return compressedLen
}

//Given a string, find the length of the longest substring without repeating characters.
func longestSubString(sentence string) int {
	longest := 0
	startIdx := 0
	mapChar := map[byte]int{}
	for i := 0; i < len(sentence); i++ {
		ch := sentence[i]
		if foundIdx, ok := mapChar[ch]; ok {
			startIdx = foundIdx + 1
		}
		if longest < (i - startIdx + 1) {
			longest = i - startIdx + 1
		}
		mapChar[ch] = i
	}
	return longest
}

//think of this exercise like a variation of mancala.
//Starting at the zeroth index, can you navigate to the last index using the number of hops given to you at each index?
//You may use less than the number of hops given to you at each index
//return true if you can, return false if it is impossible
//ex [3,3,1,0,2,0,1] is true
// hop once from 3 to 3, hop three times to two and hop twice to 1!
//ex [3,2,0,0,2,0,1] is false
// the farthest you can hop is to index 4
func advancingArray(arr []int) bool {
	farthest := arr[0]
	for i := 0; i <= farthest; i++ {
		if farthest == len(arr)-1 {
			return true
		}
		if arr[i]+i > farthest {
			farthest = arr[i] + i
		}
	}
	return false
}

//check to see if a passphrase is valid, it must not contain any duplicate words
func passPhrase(passphrase string) bool {
	wordMap := map[string]bool{}
	for _, word := range strings.Split(passphrase, " ") {
		if _, ok := wordMap[word]; ok {
			return false
		}
		wordMap[word] = true
	}
	return true
}

func passPhraseAdvent() int {
	//read from ../practice/advent_input/passphrase_input.txt
	fileBytes, err := ioutil.ReadFile("../practice/advent_input/passphrase_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	splitPhrases := strings.Split(string(fileBytes), "\n")
	count := 0
	for _, phrase := range splitPhrases {
		if passPhrase(phrase) {
			count++
		}
	}
	return count
}

//bonus, check to see if a passphrase is valid, it must not contain any words that are anagrams of eachother
// abcde == dcabe
// func passPhraseAnagram(passphrase string) bool {

// }
