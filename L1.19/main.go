package main

import "fmt"

func main() {

	str := "REDRUM"
	fmt.Println(reverse(str))

}

// reverse returns the input string reversed.
// It handles Unicode characters correctly by converting the string to a []rune slice.
func reverse(str string) string {
	res := []rune(str)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

/*
Output:
MURDER
*/
