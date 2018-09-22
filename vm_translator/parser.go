package main

import (
	"bufio"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Command int

const (
	C_NULL       Command = 0 // empty command
	C_ARITHMETIC Command = 1 // pop 2 off stack and add/sub/eq/gt/lt/or/and or pop 1 of stack and neg/not
	C_PUSH       Command = 2 // ex: push local 1, push the value of local[1] onto the stack
	C_POP        Command = 3 // ex: pop local 1, pop the top value of the stack and store it in local[1]
	C_LABEL      Command = 4 // ex: label symbol, this is a label declaration
	C_GOTO       Command = 5 // ex: goto symbol, unconditional branching
	C_IF         Command = 6 // ex: if-goto symbol, conditional branching
	C_FUNCTION   Command = 7 // ex: function functionName nLocals, Function declaration, specifying the number of the function's local variables
	C_CALL       Command = 8 // ex: call functionName nArgs, Function invocation, specifying the number of the function's arguments
	C_RETURN     Command = 9 // ex: return, transfer control back to the calling function
)

//Parser reads in a VM command and breaks it up into its important pieces
type Parser struct {
	scanner *bufio.Scanner
	text    string
	arg1    string
	arg2    int
}

type ParserInterface interface {
	HasMoreCommands() bool
	Advance()
	CommandType() Command
	Arg1() string
	Arg2() int
}

func NewParser(scanner *bufio.Scanner) *Parser {
	return &Parser{scanner, "", "", 0}
}

func (p *Parser) HasMoreCommands() bool {
	return p.scanner.Scan()
	//Scan advances the Scanner to the next token,
	//which will then be available through the Bytes or Text method.
	//It returns false when the scan stops, either by reaching the end of the input or an error.
}

func (p *Parser) Advance() {
	p.text = p.scanner.Text()
	if len(p.text) <= 0 {
		return
	}
	// normalization: remove comments
	tokens := strings.SplitN(p.text, "//", 2)
	if len(tokens) > 0 {
		p.text = tokens[0]
	}
	// normalization: remove spaces and tabs
	p.text = strings.TrimSpace(p.text)
}

func (p *Parser) CommandType() Command {
	p.arg1 = ""
	p.arg2 = 0
	if len(p.text) == 0 {
		// empty token
		return C_NULL
	}
	splitByArgs := strings.Split(p.text, " ")
	commandString := splitByArgs[0]

	if len(splitByArgs) == 1 {
		//arithmetic or return instruction
		arithmeticRegex, err := regexp.Compile("add|sub|eq|gt|lt|or|and|not|neg")
		checkErr(err)
		if commandString == "return" {
			return C_RETURN
		} else if arithmeticRegex.Match([]byte(commandString)) {
			p.arg1 = commandString
			return C_ARITHMETIC
		} else {
			log.Fatal("Unrecognized Command Type:", commandString)
		}
	} else if len(splitByArgs) == 2 {
		//label, goto, or if
		p.arg1 = splitByArgs[1]
		if commandString == "label" {
			return C_LABEL
		} else if commandString == "goto" {
			return C_GOTO
		} else if commandString == "if-goto" {
			return C_IF
		} else {
			log.Fatal("Unrecognized Command Type:", commandString)
		}
	} else if len(splitByArgs) == 3 {
		// push, pop, function, or call
		p.arg1 = splitByArgs[1]
		p.arg2, _ = strconv.Atoi(splitByArgs[2])
		if commandString == "push" {
			return C_PUSH
		} else if commandString == "pop" {
			return C_POP
		} else if commandString == "function" {
			return C_FUNCTION
		} else if commandString == "call" {
			return C_CALL
		} else {
			log.Fatal("Unrecognized Command Type:", commandString)
		}
	}
	return C_NULL
}

func (p *Parser) Arg1() string {
	return p.arg1
}

func (p *Parser) Arg2() int {
	return p.arg2
}
