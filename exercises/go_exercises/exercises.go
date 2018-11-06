package goExercises

import (
	"strings"
)

//return a string where every char in sentence is shifted by shift
func caesarCipher(shift int, sentence string) string {
	encryptedString := []string{}
	for _, ch := range strings.Split(sentence, "") {
		if ch == " " {
			encryptedString = append(encryptedString, " ")
			continue
		}
		encryptedString = append(encryptedString, shiftChar(shift, ch))
	}
	return strings.Join(encryptedString, "")
}

func shiftChar(shift int, character string) string {
	alpha := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	intVal := indexOf(character, alpha)
	newVal := alpha[(intVal+shift)%26]
	return newVal
}

func indexOf(item string, arr []string) int {
	for i, v := range arr {
		if v == item {
			return i
		}
	}
	return -1
}

//check to see if a passphrase is valid, it must not contain any duplicate words
func passPhrase(passphrase string) bool {
	m := map[string]bool{}
	for _, word := range strings.Split(passphrase, " ") {
		if _, ok := m[word]; ok {
			return false
		}
		m[word] = true
	}
	return true
}

func passPhraseAdvent() int {
	//read from ../practice/passphrase_input.txt
	//each line has a different passphrase, return how many are valid
	return 1
}

//bonus, check to see if a passphrase is valid, it must not contain any words that are anagrams of eachother
// abcde == dcabe
// func passPhraseAnagram(passphrase string) bool {

// }
