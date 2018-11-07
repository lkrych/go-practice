package goExercises
 
// //Given a string, find the length of the longest substring without repeating characters.
// func lengthOfLongestSubstring(s string) int {
// 	bs := []byte(s)
// 	var maxLen, start int
// 	idxMap := map[byte]int{}
// 	for i := 0; i < len(bs); i++ {
// 			if _, ok := idxMap[bs[i]]; ok && start <= idxMap[bs[i]] {
// 					start = idxMap[bs[i]] + 1
// 			} else {
// 					if maxLen < i - start + 1 {
// 							maxLen = i - start + 1
// 					}
// 			}
// 			idxMap[bs[i]] = i
// 	}
// 	return maxLen
// }
 
//there are two ways to solve this problem: a recursive way and an iterative way
// the iterative way is more efficient, the recursive way is easier to read
//
// recursive:
//
// func findKthElement(head *Node, k int) (*Node, int) {
// 	if head.next == nil {
// 		return head, 1
// 	}
// 	var node, ith = findKthElement(head.next, k)
// 	ith++
// 	if ith == k {
// 		return head, ith
// 	}
// 	return node, ith
// }
//
// iterative:
//
// func findKthElement(head *Node, k int) *Node {
// 	p1 := head
// 	p2 := head

// 	//move p1 k nodes into the list
// 	for i := 0; i < k; i++ {
// 		if p1 == nil {
// 			return nil
// 		}
// 		p1 = p1.next
// 	}

// 	//move them at the same pace. when p1 hits the end, p2 will be at the right ele
// 	for p1 != nil {
// 		p1 = p1.next
// 		p2 = p2.next
// 	}
// 	return p2
// }
 
