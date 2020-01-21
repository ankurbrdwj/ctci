package frequency

import "testing"

import "fmt"

func TestFindFrequency(t *testing.T) {
	var char rune = rune('a')
	var input string = "aaaaaaaaaarghhhhhhhhhhhh ah"

	result := findfrequency(char, input)
	fmt.Println(result)
	if result != 11 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 11)
	}

}
