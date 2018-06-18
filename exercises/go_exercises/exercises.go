package goExercises

import "math"

//if the sum of digits in n is less than ten, return that sum, otherwise try to find the digital root of that sum
func digitalRoot(n int) int {
	sum := n/10 + n%10
	if sum < 10 {
		return sum
	}
	return digitalRoot(sum)

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
	if math.Abs(float64(len(s1)-len(s2))) > 1 {
		return false
	}

	shorter := ""
	longer := ""

	if len(s1) < len(s2) {
		shorter = s1
		longer = s2
	} else {
		shorter = s2
		longer = s1
	}

	idx1 := 0
	idx2 := 0
	foundDiff := false

	for idx2 < len(longer) && idx1 < len(shorter) {
		if shorter[idx1] != longer[idx2] {
			if foundDiff {
				return false
			}
			foundDiff = true

			if len(shorter) == len(longer) {
				idx1++ //make sure you are moving indices together unless there is replacement
			}
		} else {
			idx1++
		}
		idx2++
	}
	return true
}

//Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
func isUnique(word string) bool {
	charMap := map[byte]bool{}
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if _, ok := charMap[ch]; ok {
			return false
		} else {
			charMap[ch] = true
		}
	}
	return true
}
