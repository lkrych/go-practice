package main

import "fmt"

type Percolation struct {
	grid        []bool
	objects     []int
	objectSize  []int
	openSquares int
	gridCount   int
}

func initialize(n int) *Percolation {
	grid := []bool{}
	objects := []int{}
	objectSize := []int{}

	for i := 0; i < n*n; i++ {
		grid[i] = false
		objects[i] = i
		objectSize[i] = 1
	}
	return &Percolation{
		grid:        grid,
		objects:     objects,
		objectSize:  objectSize,
		openSquares: n * n,
		gridCount:   n,
	}
}

//translate from row x column idx to single dimension
func translate2dToFlat(row int, col int, gridCount int) int {
	return (row * gridCount) + col
}

func (p *Percolation) Open(row int, col int) {
	if !p.isOpen(row, col) {
		//open square, check if it needs to connect to any other squares
		oneD := translate2dToFlat(row, col, p.gridCount)
		p.grid[oneD] = true
		p.UnionWithNeighbors(row, col)
		p.openSquares++
	}
}

func (p *Percolation) isOpen(row int, col int) bool {
	oneD := translate2dToFlat(row, col, p.gridCount)
	return p.grid[oneD]
}

func (p *Percolation) numberOfOpensites() int {
	return p.openSquares
}

func (p *Percolation) percolates bool {
	
	for i := 0; i < p.gridCount; i++ {

	}

	//Check if they are connected
}

func (p *Percolation) UnionWithNeighbors(row int, col int) {
	oneD := translate2dToFlat(row, col, p.gridCount)
	if (row - 1) > 0 && p.isOpen(row -1, col) {
		p.Union(oneD, translate2dToFlat(row-1, col, p.gridCount))
	}

	if (row + 1) < p.gridCount && p.isOpen(row +1, col) {
		p.Union(oneD, translate2dToFlat(row+1, col, p.gridCount))
	}

	if (col - 1) > 0 && p.isOpen(row, col -1) {
		p.Union(oneD, translate2dToFlat(row, col-1, p.gridCount))
	}

	if (col + 1) < p.gridCount && p.isOpen(row, col +1) {
		p.Union(oneD, translate2dToFlat(row, col+1, p.gridCount))
	}
}

func (u *Percolation) Union(p int, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)

	if rootP == rootQ {
		return
	}

	if u.objectSize[rootP] < u.objectSize[rootQ] {
		u.objects[rootP] = rootQ
		u.objectSize[rootQ] += u.objectSize[rootP]
	} else {
		u.objects[rootQ] = rootP
		u.objectSize[rootP] += u.objectSize[rootQ]
	}
}

func (u *Percolation) Find(p int) int {
	for p != u.objects[p] {
		p = u.objects[p]
	}
	return p
}

func (u *Percolation) Connected(p int, q int) bool {
	return u.Find(p) == u.Find(q)
}

func main() {
	p := initialize(5)
	fmt.Println(p)
}
