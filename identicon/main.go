package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	hashedInput := readAndHashInput()
	c := hashedInput[:3]
	colorDraw := color.RGBA{c[0], c[1], c[2], 255}
	grid := buildGrid(hashedInput)
	filtered := filterOddSquares(grid)

	newPNGfile := "./identicon.png"

	myimage := image.NewRGBA(image.Rect(0, 0, 250, 250))
	white := color.RGBA{255, 255, 255, 255}

	// backfill entire surface with white
	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	for idx, el := range filtered {
		if el == 0 {
			continue
		}
		topX := (idx % 5) * 50
		botX := topX + 50
		topY := (idx / 5) * 50
		botY := topY + 50

		rect := image.Rect(topX, topY, botX, botY)
		draw.Draw(myimage, rect, &image.Uniform{colorDraw}, image.ZP, draw.Src)
	}

	// ... save image

	myfile, _ := os.Create(newPNGfile)
	png.Encode(myfile, myimage)
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
