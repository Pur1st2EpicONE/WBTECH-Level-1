package main

import "fmt"

func main() {

	str := "REDRUM"
	fmt.Println(reverse(str))

}

// reverse returns the input string reversed.
// It handles Unicode characters correctly by converting the string to a []rune slice.
func reverse(str string) string {
	if len(str) == 0 {
		return str
	}
	res := []rune(str)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return string(res)
}

/*
Output:
MURDER
*/
