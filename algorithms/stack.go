package main

//stacks are a LIFO data structure
//they can be implemented with linked lists
//and arrays

type Node struct {
	item string
	next *Node
}

type StackLinkedList struct {
	first *Node
}

func (s *StackLinkedList) pop() string {
	node := s.first
	s.first = item.next
	return node.item
}

func (s *StackLinkedList) push(s string) {
	newNode := &Node{item: s, next: s.first}
	s.first = newNode
}

type StackArray struct {
	capacity int
	stack []string{}
}

func (s *StackLinkedList) pop() *Node {
	item := s.stack[s.capacity]
	s.capacity--
	return item
}

func (s *StackLinkedList) push(s string) {
	s.capacity++
	s.stack[s.capacity] = s
}
