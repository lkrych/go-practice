package main

import (
	"math/big"
	"testing"
)

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

func TestAddLinkedLists(t *testing.T) {
	list := createLinkedList(102)
	num := createNumber(list)

	if num != 102 {
		t.Errorf("Expected helper functions to work properly. The list is %v, the number is %v", list, num)
	}

	//real tests
	l1 := createLinkedList(102)
	l2 := createLinkedList(245)
	test1 := createNumber(addTwoNumbers(l1, l2))

	l1 = createLinkedList(87)
	l2 = createLinkedList(349)
	test2 := createNumber(addTwoNumbers(l1, l2))

	l1 = createLinkedList(1003)
	l2 = createLinkedList(549)
	test3 := createNumber(addTwoNumbers(l1, l2))

	// test big numbers
	b1 := big.Int(2432432432432432432432432432432432432432432432432432432432439)
	b2 := big.Int(5642432432432432432432432432432432432432432432432432432439999)
	l1 = createLinkedList(b1)
	l2 = createLinkedList(b2)
	test4 := createNumber(addTwoNumbers(l1, l2))

	if test1 != 347 {
		t.Errorf("Expected answer to be 347, not %v", test1)
	}

	if test2 != 436 {
		t.Errorf("Expected answer to be 347, not %v", test2)
	}

	if test3 != 1552 {
		t.Errorf("Expected answer to be 347, not %v", test3)
	}

	if test4 != 1552 {
		t.Errorf("Expected answer to be 347, not %v", test4)
	}
}

func TestPower(t *testing.T) {
	test0 := power(2, 0)
	test1 := power(10, 2)
	test2 := power(3, 3)
	test3 := power(4, 4)
	test4 := power(10, 3)
	test5 := power(2, 5)

	if test0 != 1 {
		t.Errorf("Expected answer to be 1, not %v", test1)
	}
	if test1 != 100 {
		t.Errorf("Expected answer to be 100, not %v", test1)
	}
	if test2 != 27 {
		t.Errorf("Expected answer to be 27, not %v", test2)
	}
	if test3 != 256 {
		t.Errorf("Expected answer to be 256, not %v", test3)
	}
	if test4 != 1000 {
		t.Errorf("Expected answer to be 1000, not %v", test4)
	}
	if test5 != 32 {
		t.Errorf("Expected answer to be 32, not %v", test5)
	}
}
