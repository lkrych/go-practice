package goExercises
 
// func isPrime(n int) bool {
// 	for i := 2; i < n/2+1; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func nthPrime(nth int) int {
// 	primeCount := 0
// 	n := 2
// 	for primeCount < nth {
// 		if isPrime(n) {
// 			primeCount++
// 		}
// 		n++
// 	}
// 	return n - 1
// }
 
// import "strings"

// func scrambleString(word string, idxArray []int) string {
// 	split := strings.Split(word, "")
// 	scrambled := []string{}
// 	for _, el := range idxArray {
// 		scrambled = append(scrambled, split[el])
// 	}
// 	return strings.Join(scrambled, "")
// }
 
