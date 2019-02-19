package goExercises
 
// //Given a string, find the length of the longest substring without repeating characters.
// func lengthOfLongestSubstring(s string) int {
// 	bs := []byte(s)
// 	var maxLen, start int
// 	idxMap := map[byte]int{}
// 	for i := 0; i < len(bs); i++ {
// 			if _, ok := idxMap[bs[i]]; ok && start <= idxMap[bs[i]] {
// 					start = idxMap[bs[i]] + 1
// 			} else {
// 					if maxLen < i - start + 1 {
// 							maxLen = i - start + 1
// 					}
// 			}
// 			idxMap[bs[i]] = i
// 	}
// 	return maxLen
// }
 
// type stackOfStacks struct {
// 	capacity int
// 	stacks   []stack
// }

// func (s *stackOfStacks) push(item int) {
// 	if len(s.stacks) == 0 || len(s.peek().data) == s.capacity {
// 		newStack := &stack{}
// 		newStack.push(item)
// 		s.stacks = append(s.stacks, *newStack)
// 	} else {
// 		currentStack := s.peek()
// 		currentStack.push(item)
// 	}
// }

// func (s *stackOfStacks) pop() int {
// 	var popped int
// 	currentStack := s.peek()
// 	if len(currentStack.data) > 0 {
// 		popped = currentStack.pop()
// 		if len(currentStack.data) == 0 {
// 			s.stacks = s.stacks[:len(s.stacks)-1] //remove empty stack
// 		}
// 	} else if len(s.stacks) == 0 { //make sure there is a stack
// 		log.Fatal("There are no more stacks!")
// 	} else {
// 		s.stacks = s.stacks[:len(s.stacks)-1] //remove empty stack
// 		popped = s.peek().pop()
// 	}
// 	return popped

// }

// func (s *stackOfStacks) peek() *stack {
// 	return &s.stacks[len(s.stacks)-1]
// }

 
