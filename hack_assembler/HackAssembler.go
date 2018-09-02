package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Filename string `short:"f" long:"file"  description:"a file to read DFA input from."`
}

func main() {
	//interpret the CL stdin
	flags.Parse(&opts)
	fileName := opts.Filename

	//open file
	fp, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Err! The file cannot open")
		os.Exit(3)
	}
	//set up scanner, this creates an interface for reading a file.
	//this is more efficient than reading all of the file into memory and then
	//converting it (which was my origional idea...)
	scanner := bufio.NewScanner(fp)
	checkErr(err)

	//exit if we aren't given an asm file
	if !strings.HasSuffix(fileName, ".asm") {
		fmt.Println("Usage: Assembler foo.asm")
		os.Exit(2)
	}

	//setup hack file
	hackfile := strings.TrimSuffix(string(fileName), ".asm") + ".hack"
	f, err := os.Create(hackfile)
	checkErr(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	//initialize symbol table and parser
	symbolTable := scanSymbol(NewParser(scanner))

	//main loop
	fp.Seek(0, 0)
	scanner = bufio.NewScanner(fp) //reset
	ramAddr := 0x0010              // starts from 16
	p := NewParser(scanner)
	for p.HasMoreCommands() {
		p.Advance()
		var output string
		switch p.CommandType() {
		case A_COMMAND:
			output = "0"
			symbol := p.Symbol()
			addr, err := strconv.Atoi(symbol)
			if err == nil {
				//number-like string
				output += int2bin(addr)
			} else {
				//symbol
				if symbolTable.Contains(symbol) {
					// known symbol
					addr = symbolTable.GetAddress(symbol)
					output += int2bin(addr)
				} else {
					// new symbol
					symbolTable.AddEntry(symbol, ramAddr)
					output += int2bin(ramAddr)
					ramAddr++
				}
			}
			fmt.Fprintln(w, output)

		case C_COMMAND:
			output = "111"
			comp, err := CodeComp(p.Comp())
			if err != nil {
				os.Exit(5)
			}
			output += comp
			dest, err := CodeDest(p.Dest())
			if err != nil {
				os.Exit(5)
			}
			output += dest
			jump, err := CodeJump(p.Jump())
			if err != nil {
				os.Exit(5)
			}
			output += jump
			fmt.Fprintln(w, output)

		case L_COMMAND:
			// do nothing
		}
		w.Flush()
	}

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// https://davidzych.com/converting-an-int-to-a-binary-string-in-c/
func int2bin(num int) string {
	var bin string
	for i := 1 << 14; i > 0; i = i >> 1 {
		fmt.Printf("%v\n", i)
		if i&num != 0 {
			bin += "1"
		} else {
			bin += "0"
		}
	}
	return bin
}

//initialize symbol table
func scanSymbol(p *Parser) SymbolTable {
	st := NewSymbolTable()
	romAddr := 0

	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType() {
		case A_COMMAND, C_COMMAND:
			romAddr++
		case L_COMMAND:
			st.AddEntry(p.Symbol(), romAddr)
		}
	}

	return st
}
