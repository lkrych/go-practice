package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./aocday3.txt")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//init matrix
	var matrix [1000][1000]string
	m := map[string]bool{}
	for scanner.Scan() {
		// sample input
		// #1 @ 100,366: 24x27
		line := scanner.Text()
		splitBySpaces := strings.Split(line, " ")
		//break down into requisite parts
		currentSquare := splitBySpaces[0]
		leftandTop := strings.Split(splitBySpaces[2], ",")
		left, err := strconv.Atoi(leftandTop[0])
		checkErr(err)
		//slice off :
		top, err := strconv.Atoi(leftandTop[1][:len(leftandTop[1])-1])
		checkErr(err)
		squareDims := strings.Split(splitBySpaces[3], "x")
		squareX, err := strconv.Atoi(squareDims[0])
		checkErr(err)
		squareY, err := strconv.Atoi(squareDims[1])
		checkErr(err)

		noOverlaps := true
		//step through and add an x to the array for each square, count the number of overlaps
		for i := left + 1; i <= left+squareX; i++ {
			for j := top + 1; j <= top+squareY; j++ {

				//check if already exists in matrix
				if matrix[i][j] != "" {
					noOverlaps = false
					m[matrix[i][j]] = false
				}
				//add to matrix
				matrix[i][j] = currentSquare
			}
		}
		m[currentSquare] = noOverlaps
	}
	for k, v := range m {
		if v == true {
			fmt.Println("The square with no overlaps is", k)
		}
	}
	fmt.Println(m)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("There was an error: %s \n", err)
	}
}
