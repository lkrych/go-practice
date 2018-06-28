func fibonacci(n int) []int {
	fibArr := []int{}
	if n < 2 {
		fibArr = append(fibArr, n)
		return fibArr
	}
	fibArr = append(fibArr, fibonacci(n-1)+fibonacci(n-2))
	return fibArr
}