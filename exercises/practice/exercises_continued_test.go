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

func TestStreamParser(t *testing.T) {
	bracket1, garbage1 := streamParser("{}")
	bracket2, garbage2 := streamParser("{{{}}}")
	bracket3, garbage3 := streamParser("{{},{}}")
	bracket4, garbage4 := streamParser("{{{},{},{{}}}}")
	bracket5, garbage5 := streamParser("{<a>,<a>,<a>,<a>}")
	bracket6, garbage6 := streamParser("{{<ab>},{<ab>},{<ab>},{<ab>}}")
	bracket7, garbage7 := streamParser("{{<!!>},{<!!>},{<!!>},{<!!>}}")
	bracket8, garbage8 := streamParser("{{<a!>},{<a!>},{<a!>},{<ab>}}")
	bracket9, garbage9 := streamParserAdvent("./advent_input/stream_input.txt")

	if bracket1 != 1 {
		t.Errorf("Expected answer to be 1, not %v", bracket1)
	}

	if garbage1 != 0 {
		t.Errorf("Expected answer to be 0, not %v", garbage1)
	}

	if bracket2 != 6 {
		t.Errorf("Expected answer to be 6, not %v", bracket2)
	}

	if garbage2 != 0 {
		t.Errorf("Expected answer to be 0, not %v", garbage2)
	}

	if bracket3 != 5 {
		t.Errorf("Expected answer to be 5, not %v", bracket3)
	}

	if garbage3 != 0 {
		t.Errorf("Expected answer to be 0, not %v", garbage3)
	}

	if bracket4 != 16 {
		t.Errorf("Expected answer to be 16, not %v", bracket4)
	}

	if garbage4 != 0 {
		t.Errorf("Expected answer to be 0, not %v", garbage4)
	}

	if bracket5 != 1 {
		t.Errorf("Expected answer to be 1, not %v", bracket5)
	}

	if garbage5 != 4 {
		t.Errorf("Expected answer to be 4, not %v", garbage5)
	}

	if bracket6 != 9 {
		t.Errorf("Expected answer to be 9, not %v", bracket6)
	}

	if garbage6 != 8 {
		t.Errorf("Expected answer to be 8, not %v", garbage6)
	}

	if bracket7 != 9 {
		t.Errorf("Expected answer to be 9, not %v", bracket7)
	}

	if garbage7 != 0 {
		t.Errorf("Expected answer to be 0, not %v", garbage7)
	}

	if bracket8 != 3 {
		t.Errorf("Expected answer to be 3, not %v", bracket8)
	}

	if garbage8 != 17 {
		t.Errorf("Expected answer to be 17, not %v", garbage8)
	}

	if bracket9 != 12505 {
		t.Errorf("Expected answer to be 12505, not %v", bracket9)
	}

	if garbage9 != 17 {
		t.Errorf("Expected answer to be 17, not %v", garbage9)
	}
}

func TestKnotHash(t *testing.T) {
	// test1 := knotHashMinimal([]int{3, 4, 1, 5})
	// test2 := knotHashMinimal([]int{3, 4})
	// test3 := knotHashAdvent("./advent_input/knot_input.txt")
	test4 := knotHashAdventAdvanced("./advent_input/knot_input.txt")

	// if test1 != 12 {
	// 	t.Errorf("Expected answer to be 12, not %v", test1)
	// }

	// if test2 != 12 {
	// 	t.Errorf("Expected answer to be 12, not %v", test2)
	// }

	// if test3 != 32 {
	// 	t.Errorf("Expected answer to be 32, not %v", test3)
	// }

	if test4 != "googlywoogly" {
		t.Errorf("Expected answer to be googlywoogly, not %v", test4)
	}
}

func TestHexMaze(t *testing.T) {
	test1 := hexMaze([]string{"ne", "ne", "ne"})
	test2 := hexMaze([]string{"ne", "ne", "sw", "sw"})
	test3 := hexMaze([]string{"ne", "ne", "s", "s"})
	test4 := hexMaze([]string{"se", "sw", "se", "sw", "sw"})
	test5 := hexMazeAdvent("./advent_input/hex_maze_input.txt")

	if test1 != 3 {
		t.Errorf("Expected answer to be 3, not %v", test1)
	}

	if test2 != 0 {
		t.Errorf("Expected answer to be 0, not %v", test2)
	}

	if test3 != 2 {
		t.Errorf("Expected answer to be 2, not %v", test3)
	}

	if test4 != 3 {
		t.Errorf("Expected answer to be 3, not %v", test4)
	}

	if test5 != 3 {
		t.Errorf("Expected answer to be 3, not %v", test5)
	}

}

func TestDigitalPlumber(t *testing.T) {
	test1 := digitalPlumber("./advent_input/digital_plumber_beginner.txt")
	// test2 := digitalPlumber("./advent_input/digital_plumber_input.txt")

	if test1 != 3 {
		t.Errorf("Expected answer to be 3, not %v", test1)
	}

	// if test2 != 0 {
	// 	t.Errorf("Expected answer to be 0, not %v", test2)
	// }
}

func TestPhoneMnemonics(t *testing.T) {
	test1 := phoneMnemonic(2276696)
	test2 := phoneMnemonic(2639874)
	test3 := phoneMnemonic(2635901)

	if len(test1) != 3 {
		t.Errorf("Expected all permutations of number to be x, not %v", test1)
	}

	if len(test2) != 0 {
		t.Errorf("Expected all permutations of number to be x, not %v", test2)
	}

	if len(test3) != 2 {
		t.Errorf("Expected all permutations of number to be x, not %v", test3)
	}

}
