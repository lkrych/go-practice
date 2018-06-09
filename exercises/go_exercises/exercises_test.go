package goExercises

import (
	"testing"
)

func TestCaesarCipher(t *testing.T) {
	test1 := caesarCipher(3, "abc")
	test2 := caesarCipher(3, "abc xyz")
	test3 := caesarCipher(29, "abc xyz")

	if test1 != "def" {
		t.Errorf("Expected caesarCipher with a shift of 3 and input of abc to be def, not %v", test1)
	}

	if test2 != "def abc" {
		t.Errorf("Expected caesarCipher with a shift of 3 and input of abc xyz to be def abc, not %v", test2)
	}

	if test3 != "def abc" {
		t.Errorf("Expected caesarCipher with a shift of 29 and input of abc xyz to be def abc, not %v", test3)
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

func TestNthPrime(t *testing.T) {
	test1 := nthPrime(3)
	test2 := nthPrime(4)
	test3 := nthPrime(15)
	test4 := nthPrime(28)
	test5 := nthPrime(77)

	if test1 != 5 {
		t.Errorf("Expected nthPrime of 3 to be 5, not %v", test1)
	}

	if test2 != 7 {
		t.Errorf("Expected nthPrime of 4 to be 7, not %v", test2)
	}

	if test3 != 47 {
		t.Errorf("Expected nthPrime of 15 to be 47, not %v", test3)
	}

	if test4 != 107 {
		t.Errorf("Expected nthPrime of 28 to be 107, not %v", test4)
	}

	if test5 != 389 {
		t.Errorf("Expected nthPrime of 77 to be 389, not %v", test5)
	}
}

func TestDasherizeNumber(t *testing.T) {
	test1 := dasherizeNumber(203)
	test2 := dasherizeNumber(303)
	test3 := dasherizeNumber(333)
	test4 := dasherizeNumber(3223)

	if test1 != "20-3" {
		t.Errorf("Expected dasherizeNumber of 203  to be 20-3, not %v", test1)
	}

	if test2 != "3-0-3" {
		t.Errorf("Expected dasherizeNumber of 303  to be 3-0-3, not %v", test2)
	}

	if test3 != "3-3-3" {
		t.Errorf("Expected dasherizeNumber of 333 to be 3-3-3, not %v", test3)
	}

	if test4 != "3-22-3" {
		t.Errorf("Expected dasherizeNumber of 3223  to be 3-22-3, not %v", test4)
	}
}
