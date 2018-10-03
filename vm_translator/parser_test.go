package main

import (
	"bufio"
	"strings"
	"testing"
)

type expectParserData struct {
	ctype Command
	text  string
	arg1  string
	arg2  int
}

var vmCode = `
	// comment line
	function SimpleFunction.test 2
	push local 0
	push local 1
	add
	not
	push argument 0
	add
	return
	call blah 2
	label SYMBOL
	goto CYMBAL
	if-goto SYMBOL
	`
var expects = []expectParserData{
	{C_NULL, "", "", 0},
	{C_NULL, "", "", 0},
	{C_FUNCTION, "function SimpleFunction.test 2", "SimpleFunction.test", 2},
	{C_PUSH, "push local 0", "local", 0},
	{C_PUSH, "push local 1", "local", 1},
	{C_ARITHMETIC, "add", "add", 0},
	{C_ARITHMETIC, "not", "not", 0},
	{C_PUSH, "push argument 0", "argument", 0},
	{C_ARITHMETIC, "add", "add", 0},
	{C_RETURN, "return", "", 0},
	{C_CALL, "call blah 2", "blah", 2},
	{C_LABEL, "label SYMBOL", "SYMBOL", 0},
	{C_GOTO, "goto CYMBAL", "CYMBAL", 0},
	{C_IF, "if-goto SYMBOL", "SYMBOL", 0},
	{C_NULL, "", "", 0},
}

func TestParser(t *testing.T) {
	buf := strings.NewReader(vmCode)
	s := bufio.NewScanner(buf)
	p := NewParser(s)

	for i := range expects {
		if !p.HasMoreCommands() {
			t.Errorf("unexpected finish at line: %d", i+1)
		}

		p.Advance()
		expect := &expects[i]
		actualType := p.CommandType()
		if expect.ctype != actualType {
			t.Fatalf("%v: expect %v, but got %v", p.text, expect.ctype, actualType)
		}

		if expect.text != p.text {
			t.Fatalf("text: expect %v, but %v", expect.text, p.text)
		}

		if expect.arg1 != p.arg1 {
			t.Fatalf("arg1: expect %v, but %v", expect.arg1, p.arg1)
		}

		if expect.arg2 != p.arg2 {
			t.Fatalf("arg2: expect %v, but %v", expect.arg2, p.arg2)
		}
	}
}
