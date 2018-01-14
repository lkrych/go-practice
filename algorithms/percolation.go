package main

type Percolation struct {
	grid       [][]bool
	objects    [][]int
	objectSize [][]int
	openSquares int
}

func initialize(n int) *Percolation {
	grid := [][]bool{}
	objects := [][]int{}
	objectSize := [][]int{}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = false
			objects[i][j] = i + j
			objectSize[i][j] = 1
		}
	}
	return &Percolation{
		grid:       grid,
		objects:    objects,
		objectSize: objectSize,
		openSquares: n * n,
	}
}

func (p *Percolation) Open(row int, col int) {
	if !p.isOpen(row, col) {
		//open square, check if it needs to connect to any other squares
		p.grid[row][col] = true
		p.UnionWithNeighbors(row, col)
		p.openSquares++
	}
}

func (p *Percolation) isOpen(row int, col int) bool {
	return p.grid[row][col]
}

func (p *Percolation) numberOfOpensites() int {
	return p.openSquares
}

func (p *Percolation) percolates bool {
	//connect phantomCell with top row
	for i := 0; i < len(p[0]); i++ {
		
	}
	//connect phantomCell with bottom row

	//Check if they are connected
}

func (p *Percolation) UnionWithNeighbors(row int, col int) {

}