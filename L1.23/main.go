package main

import (
	"fmt"
	"slices"
)

func main() {

	i := 5
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(slice)

	delete(&slice, i)
	fmt.Println(slice)

	slicesDelete(&slice, i)
	fmt.Println(slice)

}

// delete removes an element from the slice at the given index
// by creating a new slice and copying elements before and after the index.
func delete(slice *[]int, index int) {
	if index >= 0 && index < len(*slice) {
		new := make([]int, len(*slice)-1)
		copy(new, (*slice)[:index])
		copy(new[index:], (*slice)[index+1:])
		*slice = new
	}
}

// slicesDelete removes an element from the slice at the given index
// using the standard library function slices.Delete
func slicesDelete(slice *[]int, index int) {
	if index >= 0 && index < len(*slice) {
		*slice = slices.Delete(*slice, index, index+1)
	}
}

/*
Output:
[1 2 3 4 5 6 7 8]
[1 2 3 4 5 7 8]
[1 2 3 4 5 8]
*/
