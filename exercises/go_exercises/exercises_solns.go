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
 
// import "math"

// //func keep track of how far you have traveled, if you can reach the end, you made it!
// func advancingArray(arr []int) bool {
// 	farthestTraveled := arr[0]
// 	for i := 0; i <= farthestTraveled; i++ {
// 		if farthestTraveled == len(arr)-1 {
// 			return true
// 		}
// 		farthestTraveled = int(math.Max(float64(farthestTraveled), float64(arr[i]+i)))
// 	}
// 	return false
// }
 
