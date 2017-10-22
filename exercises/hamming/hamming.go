package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("The two strings need to be of equal length.")
	}
	hamming := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hamming++
		}
	}
	return hamming, nil
}
