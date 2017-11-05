package main

import (
	"fmt"
)

func main() {
	fmt.Println(compareNumbers(1, 1))
	x := 2
	switch x {
	case 3:
		fmt.Println("I am 3")
	case 2:
		fmt.Println("I am 2")
		//fallthrough // evaluates the next statement/s as true!
	case 4:
		fmt.Println("I am 4")
	}
}

func compareNumbers(i1, i2 int) (bool, int) {
	/*
		if i1 > i2 {
			fmt.Println("the first number is greater than the second number")
			return false, i1 - i2
		} else if i2 > i1 {
			fmt.Println("The second number is greater than the first number")
			return false, i2 - i1
		}
		fmt.Println("The numbers are equal!")
		return true, 0
	*/

	//alternative syntax for conditionals
	switch {
	case i1 > i2:
		fmt.Println("the first number is greater than the second number")
		return false, i1 - i2
	case i2 > i1:
		fmt.Println("The second number is greater than the first number")
		return false, i2 - i1
	}
	fmt.Println("The numbers are equal!")
	return true, 0

}
