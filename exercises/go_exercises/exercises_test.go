package goExercises

import (
	"testing"
)

func TestNearbyAZ(t *testing.T) {
	test1 := nearbyAZ("baz")
	test2 := nearbyAZ("abz")
	test3 := nearbyAZ("abcz")
	test4 := nearbyAZ("a")
	test5 := nearbyAZ("z")
	test6 := nearbyAZ("za")

	if test1 != true {
		t.Errorf("Expected nearbyAZ of baz to be true, not %v", test1)
	}

	if test2 != true {
		t.Errorf("Expected nearbyAZ of abz to be true, not %v", test2)
	}

	if test3 != true {
		t.Errorf("Expected nearbyAZ of abcz to be true, not %v", test3)
	}

	if test4 != false {
		t.Errorf("Expected nearbyAZ of a to be false, not %v", test4)
	}

	if test5 != false {
		t.Errorf("Expected nearbyAZ of z to be false, not %v", test5)
	}

	if test6 != false {
		t.Errorf("Expected nearbyAZ of za to be false, not %v", test6)
	}
}

func TestBaseConverter(t *testing.T) {
	test1 := baseConverter(5, 2)
	test2 := baseConverter(1239449, 16)
	test3 := baseConverter(112, 2)
	test4 := baseConverter(112, 16)

	if test1 != "101" {
		t.Errorf("Expected baseConverter of 5 to base 2 to equal 101, not %v", test1)
	}
	if test2 != "12e999" {
		t.Errorf("Expected baseConverter of 1239449 to base 16 to be 12e999, not %v", test2)
	}

	if test3 != "1110000" {
		t.Errorf("Expected baseConverter of 112 to base 2 to be 1110000, not %v", test3)
	}

	if test4 != "70" {
		t.Errorf("Expected baseConverter of 112 to base 16 to be 70, not %v", test4)
	}
}
