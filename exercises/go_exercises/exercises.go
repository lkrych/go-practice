package goExercises

import ("strings"; "fmt")
//given a string, write a function to check it is a permutation of a palindrome
func palindromePermutation(perm string) bool {
	sorted := quickSort(strings.Split(perm, ""))
	charMap := make(map[string]int)
	fmt.Println("The sorted input is", sorted)
	//sort the chars and iterate through them
	for _, char := range sorted {
		if char == " " { continue }
		if count, ok := charMap[char]; ok {
			charMap[char] = count + 1
		} else {
			charMap[char] = 1
		}
	}
	foundOdd := false
	fmt.Printf("charMap is %v for word %v \n", charMap, perm)
	for _, v := range charMap {
		
		if v % 2 != 0 {
			if foundOdd {
				//there can't be more than one odd count
				return false
			} else {
				foundOdd = true
			}
		}
	}
	return true
}

func quickSort(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	left := []string{}
	right := []string{}

	for i := 1; i < len(arr); i++ {
		if arr[i] >= pivot {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}

	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)
	sorted := []string{}

	sorted = append(sorted, sortedLeft...)
	sorted = append(sorted, pivot)
	sorted = append(sorted, sortedRight...)

	return sorted
} 
 
//return if the word is a palindrome
func isPalindrome(word string) bool {
	for i,j := 0, len(word) -1; i < j; i,j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}
 
