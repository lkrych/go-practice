package goExercises
 
import (
"fmt"
"testing"
"github.com/google/go-cmp/cmp"
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
 
func TestAdvancingArray(t *testing.T) {
	test1 := []int{3, 3, 1, 0, 2, 0, 1}
	test2 := []int{3, 2, 0, 0, 2, 0, 1}
	test3 := []int{5, 3, 1, 0, 2, 0, 1}
	test4 := []int{0, 3, 1, 0, 2, 0, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	if advancingArray(test1) != true {
		t.Errorf("Expected array %v to be true", test1)
	}
	if advancingArray(test2) != false {
		t.Errorf("Expected array %v to be false", test2)
	}

	if advancingArray(test3) != true {
		t.Errorf("Expected array %v to be true", test3)
	}

	if advancingArray(test4) != false {
		t.Errorf("Expected array %v to be false", test4)
	}

	if advancingArray(test5) != true {
		t.Errorf("Expected array %v to be true", test5)
	}
}
 
