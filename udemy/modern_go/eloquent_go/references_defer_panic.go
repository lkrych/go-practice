package main

import "fmt"

func main() {
	//pointers
	var pI *int //memory address of a value of type int
	var I int = 3
	pI = &I //pI now points at the memory address of var I
	fmt.Printf("before incrementing: %v \n", *pI)
	increment(pI)
	fmt.Printf("after incrementing: %v \n", *pI)

	//defers

	// defer fmt.Println("World")     //blocks until end of function, and then executes
	// defer fmt.Println("Great Big") //Multiple defers are stored on a stack
	// fmt.Println("Hello")

	//panics
	fmt.Println("Hello")
	testpanics()
	fmt.Println("World")
}

func increment(pI *int) {
	*pI++ //dereferencing
}

func testpanics() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("We caught the panic and recovered the code.")
			fmt.Printf("The panic text is %v \n", err)
		}
	}()
	panic("A panic happened")
}
