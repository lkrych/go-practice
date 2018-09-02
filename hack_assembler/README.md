# Hack Assembler

This program translates Hack assembly programs into executable Hack binary code. 

It was created while taking [Nand2Tetris](https://www.nand2tetris.org/), a course on building a modern computer from first principles.

## How to use the Hack Assembler

The source program is supplied in a text file named any_name.asm (we assume this input is error-free) and the generated code is written into a text file named any_name.hack

To initiate the assembler:

``` go HackAssembler any_name.asm ```

## How does the assembler work?

There are three important sub-components of the assembler. 

1. Parser: unpacks each instruction into its underlying fields

2. Code: translates each field into its corresponding binary value

3. Symbol Table: Manages the symbol table

Much of this code was copied or modified from https://github.com/hirak/Assembler. Thanks for putting this out there Hirak. I learned a lot reading your code!