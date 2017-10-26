package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	hashedInput := readAndHashInput()
	color := hashedInput[:3]
	fmt.Println(hashedInput)
	fmt.Println("The color is:", color)
}

func readAndHashInput() [16]byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to create an identicon: ")
	text, _ := reader.ReadString('\n')
	return md5.Sum([]byte(text))
}
