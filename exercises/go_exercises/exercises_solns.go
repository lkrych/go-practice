package goExercises
 
// func buyStock(arr []int) int {
// 	low := 1000000
// 	diff := 0
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] < low {
// 			low = arr[i]
// 		}
// 		if (arr[i] - low) > diff {
// 			diff = arr[i] - low
// 		}
// 	}
// 	return diff
// }
 
// func binarySearch(arr []int, toFind int) int {
// 	if len(arr) < 1 {
// 		return -1
// 	}
// 	midIdx := len(arr) / 2
// 	if toFind == arr[midIdx] {
// 		return midIdx
// 	} else if toFind > arr[midIdx] {
// 		search := binarySearch(arr[midIdx+1:], toFind)
// 		if search != -1 {
// 			return search + midIdx + 1
// 		}
// 		return search
// 	} else {
// 		return binarySearch(arr[0:midIdx], toFind)
// 	}

// }
 
// func isPrime(n int) bool {
// 	for i := 2; i < n/2+1; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
 
// import "errors"

// func fibonacci(n int) ([]int, error) {
// 	if n < 1 {
// 		return nil, errors.New("n must be greater than 0")
// 	}
// 	fibArr := []int{}
// 	if n < 2 {
// 		fibArr = append(fibArr, n)
// 		return fibArr, nil
// 	}
// 	fibArr = append(fibArr, (fibArr[n-1] + fibArr[n-2]))
// 	return fibArr, nil
// }
 
// //given two numbers represented by linked lists, where each node contains a single digit, add them together
// func sumLists(l1 *Node, l2 *Node) *Node {
// 	carry := 0
// 	sum := 0

// 	head := &Node{
// 		val:  0,
// 		next: nil,
// 	}
// 	currentNode := head

// forLoop:
// 	for l1 != nil || l2 != nil {
// 		intermediate := 0
// 		switch {
// 		case l1 != nil && l2 != nil:
// 			intermediate = (l1.val + l2.val + carry)
// 			sum = intermediate % 10
// 			carry = intermediate / 10
// 			l1 = l1.next
// 			l2 = l2.next

// 		case l1 != nil:
// 			intermediate = (l1.val + carry)
// 			sum = intermediate % 10
// 			carry = intermediate / 10
// 			l1 = l1.next
// 		case l2 != nil:
// 			intermediate = (l2.val + carry)
// 			sum = intermediate % 10
// 			carry = intermediate / 10
// 			l2 = l2.next
// 		default:
// 			intermediate = carry
// 			sum = intermediate % 10
// 			carry = intermediate / 10
// 		}

// 		currentNode.val = sum
// 		nextNode := &Node{
// 			val:  0,
// 			next: nil,
// 		}
// 		//breaking logic
// 		switch {
// 		case l1 == nil && l2 == nil && carry == 0:
// 			break forLoop
// 		case l1 == nil && l2 == nil && carry != 0:
// 			currentNode.next = nextNode
// 			currentNode = nextNode
// 			currentNode.val = carry
// 			break forLoop
// 		}

// 		currentNode.next = nextNode
// 		currentNode = nextNode
// 	}

// 	return head

// }

 
