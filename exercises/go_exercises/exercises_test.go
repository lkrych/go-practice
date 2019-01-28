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
 

func TestBuyStock(t *testing.T) {
	test1 := []int{3, 3, 1, 0, 2, 0, 1}
	test2 := []int{3, 2, 0, 0, 2, 0, 1, 5}
	test3 := []int{5, 3, 1, 5, 2, 0, 1}
	test4 := []int{5, 3, 2, 4, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	test6 := []int{7, 6, 5, 4, 3, 2}
	if buyStock(test1) != 2 {
		t.Errorf("Expected array %v to be 2, not %v", test1, buyStock(test1))
	}
	if buyStock(test2) != 5 {
		t.Errorf("Expected array %v to be 5, not %v", test2, buyStock(test2))
	}

	if buyStock(test3) != 4 {
		t.Errorf("Expected array %v to be 4, not %v", test3, buyStock(test3))
	}

	if buyStock(test4) != 2 {
		t.Errorf("Expected array %v to be 2, not %v", test4, buyStock(test4))
	}

	if buyStock(test5) != 6 {
		t.Errorf("Expected array %v to be 6, not %v", test5, buyStock(test5))
	}

	if buyStock(test6) != 0 {
		t.Errorf("Expected array %v to be 0, not %v", test6, buyStock(test6))
	}
}
 
