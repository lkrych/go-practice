package goExercises
 
// type queue struct {
// 	data []int
// }

// func (q *queue) add(item int) {
// 	q.data = append(q.data, item)
// }

// func (q *queue) remove() int {
// 	firstItem := q.data[0]
// 	q.data = q.data[1:]
// 	return firstItem
// }

// func (q *queue) peek() int {
// 	return q.data[0]
// }

// func (q *queue) isEmpty() bool {
// 	return len(q.data) == 0
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
 
// //Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
// func isUnique(word string) bool {
// 	//use a map
// }
 
// import "strings"

// func countVowels(sentence string) int {
// 	split := strings.Split(sentence, "")
// 	vowels := "aeiou"
// 	count := 0
// 	for _, chars := range split {
// 		count += strings.Count(vowels, chars)
// 	}

// 	return count
// }
 
// import "strings"

// // If there is a rotation, there is a pivot point
// //regardless of where the pivot is s2 will always be a substring of s1s1
// func stringRotation(s1, s2 string) bool {
// 	if len(s1) == len(s2) && len(s1) > 0 {
// 		s1 = s1 + s1
// 		return strings.Index(s1, s2) > 0
// 	}
// }
 
