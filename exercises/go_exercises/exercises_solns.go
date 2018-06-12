package goExercises
 
//before building the string, you should check to make sure that you need to build the string
// func stringCompression(in string) out string {
// 	if countCompressionLength(in) > len(in) {
// 		return in
// 	}

// 	compressed := []rune{}
// 	countConsec := 0
// 	for i := 0; i < len(in); i++ {
// 		countConsec++

// 		if (i+1 >= len(string) || in[i] != in[i+1]) {
// 			compressed = append(compressed, in[i], countConsec)
// 			countConsec = 0
// 		}
// 	}
// 	return string(compressed)
// }

// func countCompressionLength(in string) bool {
// 	compressedLength := 0
// 	countConsec := 0
// 	for i := 0; i < len(in); i++ {
// 		countConsec++

// 		if (i+1 >= len(string) || in[i] != in[i+1]) {
// 			compressedLength += (1 + len(string(countConsec)))
// 			countConsec = 0
// 		}
// 	}
// }
 
// //Given a string, find the length of the longest substring without repeating characters.
// func lengthOfLongestSubstring(s string) int {
// 	bs := []byte(s)
// 	var maxLen, start int
// 	idxMap := map[byte]int{}
// 	for i := 0; i < len(bs); i++ {
// 			if _, ok := idxMap[bs[i]]; ok && start <= idxMap[bs[i]] {
// 					start = idxMap[bs[i]] + 1
// 			} else {
// 					if maxLen < i - start + 1 {
// 							maxLen = i - start + 1
// 					}
// 			}
// 			idxMap[bs[i]] = i
// 	}
// 	return maxLen
// }
 
// import "math"

// //func keep track of how far you have traveled, if you can reach the end, you made it!
// func advancingArray(arr []int) bool {
// 	farthestTraveled := arr[0]
// 	for i := 0; i <= farthestTraveled; i++ {
// 		if farthestTraveled == len(arr)-1 {
// 			return true
// 		}
// 		farthestTraveled = int(math.Max(float64(farthestTraveled), float64(arr[i]+i)))
// 	}
// 	return false
// }
 
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"strings"
// )

// func passPhrase(passphrase string) bool {
// 	split := strings.Split(passphrase, " ")
// 	wordMap := make(map[string]int)
// 	for _, word := range split {
// 		if _, ok := wordMap[word]; ok {
// 			fmt.Printf("%v is the repeated word in %v \n", word, passphrase)
// 			return false
// 		}
// 		wordMap[word] = 1
// 	}
// 	return true
// }

// func passPhraseAdvent() int {
// 	phraseCount := 0
// 	content, err := ioutil.ReadFile("./advent_input/passphrase_input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	splitByLine := strings.Split(string(content), "\n")
// 	for _, line := range splitByLine {
// 		if passPhrase(line) {
// 			phraseCount++
// 		}
// 	}
// 	return phraseCount
// }

// func passPhraseAnagram(passphrase string) bool {
// 	split := strings.Split(passphrase, " ")
// 	wordMap := make(map[string]int)
// 	for _, word := range split {
// 		sorted := SortString(word)
// 		if _, ok := wordMap[sorted]; ok {
// 			return false
// 		}
// 		wordMap[sorted] = 1
// 	}
// 	return true
// }
 
