package goExercises

import (
	"regexp"
	"strings"
)

//build a queue data structure
type queue struct {
	data []int
}

func (q *queue) add(n int) {
	q.data = append(q.data, n)
}

func (q *queue) remove() int {
	first := q.data[0]
	q.data = q.data[1:]
	return first
}

func (q *queue) peek() int {
	return q.data[0]
}

func (q *queue) isEmpty() bool {
	return len(q.data) < 1
}

//return a sorted arr using bubblesort!
func bubbleSort(arr []int) []int {
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				sorted = false
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}

//Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
func isUnique(word string) bool {
	chMap := map[string]bool{}
	for _, ch := range strings.Split(word, "") {
		if ok, _ := chMap[ch]; ok {
			return false
		} else {
			chMap[ch] = true
		}
	}
	return true
}

//return num of vowels in sentence
func countVowels(sentence string) int {
	vowelCount := 0
	vowelRegexp := regexp.MustCompile("[aeiou]")
	for _, ch := range strings.Split(sentence, "") {
		if vowelRegexp.Match([]byte(ch)) {
			vowelCount++
		}
	}
	return vowelCount
}

//assume you have a method isSubstring which checks if one word is asubstring of another
//given two strings s1 and s2, write code to check if s2 is a rotation of s1 using
//only 1 call to isSubstring
//ex: waterbottle is a rotation of erbottlewat
func stringRotation(s1, s2 string) bool {
	newString := s2 + s2
	s1regex := regexp.MustCompile(s1)
	return s1regex.Match([]byte(newString))
}
