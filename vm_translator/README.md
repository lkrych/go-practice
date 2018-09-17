## VM Translator

This program translates a stack-oriented VM code into Hack Assembly code. It was built in conjunction with Nand2Tetris, an online course. 

The code is broken up into three segments

1. Parser: parses each VM command into its elements

2. CodeWriter: writes the assembly code that implements the parsed command

3. VMTranslator: the main file