func TestMemoryBank(t *testing.T) {
	test1 := memoryBank([]int{0, 2, 7, 0})
	test2 := memoryBank([]int{0, 1, 0, 0})
	test3 := memoryBankAdvent()

	if test1 != 5 {
		t.Errorf("Expected answer to be 5, not %v", test1)
	}

	if test2 != 4 {
		t.Errorf("Expected answer to be 4, not %v", test2)
	}

	if test3 != 11137 {
		t.Errorf("Expected answer to be 11137, not %v", test3)
	}
}