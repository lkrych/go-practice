package main

import (
	"bufio"
	"strings"
)

type Command int

const (
	N_COMMAND Command = 0 // empty command
	A_COMMAND Command = 1 // address ex)@123
	C_COMMAND Command = 2 // compute ex)A=M;JGT
	L_COMMAND Command = 3 // label   ex)(LOOP)
)

//Parser reads in a Hack assembly line and breaks it up into its important pieces
type Parser struct {
	scanner *bufio.Scanner
	text    string
	dest    string
	comp    string
	jump    string
}

type ParserInterface interface {
	HasMoreCommands() bool
	Advance()
	CommandType() Command
	Symbol() string
	Dest() string
	Comp() string
	Jump() string
}

func NewParser(scanner *bufio.Scanner) *Parser {
	return &Parser{scanner, "", "", "", ""}
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
	p.dest = ""
	p.comp = ""
	p.jump = ""
	if len(p.text) == 0 {
		// empty token
		return N_COMMAND
	}
	if p.text[0] == '@' {
		// @Xxx
		return A_COMMAND
	}
	if p.text[0] == '(' {
		// (LABEL)
		return L_COMMAND
	}
	// dest=comp;jump
	p.analyze()
	return C_COMMAND
}

// parse "dest=comp;jump"
func (p *Parser) analyze() {
	// "dest=comp", "jump"
	tokens := strings.SplitN(p.text, ";", 2)
	destcomp := tokens[0]
	if len(tokens) == 2 {
		p.jump = tokens[1]
	} else {
		p.jump = ""
	}

	// "dest", "comp"
	tokens = strings.SplitN(destcomp, "=", 2)
	if len(tokens) == 2 {
		p.dest = tokens[0]
		p.comp = tokens[1]
	} else {
		p.dest = ""
		p.comp = destcomp
	}
}

func (p *Parser) Symbol() string {
	if strings.HasPrefix(p.text, "@") {
		return strings.TrimPrefix(p.text, "@")
	} else {
		return strings.Trim(p.text, "()")
	}
}

func (p *Parser) Dest() string {
	return p.dest
}

func (p *Parser) Comp() string {
	return p.comp
}

func (p *Parser) Jump() string {
	return p.jump
}
