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

 
// import (
// 	"strconv"
// 	"strings"
// )

// func dasherizeNumber(n int) string {
// 	dasherized := []string{}
// 	toString := strconv.Itoa(n)
// 	for i, el := range strings.Split(toString, "") {

// 		if int, _ := strconv.Atoi(el); int%2 != 0 {
// 			if i == 0 {
// 				dasherized = append(dasherized, el, "-")
// 			} else if i == len(toString)-1 {
// 				dasherized = append(dasherized, "-", el)
// 			} else {
// 				dasherized = append(dasherized, "-", el, "-")
// 			}
// 		} else {
// 			dasherized = append(dasherized, el)
// 		}
// 	}
// 	joined := strings.Join(dasherized, "")
// 	return strings.Replace(joined, "--", "-", -1)
// }
 
