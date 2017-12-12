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