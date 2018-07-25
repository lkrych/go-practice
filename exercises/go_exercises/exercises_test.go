package goExercises

import (
	"testing"
)

func TestCountVowels(t *testing.T) {
	test1 := countVowels("abcd")
	test2 := countVowels("colour")
	test3 := countVowels("cecilia nutley")

	if test1 != 1 {
		t.Errorf("Expected countVowels of abcd to be 1, not %v", test1)
	}

	if test2 != 3 {
		t.Errorf("Expected countVowels of colour to be 3, not %v", test2)
	}

	if test3 != 6 {
		t.Errorf("Expected countVowels of mom to be 6, not %v", test3)
	}
}

func TestCreatingAQueue(t *testing.T) {
	q := &queue{
		data: []int{1, 2, 3, 4, 5, 6},
	}
	testremove := q.remove()
	testremove2 := q.remove()

	if len(q.data) != 4 {
		t.Errorf("Expected remove to return elements from the top of the queue: %v, removeped: %v", q.data, testremove)
	}
	if q.peek() != 3 {
		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
	}
	if len(q.data) != 4 {
		t.Errorf("Expected peek to not delete elements from the top of the queue")
	}
	if q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue has %v elements", len(q.data))
	}

	q.add(testremove)
	if len(q.data) != 5 {
		t.Errorf("Expected remove to return elements from the top of the queue: %v, removeped: %v", q.data, testremove2)
	}
	if q.peek() != 3 {
		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
	}
	if len(q.data) != 5 {
		t.Errorf("Expected peek to not delete elements from the top of the queue")
	}
	if q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue has %v elements\n", len(q.data))
	}

	q.add(testremove2)
	if len(q.data) != 6 {
		t.Errorf("Expected push to return elements to the top of the queue")
	}
	if q.peek() != 3 {
		t.Errorf("Expected peek to return the element at the top of the queue, not %v", q.peek())
	}
	if len(q.data) != 6 {
		t.Errorf("Expected peek to not delete elements from the top of the queue")
	}
	if q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue has %v elements\n", len(q.data))
	}

	for i := 0; i < 6; i++ {
		q.remove()
	}
	if !q.isEmpty() {
		t.Errorf("The isEmpty function doesnt work because the queue should be empty")
	}
}
