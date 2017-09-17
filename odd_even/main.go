package main

import "fmt"

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range ints {
		checkIfEvenOrOdd(val)
	}
}

func checkIfEvenOrOdd(n int) {
	if n%2 == 0 {
		fmt.Printf("%d is even \n", n)
	} else {
		fmt.Printf("%d is odd \n", n)
	}

}
