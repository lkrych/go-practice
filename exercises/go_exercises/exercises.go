package goExercises

import ("strconv";"strings")
//build a queue data structure
// fifo, has four functions: add, remove, peek, isEmpty 
type queue struct{
	data []int
}

func (q *queue) add (x int) {
	q.data = append(q.data, x)
}

func (q *queue)	remove () int {
	first := q.data[0]
	q.data = q.data[1:]
	return first
}

func (q *queue)	peek  () int {
	return q.data[0]
}

func (q *queue)	isEmpty () bool {
	if len(q.data) == 0 {
		return true
	}
	return false
}
 
//add dashes around an odd number, except on first or last el. Only allow one dash
//ex: 303 => 3-0-3
//ex: 333 => 3-3-3
//ex: 3223 => 3-22-3
func dasherizeNumber(n int) string {
	toString := strconv.Itoa(n%10)
	
	if n / 10 == 0 {
		if isOdd(n) {
			return toString + "-"	
		}
		return toString
	}
	if isOdd(n % 10) {
		toString = "-" + toString + "-"
	}
	ans := dasherizeNumber(n / 10) + toString
	ans = strings.Replace(ans, "--", "-", -1)
	return strings.TrimRight(ans, "-")
}

func isOdd (n int) bool {
	if n % 2 == 0 {
		return false
	}
	return true
}
 
