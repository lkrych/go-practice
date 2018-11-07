package goExercises

//Given a string, find the length of the longest substring without repeating characters.
func longestSubString(sentence string) int {
	mIdxs := map[string]int{}
	idx1 := 0
	longestSub := 0
	lenCurrSub := 0
	for idx1 < len(sentence)-1 {
		currChar := string(sentence[idx1])
		if lastIdx, ok := mIdxs[currChar]; !ok {
			lenCurrSub++
		} else {
			lenCurrSub = idx1 - lastIdx
		}
		//keep track of longest sub
		if lenCurrSub > longestSub {
			longestSub = lenCurrSub
		}
		mIdxs[currChar] = idx1
		idx1++
	}
	return longestSub
}

type Node struct {
	val  int
	next *Node
}

//find the kth to last element of a singly-linked list
func findKthElement(head *Node, k int) (*Node, int) {
	if head.next == nil {
		return head, 1
	}

	node, i := findKthElement(head.next, k)
	i++
	if i == k {
		return head, i
	}

	return node, i
}
