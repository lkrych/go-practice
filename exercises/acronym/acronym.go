package acronym

import "strings"

const testVersion = 3

//FieldsFunc allows you to split by multiple runes

func Abbreviate(input string) string {
	acronym := []string{}
	for _, word := range strings.FieldsFunc(input, SplitByTwo) {
		acronym = append(acronym, strings.ToUpper(string(word[0])))
	}
	return strings.Join(acronym, "")
}

func SplitByTwo(r rune) bool {
	return r == ' ' || r == '-'
}
