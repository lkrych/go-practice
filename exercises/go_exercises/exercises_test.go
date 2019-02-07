package goExercises
 
import (
"testing"
)
 

func TestBuyStock(t *testing.T) {
	test1 := []int{3, 3, 1, 0, 2, 0, 1}
	test2 := []int{3, 2, 0, 0, 2, 0, 1, 5}
	test3 := []int{5, 3, 1, 5, 2, 0, 1}
	test4 := []int{5, 3, 2, 4, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	test6 := []int{7, 6, 5, 4, 3, 2}
	if buyStock(test1) != 2 {
		t.Errorf("Expected array %v to be 2", test1)
	}
	if buyStock(test2) != 5 {
		t.Errorf("Expected array %v to be 5", test2)
	}

	if buyStock(test3) != 4 {
		t.Errorf("Expected array %v to be 4", test3)
	}

	if buyStock(test4) != 2 {
		t.Errorf("Expected array %v to be 2", test4)
	}

	if buyStock(test5) != 6 {
		t.Errorf("Expected array %v to be 6", test5)
	}

	if buyStock(test6) != 0 {
		t.Errorf("Expected array %v to be 0", test6)
	}
}
 
func TestDigitalRoot(t *testing.T) {
	test1 := digitalRoot(10)
	test2 := digitalRoot(83)
	test3 := digitalRoot(123)
	test4 := digitalRoot(789)
	test5 := digitalRoot(452982)

	if test1 != 1 {
		t.Errorf("Expected digitalRoot of 10 to be 1, not %v", test1)
	}
	if test2 != 2 {
		t.Errorf("Expected digitalRoot of 83 to be 2, not %v", test2)
	}

	if test3 != 6 {
		t.Errorf("Expected digitalRoot of 123 to be 6, not %v", test3)
	}

	if test4 != 6 {
		t.Errorf("Expected digitalRoot of 789 to be 6, not %v", test4)
	}

	if test5 != 3 {
		t.Errorf("Expected digitalRoot of 452982 to be 3, not %v", test5)
	}
}
 
