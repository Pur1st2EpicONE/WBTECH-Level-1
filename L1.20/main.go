package main

import (
	"fmt"
)

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str))

}

// reverseWords reverses the order of words in a string.
// It works "in-place" using a []rune slice, so it correctly handles Unicode.
func reverseWords(str string) string {
	res := []rune(str)
	reverse(res)
	j := 0
	for i, val := range res {
		if val == ' ' {
			reverse(res[j:i])
			j = i + 1
		}
	}
	reverse(res[j:])
	return string(res)
}

// reverse reverses a slice of runes "in-place" using the two-pointers technique
func reverse(str []rune) {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}

/*
Output:
sun dog snow
*/
