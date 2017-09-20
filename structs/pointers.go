package main

func (pointerToPerson *person) updateNamePointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	//the star function asks for the value at the memory address, this will be the person struct
}
