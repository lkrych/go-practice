package goExercises

// func incrementArb(arr []int) []int {
// 	lastIdx := len(arr) - 1
// 	arr[lastIdx] = (arr[lastIdx] + 1) % 10
// 	if arr[lastIdx] == 0 {
// 		arr = append(incrementArb(arr[:lastIdx]), arr[lastIdx])
// 	}
// 	return arr
// }

// func factorial(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	return n * factorial(n-1)
// }

// import "strconv"

// func timeConversion(timeInMinutes int) string {
// 	minutes := timeInMinutes % 60
// 	hours := timeInMinutes / 60
// 	var time string
// 	if minutes < 10 {
// 		time = strconv.Itoa(hours) + ":0" + strconv.Itoa(minutes)
// 	} else {
// 		time = strconv.Itoa(hours) + ":" + strconv.Itoa(minutes)
// 	}
// 	return time
// }

//before building the string, you should check to make sure that you need to build the string
// func stringCompression(in string) out string {
// 	if countCompressionLength(in) > len(in) {
// 		return in
// 	}

// 	compressed := []rune{}
// 	countConsec := 0
// 	for i := 0; i < len(in); i++ {
// 		countConsec++

// 		if (i+1 >= len(string) || in[i] != in[i+1]) {
// 			compressed = append(compressed, in[i], countConsec)
// 			countConsec = 0
// 		}
// 	}
// 	return string(compressed)
// }

// func countCompressionLength(in string) bool {
// 	compressedLength := 0
// 	countConsec := 0
// 	for i := 0; i < len(in); i++ {
// 		countConsec++

// 		if (i+1 >= len(string) || in[i] != in[i+1]) {
// 			compressedLength += (1 + len(string(countConsec)))
// 			countConsec = 0
// 		}
// 	}
// }
