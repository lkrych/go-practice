package goExercises

import (
	"testing"
)

func TestDigitalRoot(t *testing.T) {
	test1 := digitalRoot(10)
	test2 := digitalRoot(83)
	test3 := digitalRoot(123)
	test4 := digitalRoot(789)
	test5 := digitalRoot(452982)

	if test1 != 1 {
		t.Errorf("Expected digitalRoot of 10 to be 1, not %v", test1)
	}
	if test2 != 2 {
		t.Errorf("Expected digitalRoot of 83 to be 2, not %v", test2)
	}

	if test3 != 6 {
		t.Errorf("Expected digitalRoot of 123 to be 6, not %v", test3)
	}

	if test4 != 6 {
		t.Errorf("Expected digitalRoot of 789 to be 6, not %v", test4)
	}

	if test5 != 3 {
		t.Errorf("Expected digitalRoot of 452982 to be 3, not %v", test5)
	}
}

func TestOneAway(t *testing.T) {
	test1 := oneAway("pale", "ple")
	test2 := oneAway("pales", "pale")
	test3 := oneAway("pale", "bale")
	test4 := oneAway("pale", "bake")
	test5 := oneAway("baleen", "baleen")

	if test1 != true {
		t.Errorf("Expected oneAway of tact coa to be true, not %v", test1)
	}

	if test2 != true {
		t.Errorf("Expected oneAway of aaa to be true, not %v", test2)
	}

	if test3 != true {
		t.Errorf("Expected oneAway of abab to be true, not %v", test3)
	}

	if test4 != false {
		t.Errorf("Expected oneAway of cadac to be false, not %v", test4)
	}

	if test5 != true {
		t.Errorf("Expected oneAway of abcde to be true, not %v", test5)
	}
}

func TestisUnique(t *testing.T) {
	test1 := isUnique("blah")
	test2 := isUnique("banana")
	test3 := isUnique("fike hyjrant")
	test4 := isUnique("fire hydrant")
	test5 := isUnique("Another Day")

	if test1 != true {
		t.Errorf("Expected blah to be true because all letters are unique ")
	}

	if test2 != false {
		t.Errorf("Expected banana to be false because all letters are not unique ")
	}

	if test3 != true {
		t.Errorf("Expected fike hyjrant to be true because all letters are unique ")
	}

	if test4 != false {
		t.Errorf("Expected fire hydrant to be false because all letters are not unique ")
	}

	if test5 != false {
		t.Errorf("Expected Another day to be false because all letters are not unique ")
	}
}
