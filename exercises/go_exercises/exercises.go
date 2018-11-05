package goExercises

//return the GCF of the two inputs
func greatestCommonFactor(n1 int, n2 int) int {
	greater := 0
	lesser := 0
	if n1 == n2 {
		return n1
	} else if n1 > n2 {
		greater = n1
		lesser = n2
	} else {
		greater = n2
		lesser = n1
	}

	for i := lesser; i > 1; i-- {
		if (greater%i) == 0 && (lesser%i) == 0 {
			return i
		}
	}
	return 1
}

//return a sorted arr using bubblesort!
func bubbleSort(arr []int) []int {
	var switched bool
	switched = true
	for switched {
		switched = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				switched = true
			}
		}
	}
	return arr
}
