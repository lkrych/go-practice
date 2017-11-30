func TestBuyStockTwice(t *testing.T) {
	test1 := []int{12, 11, 13, 9, 12, 8, 14, 13, 15}
	test2 := []int{3, 2, 0, 0, 2, 0, 1, 5}
	test3 := []int{5, 3, 1, 5, 2, 0, 1}
	test4 := []int{5, 3, 2, 4, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	test6 := []int{7, 6, 5, 4, 3, 2}
	if buyStockTwice(test1) != 10 {
		t.Errorf("Expected array %v to be 2", test1)
	}
	if buyStockTwice(test2) != 7 {
		t.Errorf("Expected array %v to be 5", test2)
	}

	if buyStockTwice(test3) != 5 {
		t.Errorf("Expected array %v to be 4", test3)
	}

	if buyStockTwice(test4) != 2 {
		t.Errorf("Expected array %v to be 2", test4)
	}

	if buyStockTwice(test5) != 9 {
		t.Errorf("Expected array %v to be 6", test5)
	}

	if buyStockTwice(test6) != 0 {
		t.Errorf("Expected array %v to be 0", test6)
	}
}