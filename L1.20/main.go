package main

import (
	"fmt"
)

func main() {

	str := "snow dog sun"
	fmt.Println(reverseWords(str))

}

// reverseWords reverses the order of words in a given string.
// Words are assumed to be separated by single spaces. The reversal
// is done in-place using a rune slice to properly handle Unicode characters.
func reverseWords(str string) string {

	if len(str) == 0 {
		return str
	}

	runes := []rune(str)
	reverseSlice(runes, 0, len(runes)-1)

	left := 0
	for right := 0; right <= len(runes); right++ {
		if right == len(runes) || runes[right] == ' ' {
			reverseSlice(runes, left, right-1)
			left = right + 1
		}
	}

	return string(runes)

}

// reverseSlice reverses the elements of a rune slice in-place
// between the indices start and end inclusive.
func reverseSlice(runes []rune, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

/*
Output:
sun dog snow
*/
