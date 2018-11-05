package goExercises

import "strings"

//return the number of repeated characters in word
func numRepeats(word string) int {
	m := make(map[string]int)
	count := 0
	for _, ch := range strings.Split(word, "") {
		if val, ok := m[ch]; ok {
			m[ch] = val + 1
		} else {
			m[ch] = 1
		}
	}

	for _, v := range m {
		if v > 1 {
			count++
		}
	}

	return count

}

//there are three types of edits that can be performed on strings
// 1) insert a character
// 2) remove a character
// 3) replace a character
// given two strings, write a function to check if they are one edit(or zero edits) away
//ex: pale, ple -> true
//	  pales, pale -> true
//    pale, bale -> true
//    pale, bake -> false
func oneAway(s1, s2 string) bool {
	sameLength := len(s1) == len(s2)
	foundDiff := false
	if sameLength {
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if foundDiff {
					//there are more than two edit differences in a same length string
					return false
				}
				foundDiff = true
			}
		}
		return foundDiff
	} else {
		if len(s1) > len(s2) {
			for i := 0; i < len(s2); i++ {
				if foundDiff {
					if s1[i] != s2[i-1] {
						return false
					}
				}
				if s1[i] != s2[i] {
					//found deleted char in s2
					foundDiff = true
				}
			}
			//s1 > s2, if foundDiff not found than s1 has an extra char
			return true
		} else {
			for i := 0; i < len(s1); i++ {
				if foundDiff {
					if s2[i] != s1[i-1] {
						return false
					}
				}
				if s1[i] != s2[i] {
					//found deleted char in s1
					foundDiff = true
				}
			}
			//same case as above
			return true
		}
	}
	return foundDiff
}
