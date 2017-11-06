package main

import "fmt"

func main() {
	// a := [5]int{1, 2, 3, 4, 5}
	// myslice := []int{1, 2, 3, 4, 5}
	// myslice = append(myslice, 6, 7, 8, 9, 10)
	// fmt.Println(a)
	// fmt.Println(myslice)

	s := make([]int, 5)                          //make a slice of length and capacity 5
	s[0], s[1], s[2], s[3], s[4] = 1, 2, 3, 4, 5 //capacity: cap(s), length, len(s)
	fmt.Printf("%v, capacity: %v, length: %v \n", s, cap(s), len(s))
	s1 := s[1:3]
	// Slices still have access to full capacity of array.
	fmt.Printf("%v, capacity: %v, length: %v \n", s1, cap(s1), len(s1))
	fmt.Printf("Surprised by this behavior? Slices are just pointers to an index, with a length and capacity. %v \n", s1[:cap(s1)])

	//to create a slice with only two spots of allocated memory
	s2 := make([]int, 2)
	copy(s2, s[1:3])
	fmt.Printf("%v, capacity: %v, length: %v \n", s2, cap(s2), len(s2))

}
