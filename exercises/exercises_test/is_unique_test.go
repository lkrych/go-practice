import "testing"

func TestisUnique(t *testing.T) {
	test1 := isUnique("blah")
	test2 := isUnique("banana")
	test3 := isUnique("fike hyjrant")
	test4 := isUnique("fire hydrant")
	test5 := isUnique("Another Day")

	if test1 != true {
		t.ErrorF("Expected blah to be true because all letters are unique ")
	}

	if test2 != false {
		t.ErrorF("Expected banana to be false because all letters are not unique ")
	}

	if test3 != true {
		t.ErrorF("Expected fike hyjrant to be true because all letters are unique ")
	}

	if test4 != false {
		t.ErrorF("Expected fire hydrant to be false because all letters are not unique ")
	}

	if test5 != false {
		t.ErrorF("Expected Another day to be false because all letters are not unique ")
	}
}