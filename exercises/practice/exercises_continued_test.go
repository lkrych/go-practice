package main

import "testing"

func TestRecursiveCircus(t *testing.T) {
	test1 := recursiveCircus("./advent_input/recursive_circus_beginner_input.txt")
	test2 := recursiveCircus("./advent_input/recursive_circus_input.txt")

	if test1 != "tknk" {
		t.Errorf("Expected answer to be 'tknk', not %v", test1)
	}

	if test2 != "fknj" {
		t.Errorf("Expected answer to be 'fknj', not %v", test2)
	}
}

func TestRecursiveCircusTree(t *testing.T) {
	// test1 := recursiveCircusTree("./advent_input/recursive_circus_beginner_input.txt")
	test2 := recursiveCircusTree("./advent_input/recursive_circus_input.txt")

	// if test1 != 3 {
	// 	t.Errorf("Expected answer to be 3, not %v", test1)
	// }

	if test2 != 3 {
		t.Errorf("Expected answer to be 'fknj', not %v", test2)
	}
}

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
