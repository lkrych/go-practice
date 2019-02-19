package goExercises
 
import (
"fmt"
"testing"
)


func TestLongestSubString(t *testing.T) {
	test1 := longestSubString("abcabcbb")
	test2 := longestSubString("bbbbb")
	test3 := longestSubString("pwwkew")

	if test1 != 3 {
		t.Errorf("Expected longestSubString of abcabcbb to be 3, not %v", test1)
	}

	if test2 != 1 {
		t.Errorf("Expected longestSubString of abba to be 1, not %v", test2)
	}

	if test3 != 3 {
		t.Errorf("Expected longestSubString of pwwkew to be 3, not %v", test3)
	}
}
 

func TestCreatingAStackOfStacks(t *testing.T) {
	s := &stackOfStacks{
		capacity: 2,
	}
	for i := 1; i < 8; i++ {
		s.push(i)
	}

	if len(s.stacks) != 4 {
		t.Errorf("Expected stack of stacks to have: 4 stacks not %v stacks", len(s.stacks))
		fmt.Printf("stacks: %v \n", s.stacks)
	}

	testPop := s.pop()

	if testPop != 7 {
		fmt.Println(testPop)
		t.Errorf("Expected stackOfStacks pop to work like stack pop")
	}

	if len(s.stacks) != 3 {
		t.Errorf("Expected stack of stacks to have: 3 stacks not %v stacks", len(s.stacks))
		fmt.Printf("stacks: %v \n", s.stacks)
	}

	s.push(testPop)
	if len(s.stacks) != 4 {
		t.Errorf("Expected stack of stacks to have: 4 stacks not %v stacks", len(s.stacks))
		fmt.Printf("stacks: %v \n", s.stacks)
	}

	s.push(0)
	s.push(99)
	if len(s.stacks) != 5 {
		t.Errorf("Expected stack of stacks to have: 5 stacks not %v stacks", len(s.stacks))
		fmt.Printf("stacks: %v \n", s.stacks)
	}

}
 
