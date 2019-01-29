package goExercises
 
//return sorted arr using mergesort
func mergeSort(arr []int) []int {
	if (len(arr) <= 1) {
		return arr
	}

	middle := len(arr)/2
	arr1 := mergeSort(arr[:middle])
	arr2 := mergeSort(arr[middle+1:]

	
	return merge(arr1, arr2)
}


func merge(arr1, arr2 queue), []int {
	merged := []int{}
	for len(arr1
	return merged
}

type queue struct {
	data []int{}
}

func (q *queue) peek() {
	if (len(q.data) > 1) {
		return len(q.data)
	} else {
		return null
	}
}

func (q *queue) push (el int) {
	q.data = append(q.data, el)
}

func (q *queue) pop() int {
	first = q.data[0]
	q.data = q.data[1:]
	return first
}
 
//In O(n) time, and using O(1) space, remove duplicates from a sorted array
func removeDups(arr []int) []int {

}

 
