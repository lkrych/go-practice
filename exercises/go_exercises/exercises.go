package goExercises

import (
	"regexp"
	"strconv"
	"strings"
)

//return true if the letter a is within two spaces of the letter z
func nearbyAZ(word string) bool {
	r, _ := regexp.Compile("z")
	for i, ch := range strings.Split(word, "") {
		if ch == "a" {
			search := false
			if i+4 > len(word)-1 {
				search = r.MatchString(word[i+1:])
			} else {
				search = r.MatchString(word[i+1 : i+4])
			}
			if search {
				return true
			}
		}
	}
	return false
}

//convert base ten int n, into hexadecimal or binary representation
func baseConverter(n int, base int) string {
	num := []string{}
	hexMap := map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
	}
	for n > 0 {

		if base == 2 {
			num = append(num, strconv.Itoa(n%2))
		} else {
			num = append(num, hexMap[n%16])
		}
		n = n / base
	}
	return reverse(strings.Join(num, ""))
}

func reverse(word string) string {
	s := []rune(word)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
