package frequency

import (
	"strings"
)

var m map[string]int

func wordFrequency(input string, text string) int {
	fillMap(m, text)
	return m[input]
}

func fillMap(m map[string]int, text string) {
	a := strings.Split(text, " ")
	for _, s := range a {
		if val, ok := m[s]; !ok {
			m[s] = val
		}
	}
}
