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
	grid := buildGrid(hashedInput)
	fmt.Println(hashedInput)
	fmt.Println("The color is:", color)
	fmt.Println("The grid is ", grid)
}

func readAndHashInput() [16]byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to create an identicon: ")
	text, _ := reader.ReadString('\n')
	return md5.Sum([]byte(text))
}

func buildGrid(hash [16]byte) []byte {
	chunked := [5][]byte{}
	//chunk the bytes
	for i := 0; i < len(hash)-1; i++ {
		chunked[i/3] = append(chunked[i/3], hash[i])
	}

	// mirror the first and zeroth index of the chunks
	// flatten into slice
	flattened := []byte{}
	for idx := range chunked {
		chunked[idx] = append(chunked[idx], chunked[idx][1])
		chunked[idx] = append(chunked[idx], chunked[idx][0])
		flattened = append(flattened, chunked[idx]...)
	}
	return flattened
}
