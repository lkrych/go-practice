
i
func TestBuyStockTwice(t *testing.T) {
	test1 := []int{12, 11, 13, 9, 12, 8, 14, 13, 15}
	test2 := []int{3, 2, 0, 0, 2, 0, 1, 5}
	test3 := []int{5, 3, 1, 5, 2, 0, 1}
	test4 := []int{5, 3, 2, 4, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	test6 := []int{7, 6, 5, 4, 3, 2}
	ans1 := 10
	ans2 := 7
	ans3 := 5
	ans4 := 2
	ans5 := 9
	ans6 := 0
	if buyStockTwice(test1) != ans1 {
		t.Errorf("Expected array %v to be 2, not %v", test1, ans1)
	}
	if buyStockTwice(test2) != ans2 {
		t.Errorf("Expected array %v to be 5, not %v", test2, ans2)
	}

	if buyStockTwice(test3) != ans3 {
		t.Errorf("Expected array %v to be 4, not %v", test3, ans3)
	}

	if buyStockTwice(test4) != ans4 {
		t.Errorf("Expected array %v to be 2, not %v", test4, ans4)
	}

	if buyStockTwice(test5) != ans5 {
		t.Errorf("Expected array %v to be 6, not %v", test5, ans5)
	}

	if buyStockTwice(test6) != ans6 {
		t.Errorf("Expected array %v to be 0, not %v", test6, ans6)
	}
}