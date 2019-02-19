package goExercises

import "strings"

//Given a string, find the length of the longest substring without repeating characters.
func longestSubString(sentence string) int {
	longest := 0
	currStartIdx := 0
	charMap := map[string]int{}
	for i, char := range strings.Split(sentence, "") {
		if _, ok := charMap[char]; ok {	
			currStartIdx = i
		}
		charMap[char] = i
		if (i - currStartIdx) > longest {
			longest = (i - currStartIdx)
		}
		
	}
	return longest
}
 
//implement a data structure stackOfStacks (setOfStacks really) that is composed of
//several stacks and should create a new stack once the previous one exceeds capacity
//push and pop should behave identically to a single stack
type stack struct {
	data []int
}

func (s *stack) push (x int) {
	s.data = append([]int{x}, s.data...)
}

func (s *stack) pop () int {
	first := s.data[0]
	s.data = s.data[1:]
	return first
}
	
type stackOfStacks struct {
	stacks []stack
	capacity int
}

func (s *stackOfStacks) push (x int) {
	topStack := s.stacks[0]
	if len(topStack.data) >= s.capacity {
		newStack := stack{
			data: []int{x},
		}
		s.stacks = append( []stack{newStack}, s.stacks...) 
	} else {
		topStack.push(x)
	}
}

func (s *stackOfStacks) pop () int {
	topStack := s.stacks[0]
	popped := topStack.pop()
	if len(topStack.data) == 0 {
		s.stacks = s.stacks[1:]
	}
	return popped
}
 
