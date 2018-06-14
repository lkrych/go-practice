package main

import "log"

type stack struct {
	data []int
}

func (s *stack) pop() int {
	popped := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return popped
}

func (s *stack) push(item int) {
	s.data = append(s.data, item)
}

func (s *stack) peek() int {
	return s.data[len(s.data)-1]
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

type queue struct {
	data []int
}

func (q *queue) add(item int) {
	q.data = append(q.data, item)
}

func (q *queue) remove() int {
	firstItem := q.data[0]
	q.data = q.data[1:]
	return firstItem
}

func (q *queue) peek() int {
	return q.data[0]
}

func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}

//How would you design a stack which in addition to push and pop, has a function
// min which returns the minimum element? Push, pop and min should operate in O(1)

//use a stack
type stackWithMin struct {
	data     []int
	minStack stack
}

func (s *stackWithMin) pop() int {
	popped := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	if popped == s.min() { //if you've popped the min, move it off the minStack!
		s.minStack.pop()
	}
	return popped
}

func (s *stackWithMin) push(item int) {
	s.data = append(s.data, item)
	if item <= s.min() {
		s.minStack.push(item)
	}
}

func (s *stackWithMin) min() int {
	if s.minStack.isEmpty() {
		return 100000000000000000 //really large number
	}
	return s.minStack.peek()
}

//using a separate struct
// type stackVal struct {
// 	val int
// 	min int
// }

// type stackWithMin struct {
// 	data []stackVal
// }

// func (s *stackWithMin) pop() int {
// 	popped := s.data[len(s.data)-1]
// 	s.data = s.data[:len(s.data)-1]
// 	return popped.val
// }

// func (s *stackWithMin) push(item int) {
// 	currentMin := s.min()
// 	if item < currentMin {
// 		currentMin = item
// 	}
// 	val := stackVal{
// 		val: item,
// 		min: currentMin,
// 	}
// 	s.data = append(s.data, val)
// }

// func (s *stackWithMin) min() int {
// 	if s.isEmpty() {
// 		return 100000000000000000 //really large number
// 	}
// 	data := s.data[len(s.data)-1]
// 	return data.min
// }

// func (s *stackWithMin) peek() int {
// 	data := s.data[len(s.data)-1]
// 	return data.val
// }

// func (s *stackWithMin) isEmpty() bool {
// 	return len(s.data) == 0
// }

// type stackOfStacks struct {
// 	capacity int
// 	stacks   []stack
// }

// func (s *stackOfStacks) push(item int) {
// 	if len(s.stacks) == 0 || len(s.peek().data) == s.capacity {
// 		newStack := &stack{}
// 		newStack.push(item)
// 		s.stacks = append(s.stacks, *newStack)
// 	} else {
// 		currentStack := s.peek()
// 		currentStack.push(item)
// 	}
// }

// func (s *stackOfStacks) pop() int {
// 	var popped int
// 	currentStack := s.peek()
// 	if len(currentStack.data) > 0 {
// 		popped = currentStack.pop()
// 		if len(currentStack.data) == 0 {
// 			s.stacks = s.stacks[:len(s.stacks)-1] //remove empty stack
// 		}
// 	} else if len(s.stacks) == 0 { //make sure there is a stack
// 		log.Fatal("There are no more stacks!")
// 	} else {
// 		s.stacks = s.stacks[:len(s.stacks)-1] //remove empty stack
// 		popped = s.peek().pop()
// 	}
// 	return popped

// }

// func (s *stackOfStacks) peek() *stack {
// 	return &s.stacks[len(s.stacks)-1]
// }

type myQueue struct {
	in  stack
	out stack
}

func (q *myQueue) add(item int) {
	q.in.push(item)
}

func (q *myQueue) remove() int {
	if q.out.isEmpty() {
		q.reshuffle()
	}
	return q.out.pop()
}

func (q *myQueue) peek() int {
	if q.isEmpty() {
		log.Fatal("queue is empty!")
	} else if q.out.isEmpty() {
		q.reshuffle()
	}
	return q.out.peek()
}

func (q *myQueue) reshuffle() {
	for !q.in.isEmpty() {
		q.out.push(q.in.pop())
	}
}

func (q *myQueue) isEmpty() bool {
	return len(q.in.data) == 0 && len(q.out.data) == 0
}
