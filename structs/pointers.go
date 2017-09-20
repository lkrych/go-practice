package main

func (pointerToPerson *person) updateNamePointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	//the star operator asks for the value at the memory address, this will be the person struct
	//the star in front of the type is a type description
}
