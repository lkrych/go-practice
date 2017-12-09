func TestRegister(t *testing.T) {
	currMax1, max1 := register("./advent_input/register_beginner_input.txt")
	currMax2, max2 := register("./advent_input/register_input.txt")

	if currMax1 != 1 {
		t.Errorf("Expected currentMax of test 1 to be 1, not %v", currMax1)
	}

	if max1 != 10 {
		t.Errorf("Expected max of test 1 to be 10, not %v", max1)
	}

	if currMax2 != 4888 {
		t.Errorf("Expected currentMax of test 1 to be 4888, not %v", currMax2)
	}

	if max2 != 7774 {
		t.Errorf("Expected max of test 1 to be 7774, not %v", max2)
	}
}