package goExercises

import (
	"testing"
)

func TestNumRepeats(t *testing.T) {
	test1 := numRepeats("abdbc")
	test2 := numRepeats("aaa")
	test3 := numRepeats("abab")
	test4 := numRepeats("cadac")
	test5 := numRepeats("abcde")

	if test1 != 1 {
		t.Errorf("Expected numRepeats of abdbc to be 1, not %v", test1)
	}

	if test2 != 1 {
		t.Errorf("Expected numRepeats of aaa to be 1, not %v", test2)
	}

	if test3 != 2 {
		t.Errorf("Expected numRepeats of abab to be 2, not %v", test3)
	}

	if test4 != 2 {
		t.Errorf("Expected numRepeats of cadac to be 2, not %v", test4)
	}

	if test5 != 0 {
		t.Errorf("Expected numRepeats of abcde to be 0, not %v", test5)
	}
}

func TestOneAway(t *testing.T) {
	test1 := oneAway("pale", "ple")
	test2 := oneAway("pales", "pale")
	test3 := oneAway("pale", "bale")
	test4 := oneAway("pale", "bake")
	test5 := oneAway("baleen", "baleen")

	if test1 != true {
		t.Errorf("Expected oneAway of pale to ple to be true, not %v", test1)
	}

	if test2 != true {
		t.Errorf("Expected oneAway of pales to pale to be true, not %v", test2)
	}

	if test3 != true {
		t.Errorf("Expected oneAway of pale to bale to be true, not %v", test3)
	}

	if test4 != false {
		t.Errorf("Expected oneAway of pale to bake to be false, not %v", test4)
	}

	if test5 != false {
		t.Errorf("Expected oneAway of baleen to baleen to be false, not %v", test5)
	}
}
