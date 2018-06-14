package main

import (
	"testing"
)

// func TestCreatingAQueue(t *testing.T) {
// 	q := &queue{
// 		data: []int{1, 2, 3, 4, 5, 6},
// 	}
// 	testremove := q.remove()
// 	testremove2 := q.remove()

// 	if len(q.data) != 4 {
// 		t.Errorf("Expected remove to return elements from the top of the queue: %v, removeped: %v", q.data, testremove)
// 	}
// 	if q.peek() != 3 {
// 		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
// 	}
// 	if len(q.data) != 4 {
// 		t.Errorf("Expected peek to not delete elements from the top of the queue")
// 	}
// 	if q.isEmpty() {
// 		t.Errorf("The isEmpty function doesnt work because the queue has %v elements", len(q.data))
// 	}

// 	q.add(testremove)
// 	if len(q.data) != 5 {
// 		t.Errorf("Expected remove to return elements from the top of the queue: %v, removeped: %v", q.data, testremove2)
// 	}
// 	if q.peek() != 3 {
// 		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
// 	}
// 	if len(q.data) != 5 {
// 		t.Errorf("Expected peek to not delete elements from the top of the queue")
// 	}
// 	if q.isEmpty() {
// 		t.Errorf("The isEmpty function doesnt work because the queue has %v elements\n", len(q.data))
// 	}

// 	q.add(testremove2)
// 	if len(q.data) != 6 {
// 		t.Errorf("Expected push to return elements to the top of the queue")
// 	}
// 	if q.peek() != 3 {
// 		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
// 	}
// 	if len(q.data) != 6 {
// 		t.Errorf("Expected peek to not delete elements from the top of the queue")
// 	}
// 	if q.isEmpty() {
// 		t.Errorf("The isEmpty function doesnt work because the queue has %v elements\n", len(q.data))
// 	}

// 	for i := 0; i < 6; i++ {
// 		q.remove()
// 	}
// 	if !q.isEmpty() {
// 		t.Errorf("The isEmpty function doesnt work because the queue should be empty")
// 	}
// }

// func TestCreatingAStack(t *testing.T) {
// 	s := &stack{
// 		data: []int{1, 2, 3, 4, 5, 6},
// 	}
// 	testPop := s.pop()
// 	testPop2 := s.pop()

// 	if len(s.data) != 4 {
// 		t.Errorf("Expected pop to return elements from the top of the stack: %v, popped: %v", s.data, testPop)
// 	}
// 	if s.peek() != 4 {
// 		t.Errorf("Expected peek to return the element at the top of the stack, not %v", s.peek())
// 	}
// 	if len(s.data) != 4 {
// 		t.Errorf("Expected peek to not delete elements from the top of the stack")
// 	}
// 	if s.isEmpty() {
// 		t.Errorf("The isEmpty function doesnt work because the stack has %v elements", len(s.data))
// 	}

// 	s.push(testPop)
// 	if len(s.data) != 5 {
// 		t.Errorf("Expected pop to return elements from the top of the stack: %v, popped: %v", s.data, testPop2)
// 	}
// 	if s.peek() != 6 {
// 		t.Errorf("Expected peek to return the element at the top of the stack, not %v", s.peek())
// 	}
// 	if len(s.data) != 5 {
// 		t.Errorf("Expected peek to not delete elements from the top of the stack")
// 	}
// 	if s.isEmpty() {
// 		t.Errorf("The isEmpty function doesnt work because the stack has %v elements\n", len(s.data))
// 	}

// 	s.push(testPop2)
// 	if len(s.data) != 6 {
// 		t.Errorf("Expected push to return elements to the top of the stack")
// 	}
// 	if s.peek() != 5 {
// 		t.Errorf("Expected peek to return the element at the top of the stack, not %v", s.peek())
// 	}
// 	if len(s.data) != 6 {
// 		t.Errorf("Expected peek to not delete elements from the top of the stack")
// 	}
// 	if s.isEmpty() {
// 		t.Errorf("The isEmpty function doesnt work because the stack has %v elements\n", len(s.data))
// 	}

