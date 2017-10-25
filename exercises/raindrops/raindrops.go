package raindrops

import (
	"strconv"
	"strings"
)

const testVersion = 3

func Convert(num int) string {
	rain := []string{}
	if num%3 == 0 {
		rain = append(rain, "Pling")
	}
	if num%5 == 0 {
		rain = append(rain, "Plang")
	}
	if num%7 == 0 {
		rain = append(rain, "Plong")
	}
	if len(rain) == 0 {
		rain = append(rain, strconv.Itoa(num))
	}
	return strings.Join(rain, "")
}
