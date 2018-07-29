package goExercises
 
// //Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
// func isUnique(word string) bool {
// 	//use a map
// }
 
// func bubbleSort(arr []int) []int {
// 	sorted := false
// 	for !sorted {
// 		sorted = true
// 		for i := 0; i < len(arr)-1; i++ {
// 			if arr[i] > arr[i+1] {
// 				arr[i], arr[i+1] = arr[i+1], arr[i]
// 				sorted = false
// 			}
// 		}
// 	}
// 	return arr
// }
 
// import "strings"

// func numRepeats(word string) int {
// 	lettersMap := map[string]int{}
// 	for _, ch := range strings.Split(word, "") {
// 		if ch == " " {
// 			continue
// 		}
// 		if lettersMap[ch] >= 1 {
// 			lettersMap[ch]++
// 		} else {
// 			lettersMap[ch] = 1
// 		}
// 	}
// 	count := 0
// 	for _, v := range lettersMap {
// 		if v > 1 {
// 			count++
// 		}
// 	}
// 	return count
// }
 
