package goExercises

//return true if n is a power of two
func isPowerOfTwo(n float64) bool {
	power := 0
	for power(2, power) < n {
		if power(2, power) == n {
			return true
		}
		power++
	}
}

func power(base, exp int) float64 {

}

//return a sorted arr, partitioned by odd and even, with even in the front and odd in the back of the result array
func partitionEvenOdd(arr []int) []int {

}

//check to see if a passphrase is valid, it must not contain any duplicate words
func passPhrase(passphrase string) bool {

}

func passPhraseAdvent() int {
	//read from ../practice/passphrase_input.txt
	//each line has a different passphrase, return how many are valid
}

//bonus, check to see if a passphrase is valid, it must not contain any words that are anagrams of eachother
// abcde == dcabe
// func passPhraseAnagram(passphrase string) bool {

// }

//write code to partition a linked list around a value x, such that all nodes less than
//x come before all nodes greater than or equal to x
func partitionLL(head *Node, partition int) *Node {

}

//return sum of n + n-1 + ... 1
func sumNums(n int) int {

}
