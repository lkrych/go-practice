package goExercises

import "strings"

//return true if the letter a is within two spaces of the letter z
func nearbyAZ(word string) bool {

}

//return a string where every char in sentence is shifted by shift
func caesarCipher(shift int, sentence string) string {
	shifted := []string{}
	for _, ch := range strings.Split(sentence, "") {
		if ch == " " {
			shifted = append(shifted, ch)
		} else {
			shifted = append(shifted, charShift(ch, shift))
		}
	}
	return strings.Join(shifted, "")
}

const alpha = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func charShift(ch string, shift int) string {
	idx := indexOfCh(ch)
	newIdx := (idx + shift) % 26
	return alpha[newIdx]
}

func indexOfCh(x string) int {
	for idx, ch := range alpha {
		if ch == x {
			return idx
		}
	}
}
