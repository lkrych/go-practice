package main

import (
	"bufio"
	"fmt"
)

type CodeWriter struct {
	writer *bufio.Writer
}

type CodeWriterInterface interface {
	setFileName()
	WriteArithmetic()
	WritePushPop()
}

func NewCodeWriter(writer *bufio.Writer) *CodeWriter {
	return &CodeWriter{writer}
}

func (c *CodeWriter) setFileName(fileName string) {
	output := "// Assembly code for " + fileName
	fmt.Fprintln(c.writer, output)
}

func (c *CodeWriter) WriteArithmetic(command string) {
	switch command {
	case "add":
	case "sub":
	case "eq":
	case "gt":
	case "lt":
	case "or":
	case "and":
	case "not":
	case "neg":
	}
}

func (c *CodeWriter) WritePushPop(command Command, segment string, index int) {
	output := ""
	switch command {
	case C_POP:
		decrStackPointer := fmt.Sprintf("@%v // %v-- \n M=M-1 \n", index, index)
		loadConstant := fmt.Sprintf("D=M \n")

	case C_PUSH:
		constant := fmt.Sprintf("@%v // D=%v \n", index, index)
		saveToStack := fmt.Sprintf("@%v // *%v = D \n A=M \n M=D \n", segment, segment)
		incrStackPointer := fmt.Sprintf("@%v // %v++ \n M=M+1 \n", segment, segment)
		output = constant + saveToStack + incrStackPointer
	}
	fmt.Fprintln(c.writer, output)
}
