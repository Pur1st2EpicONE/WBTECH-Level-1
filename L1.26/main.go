package main

import (
	"fmt"
	"unicode"
)

func main() {

	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("aabcd"))
	fmt.Println(isUnique("abCdefAaf"))

}

// isUnique iterates over each rune, converts it to lowercase, and tracks seen characters using a map.
// Returns true if all characters are unique, false if any duplicates are found.
func isUnique(str string) bool {
	hm := make(map[rune]struct{})
	var lowerVal rune
	for _, val := range str {
		lowerVal = unicode.ToLower(val)
		if _, ok := hm[lowerVal]; ok {
			return false
		}
		hm[lowerVal] = struct{}{}
	}
	return true
}

/*
Output:
true
false
false
*/
