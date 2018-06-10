import "testing"

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