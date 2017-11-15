package accumulate

const testVersion = 1

func Accumulate(input []string, accFunc func(string) string) []string {
	acc := []string{}
	for _, el := range input {
		acc = append(acc, accFunc(el))
	}
	return acc
}
