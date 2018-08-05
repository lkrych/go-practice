package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
	flags "github.com/jessevdk/go-flags"
)

// Write a program that calculates a postfix arithmetic expression (reverse polish notation)
// Ex: 3 4 5 × − => 3 - (5x4) = -17

var opts struct {
	UserInput string `short:"u" long:"userinput" description:"a string of user input, only integers and the characters +-*/ are valid"`
}

func main() {
	//interpret the CL stdin
	flags.Parse(&opts)
	reversePolish(opts.UserInput)
}

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) insert(n int) {
	//add to front of linked list
	newNode := &Node{
		val:  n,
		next: nil,
	}
	newNode.next = ll.head
	ll.head = newNode
}

func (ll *LinkedList) removeFirstNode() *Node {
	firstNode := ll.head
	ll.head = ll.head.next
	return firstNode
}

type Stack struct {
	data LinkedList
	size int
}

func (s *Stack) push(n int) {
	s.data.insert(n)
	s.size = s.size + 1
}

func (s *Stack) pop() int {
	first := s.data.removeFirstNode()
	s.size = s.size - 1
	return first.val
}

func (s *Stack) isEmpty() bool {
	return s.size == 0
}

func reversePolish(arithmeticExpression string) int {
	//create regexp for matching against operators
	r := regexp.MustCompile("[+-/*]")
	//split input by spaces
	splitInput := strings.Split(arithmeticExpression, " ")
	//init Stack
	s := &Stack{}

	for _, char := range splitInput {
		if r.Match([]byte(char)) {
			int1 := s.pop()
			int2 := s.pop()
			expression := fmt.Sprintf("%v %v %v", int1, char, int2)

			//use the govaluate package to evalute the arithmetic expression
			eval, err := govaluate.NewEvaluableExpression(expression)
			if err != nil {
				fmt.Printf("There was an error building your expression: %v \n", expression)
				log.Fatal("The program is shutting down")
			}

			result, err := eval.Evaluate(nil)
			if err != nil {
				fmt.Printf("There was an error evaluating your expression: %v \n", expression)
				log.Fatal("The program is shutting down")
			}
			fmt.Printf("The expression being evaluated is %v = %v \n", expression, result)

			//govaluates returns an interface so we need to convert from interface to int
			intResult := int(result.(float64))
			s.push(intResult)

		} else {
			//if there isn't an operator, push the number onto the stack
			intVal, err := strconv.Atoi(char)
			if err != nil {
				fmt.Printf("%v is not a valid integer \n", char)
				log.Fatal("Please use valid characters for this program: 0-9 and *-/+")
			}
			s.push(intVal)
		}
	}
	//return the end result
	total := s.pop()

	fmt.Printf("The evaluation of %v yields %v \n", opts.UserInput, total)
	return total
}
