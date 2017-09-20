package main

import (
	"fmt"
)

func main() {
	reversed := reverse("reversed")
	longest := longestWord("This is a long sentence with many lengthy words")
	fmt.Println(reversed)
	fmt.Println("The longest word in the sentence is,", longest)
}
