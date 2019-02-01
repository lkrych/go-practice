package goExercises
 
import (
"testing"
)
 
func TestNthPrime(t *testing.T) {
	test1 := nthPrime(3)
	test2 := nthPrime(4)
	test3 := nthPrime(15)
	test4 := nthPrime(28)
	test5 := nthPrime(77)

	if test1 != 5 {
		t.Errorf("Expected nthPrime of 3 to be 5, not %v", test1)
	}

	if test2 != 7 {
		t.Errorf("Expected nthPrime of 4 to be 7, not %v", test2)
	}

	if test3 != 47 {
		t.Errorf("Expected nthPrime of 15 to be 47, not %v", test3)
	}

	if test4 != 107 {
		t.Errorf("Expected nthPrime of 28 to be 107, not %v", test4)
	}

	if test5 != 389 {
		t.Errorf("Expected nthPrime of 77 to be 389, not %v", test5)
	}
}
 
func TestScrambleString(t *testing.T) {
	test1 := scrambleString("abcd", []int{3, 1, 2, 0})
	test2 := scrambleString("markov", []int{5, 3, 1, 4, 2, 0})

	if test1 != "dbca" {
		t.Errorf("Expected scrambleString of 'abcd' with inputs [3,1,2,0] to be 'dcba', not %v", test1)
	}

	if test2 != "vkaorm" {
		t.Errorf("Expected scrambleString of 'markov' with inputs [5,3,1,4,2,0] to be 'vkaorm', not %v", test2)
	}
}
 
