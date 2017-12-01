func TestReverseCaptcha(t *testing.T) {
	test1 := reverseCaptcha([]int{1, 1, 2, 2})
	test2 := reverseCaptcha([]int{1, 1, 1, 1})
	test3 := reverseCaptcha([]int{1, 2, 3, 4})
	test4 := reverseCaptcha([]int{9, 1, 2, 9})
	test5 := reverseCaptcha([]int{5, 4, 2, 4, 2, 2, 7, 7, 3, 2, 5})
	test6 := adventCaptcha()

	if test1 != 3 {
		t.Errorf("Expected answer to be 3, not %v", test1)
	}
	if test2 != 4 {
		t.Errorf("Expected answer to be 4, not %v", test2)
	}

	if test3 != 0 {
		t.Errorf("Expected answer to be 0, not %v", test3)
	}

	if test4 != 9 {
		t.Errorf("Expected answer to be 9, not %v", test4)
	}

	if test5 != 14 {
		t.Errorf("Expected answer to be 14, not %v", test5)
	}

	if test6 != 1102 {
		t.Errorf("Expected answer to be 1102, not %v", test6)
	}

}