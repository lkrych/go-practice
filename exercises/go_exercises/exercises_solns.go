package goExercises

// type stack struct {
// 	data []int
// }

// func (s *stack) pop() int {
// 	popped := s.data[len(s.data)-1]
// 	s.data = s.data[:len(s.data)-1]
// 	return popped
// }

// func (s *stack) push(item int) {
// 	s.data = append(s.data, item)
// }

// func (s *stack) peek() int {
// 	return s.data[len(s.data)-1]
// }

// func (s *stack) isEmpty() bool {
// 	return len(s.data) == 0
// }

// func partitionLL(head *Node, partition int) *Node {
// 	beforeStart := &Node{}
// 	beforeEnd := &Node{}
// 	afterStart := &Node{}
// 	afterEnd := &Node{}

// 	//partition
// 	currentNode := head
// 	for currentNode != nil {
// 		next := currentNode.next
// 		currentNode.next = nil
// 		if currentNode.val < partition {
// 			//insert node into end of before list
// 			if beforeStart.val == 0 {
// 				beforeStart = currentNode
// 				beforeEnd = beforeStart
// 			} else {
// 				beforeEnd.next = currentNode
// 				beforeEnd = currentNode
// 			}
// 		} else {
// 			//insert node into end of after list
// 			if afterStart.val == 0 {
// 				afterStart = currentNode
// 				afterEnd = afterStart
// 			} else {
// 				afterEnd.next = currentNode
// 				afterEnd = currentNode
// 			}
// 		}
// 		currentNode = next
// 	}
// 	if beforeStart.val == 0 {
// 		return afterStart
// 	}

// 	//merge the two lists
// 	beforeEnd.next = afterStart
// 	return beforeStart
// }

// func deleteDupsLL(head *Node) *Node {
// 	intMap := map[int]bool{}
// 	prev := &Node{}
// 	currentNode := head
// 	for currentNode.next != nil {
// 		if _, ok := intMap[currentNode.val]; ok {
// 			prev.next = currentNode.next
// 		} else {
// 			intMap[currentNode.val] = true
// 			prev = currentNode
// 		}
// 		currentNode = currentNode.next
// 	}
// 	if _, ok := intMap[currentNode.val]; ok {
// 		prev.next = currentNode.next
// 	}
// 	return head
// }
