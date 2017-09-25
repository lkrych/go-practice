package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	square := square{
		sideLength: 4,
	}
	triangle := triangle{
		base:   4,
		height: 8,
	}
	fmt.Println("The area of the square is: ")
	printArea(square)
	fmt.Println("The area of the triangle is: ")
	printArea(triangle)

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
