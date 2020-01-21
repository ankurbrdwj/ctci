package frequency

import (
	"unicode/utf8"
)

func findfrequency(char rune, input string) int {
	var result int = 0
	for len(input) > 0 {
		r, size := utf8.DecodeRuneInString(input)
		//fmt.Printf("%c %v\n", r, size)
		if r == char {
			result++
		}
		input = input[size:]
	}
	return result
}
