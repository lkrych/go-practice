package main

import "testing"

func TestReversePolish(t *testing.T) {
	test1 := reversePolish("1 2 3 + 4 5 * * +")
	test2 := reversePolish("33 3 * 101 +")
	test3 := reversePolish("2 452 / 2 *")
	test4 := reversePolish("8 8 8 8 + + +")
	test5 := reversePolish("17 4 * 3 6 + +")
	test6 := reversePolish("45 120 - 78 *")

	ans1 := 101
	ans2 := 200
	ans3 := 452
	ans4 := 32
	ans5 := 77
	ans6 := 5850

	if test1 != ans1 {
		t.Errorf("Expected (2 + 3) * (4 * 5) + 1 to be %v, not %v", ans1, test1)
	}
	if test2 != ans2 {
		t.Errorf("Expected (33 * 3) + 101 to be %v, not %v", ans2, test2)
	}
	if test3 != ans3 {
		t.Errorf("Expected (452 / 2) * 2 to be %v, not %v", ans3, test3)
	}
	if test4 != ans4 {
		t.Errorf("Expected 8+8+8+8 to be %v, not %v", ans4, test4)
	}
	if test5 != ans5 {
		t.Errorf("Expected (17 * 4) + (3 + 6) to be %v, not %v", ans5, test5)
	}
	if test6 != ans6 {
		t.Errorf("Expected (120 - 45) * 78 to be %v, not %v", ans6, test6)
	}
}
