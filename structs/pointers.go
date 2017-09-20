package main

func (pointerToPerson *person) updateNamePointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