// 	for i := 0; i < 6; i++ {
// 		s.pop()
// 	}
// 	if !s.isEmpty() {
// 		t.Errorf("The isEmpty function doesnt work because the stack should be empty")
// 	}
// }

// func TestCreatingAMinStack(t *testing.T) {
// 	s := &stackWithMin{}
// 	for i := 1; i < 7; i++ {
// 		s.push(i)
// 	}
// 	testPop := s.pop()
// 	testPop2 := s.pop()

// 	if len(s.data) != 4 {
// 		t.Errorf("Expected pop to return elements from the top of the stack: %v, popped: %v", s.data, testPop)
// 	}
// 	if s.min() != 1 {
// 		t.Errorf("Expected min to return the min el in the stack: %v", s.min())
// 	}

// 	s.push(testPop)
// 	if len(s.data) != 5 {
// 		t.Errorf("Expected pop to return elements from the top of the stack: %v, popped: %v", s.data, testPop2)
// 	}
// 	if s.min() != 1 {
// 		t.Errorf("Expected min to return the min el in the stack: %v", s.min())
// 	}

// 	s.push(testPop2)
// 	if len(s.data) != 6 {
// 		t.Errorf("Expected push to return elements to the top of the stack")
// 	}

// 	if s.min() != 1 {
// 		t.Errorf("Expected min to return the min el in the stack: %v", s.min())
// 	}

// 	s.push(0)
// 	if len(s.data) != 7 {
// 		t.Errorf("Expected push to return elements to the top of the stack")
// 	}

// 	if s.min() != 0 {
// 		t.Errorf("Expected min to return the min el in the stack: %v", s.min())
// 	}

// }

// func TestCreatingAStackOfStacks(t *testing.T) {
// 	s := &stackOfStacks{
// 		capacity: 2,
// 	}
// 	for i := 1; i < 8; i++ {
// 		s.push(i)
// 	}

// 	if len(s.stacks) != 4 {
// 		t.Errorf("Expected stack of stacks to have: 4 stacks not %v stacks", len(s.stacks))
// 		fmt.Printf("stacks: %v \n", s.stacks)
// 	}

// 	testPop := s.pop()

// 	if testPop != 7 {
// 		fmt.Println(testPop)
// 		t.Errorf("Expected stackOfStacks pop to work like stack pop")
// 	}

// 	if len(s.stacks) != 3 {
// 		t.Errorf("Expected stack of stacks to have: 3 stacks not %v stacks", len(s.stacks))
// 		fmt.Printf("stacks: %v \n", s.stacks)
// 	}

// 	s.push(testPop)
// 	if len(s.stacks) != 4 {
// 		t.Errorf("Expected stack of stacks to have: 4 stacks not %v stacks", len(s.stacks))
// 		fmt.Printf("stacks: %v \n", s.stacks)
// 	}

// 	s.push(0)
// 	s.push(99)
// 	if len(s.stacks) != 5 {
// 		t.Errorf("Expected stack of stacks to have: 5 stacks not %v stacks", len(s.stacks))
// 		fmt.Printf("stacks: %v \n", s.stacks)
// 	}

// }

func TestCreatingAQueueOfStacks(t *testing.T) {
	q := &myQueue{}
	for i := 1; i < 7; i++ {
		q.add(i)
	}

	testremove := q.remove()
	testremove2 := q.remove()

	if q.peek() != 3 {
		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
	}

	if q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue is not empty: %v", q)
	}

	q.add(testremove)

	if q.peek() != 3 {
		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
	}

	if q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue is not empty: %v", q)
	}

	q.add(testremove2)

	if q.peek() != 3 {
		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
	}

	if q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue is not empty: %v", q)
	}

	for i := 0; i < 6; i++ {
		q.remove()
	}
	if !q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue should be empty")
	}
}
