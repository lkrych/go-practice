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
