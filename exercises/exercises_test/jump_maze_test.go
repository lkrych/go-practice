func TestJumpMaze(t *testing.T) {
	test1 := jumpMaze([]int{0, 3, 0, 1, -3})
	test2 := jumpMaze([]int{1, 3, 0, 1, -3})
	test3 := jumpMaze([]int{0, 2, 0, 1, -3})
	test4 := jumpMazeAdvent()

	if test1 != 5 {
		t.Errorf("Expected answer to be 5, not %v", test1)
	}

	if test2 != 4 {
		t.Errorf("Expected answer to be 4, not %v", test2)
	}

	if test3 != 10 {
		t.Errorf("Expected answer to be 10, not %v", test3)
	}

	if test4 != 388611 {
		t.Errorf("Expected answer to be 388611, not %v", test4)
	}
}