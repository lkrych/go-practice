package goExercises
 
// func digitalRoot(n int) int {
// 	if n < 10 {
// 		return n
// 	}
// 	return digitalRoot(n%10 + n/10)
// }
 
// import "math"

// //the key insight to this one is noticing that these edits are very similar and that all
// //we need to do is check to see if the characters are same
// func oneAway(first, second string) bool {
// 	//check lengths
// 	if math.Abs(len(first)-len(second)) > 1 {
// 		return false
// 	}

// 	//get shorter or longer string
// 	s1 = len(first) < len(second) ? first : second
// 	s2 = len(first) < len(second) ? second : first

// 	idx1 := 0
// 	idx2 := 0
// 	foundDiff := false
// 	for idx2 < len(s2) && idx1 < len(s1) {
// 		if s1[idx1] != s2[idx2] {
// 			if foundDiff { //there should only be one edit
// 				return false
// 			}
// 			foundDiff = true

// 			if len(s1) == len(s2) {
// 				idx1++ //on replace, move shorter idx
// 			}
// 		} else {
// 			idx1++ //if matching, move shorter idx
// 		}
// 		idx2++ //always move idx for longer
// 	}
// 	return true
// }
 
// //Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
// func isUnique(word string) bool {
// 	//use a map
// }
 
