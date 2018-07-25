package goExercises

import (
	"regexp"
	"strings"
)

//return num of vowels in sentence
func countVowels(sentence string) int {
	re := regexp.MustCompile("[aeiou]")
	vowelCount := 0
	for _, ch := range strings.Split(sentence, "") {
		if ok := re.Match([]byte(ch)); ok {
			vowelCount++
		}
	}
	return vowelCount
}

//build a queue data structure
type queue struct {
	data []int
}

//isEmpty, peek, add, remove

func (q *queue) isEmpty() bool {
	return len(q.data) < 1
}

func (q *queue) peek() int {
	return q.data[0]
}

func (q *queue) add(n int) {
	q.data = append(q.data, n)
}

func (q *queue) remove() int {
	first := q.data[0]
	q.data = q.data[1:]
	return first
}
