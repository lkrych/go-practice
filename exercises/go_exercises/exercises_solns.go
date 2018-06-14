package goExercises
 
func deleteDupsLL(head *Node) *Node {
	intMap := map[int]bool{}
	prev := &Node{}
	currentNode := head
	for currentNode.next != nil {
		if _, ok := intMap[currentNode.val]; ok {
			prev.next = currentNode.next
		} else {
			intMap[currentNode.val] = true
			prev = currentNode
		}
		currentNode = currentNode.next
	}
	if _, ok := intMap[currentNode.val]; ok {
		prev.next = currentNode.next
	}
	return head
}
 
// import "strings"

// func scrambleString(word string, idxArray []int) string {
// 	split := strings.Split(word, "")
// 	scrambled := []string{}
// 	for _, el := range idxArray {
// 		scrambled = append(scrambled, split[el])
// 	}
// 	return strings.Join(scrambled, "")
// }
 
// func partitionEvenOdd(arr []int) []int {
// 	partition := 0
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i]%2 == 0 {
// 			arr[i], arr[partition] = arr[partition], arr[i]
// 			partition++
// 		}
// 	}
// 	front := quickSort(arr[:partition])
// 	back := quickSort(arr[partition:])
// 	return append(front, back...)
// }

// func quickSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	pivot := arr[0]
// 	left := []int{}
// 	right := []int{}

// 	for _, val := range arr[1:] {
// 		if val >= pivot {
// 			right = append(right, val)
// 		} else {
// 			left = append(left, val)
// 		}
// 	}
// 	answer := []int{}
// 	sortedLeft := quickSort(left)
// 	sortedRight := quickSort(right)
// 	answer = append(answer, sortedLeft...)
// 	answer = append(answer, pivot)
// 	answer = append(answer, sortedRight...)
// 	return answer
// }
 
// import "math"

// func multiplyArb(arr1 []int, arr2 []int) []int {
// 	sign := 1
// 	//check if numbers are negative
// 	if (arr1[0] * arr2[0]) > 0 {
// 		sign = 1
// 	} else {
// 		sign = -1
// 	}
// 	//set to absolute val
// 	arr1[0] = int(math.Abs(float64(arr1[0])))
// 	arr2[0] = int(math.Abs(float64(arr2[0])))

// 	total := make([]int, len(arr1)+len(arr2))
// 	for i := len(arr1) - 1; i >= 0; i-- {
// 		for j := len(arr2) - 1; j >= 0; j-- {
// 			total[i+j+1] += arr1[i] * arr2[j]
// 			total[i+j] += total[i+j+1] / 10
// 			total[i+j+1] %= 10
// 		}
// 	}
// 	zeroIdx := 0
// 	for {
// 		if total[zeroIdx] != 0 {
// 			break
// 		}
// 		zeroIdx++
// 	}
// 	total = total[zeroIdx:]
// 	total[0] = total[0] * sign

// 	return total
// }
 
// import "math"

// func isPowerOfTwo(n float64) bool {
// 	power := float64(0)
// 	value := math.Exp2(power)
// 	for value <= n {
// 		if value == n {
// 			return true
// 		}
// 		power++
// 		value = math.Exp2(power)
// 	}
// 	return false
// }
 
