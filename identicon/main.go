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
	filtered := filterOddSquares(grid)
	fmt.Println(hashedInput)
	fmt.Println("The color is:", color)
	fmt.Println("The grid is ", grid)
	fmt.Println("The filtered grid is ", filtered)
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
	for i := range chunked {
		chunked[i] = append(chunked[i], chunked[i][1])
		chunked[i] = append(chunked[i], chunked[i][0])
		flattened = append(flattened, chunked[i]...)
	}
	return flattened
}

func filterOddSquares(grid []byte) []byte {
	for i, el := range grid {
		if el%2 != 0 {
			grid[i] = 0
		}
	}
	return grid
}
