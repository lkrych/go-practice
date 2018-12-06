package main

// func main() {
// 	file, err := os.Open("./aocday3.txt")
// 	checkErr(err)
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	//init matrix
// 	var matrix [1000][1000]int
// 	overlaps := 0
// 	for scanner.Scan() {
// 		// sample input
// 		// #1 @ 100,366: 24x27
// 		line := scanner.Text()
// 		splitBySpaces := strings.Split(line, " ")
// 		//break down into requisite parts
// 		currentSquare := splitBySpaces[0]
// 		leftandTop := strings.Split(splitBySpaces[2], ",")
// 		left, err := strconv.Atoi(leftandTop[0])
// 		checkErr(err)
// 		//slice off :
// 		top, err := strconv.Atoi(leftandTop[1][:len(leftandTop[1])-1])
// 		checkErr(err)
// 		squareDims := strings.Split(splitBySpaces[3], "x")
// 		squareX, err := strconv.Atoi(squareDims[0])
// 		checkErr(err)
// 		squareY, err := strconv.Atoi(squareDims[1])
// 		checkErr(err)

// 		fmt.Println("Current Square", currentSquare)
// 		fmt.Println("Left and Top", left, top)
// 		fmt.Println("X and Y", squareX, squareY)
// 		fmt.Println()

// 		//step through and add an x to the array for each square, count the number of overlaps
// 		for i := left + 1; i <= left+squareX; i++ {
// 			for j := top + 1; j <= top+squareY; j++ {
// 				//add to matrix
// 				matrix[i][j] = matrix[i][j] + 1

// 				//check if already exists in matrix
// 				if matrix[i][j] > 1 && matrix[i][j] < 3 {
// 					overlaps++
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println("The number of overlapping square inches is", overlaps)
// }

// func checkErr(err error) {
// 	if err != nil {
// 		log.Fatalf("There was an error: %s \n", err)
// 	}
// }
