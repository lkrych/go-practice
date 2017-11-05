package main

import (
	"fmt"
)

func main() {
	fmt.Println(CompareNumbers(1, 1))
}

func CompareNumbers(i1, i2 int) (bool, int) {
	if i1 > i2 {
		fmt.Println("the first number is greater than the second number")
		return false, i1 - i2
	} else if i2 > i1 {
		fmt.Println("The second number is greater than the first number")
		return false, i2 - i1
	}
	fmt.Println("The numbers are equal!")
	return true, 0
}
