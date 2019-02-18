package goExercises
 
// func removeDups(arr []int) []int {
// 	writeIdx := 1
// 	for i := 1; i < len(arr); i++ {
// 		if arr[writeIdx-1] != arr[i] {
// 			arr[writeIdx] = arr[i]
// 			writeIdx++
// 		}
// 	}
// 	return arr[:writeIdx]
// }
 
// import "log"

// //copy the info from the next node into the current node and delete the next node
// func deleteMiddleNode(n *Node) {
// 	if n == nil || n.next == nil {
// 		log.Fatal("this is not a middle node!")
// 	}
// 	n.val = n.next.val
// 	n.next = n.next.next
// }
 
