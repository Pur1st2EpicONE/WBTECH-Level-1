package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {

	i := 5
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(slice)

	if err := deleteElement(&slice, i); err != nil {
		fmt.Fprintf(os.Stderr, "error deleting element: %v\n", err)
	} else {
		fmt.Println(slice)
	}

	if err := slicesDelete(&slice, i); err != nil {
		fmt.Fprintf(os.Stderr, "error deleting element: %v\n", err)
	} else {
		fmt.Println(slice)
	}

}

// deleteElement removes the element at the given index from the slice
// by shifting the elements after it to the left and reducing the slice length.
// Returns an error if the slice is nil or the index is out of range.
func deleteElement(slice *[]int, index int) error {

	if slice == nil {
		return fmt.Errorf("deleteElement: slice is nil")
	} else if index < 0 || index >= len(*slice) {
		return fmt.Errorf("deleteElement: index out of range")
	}

	copy((*slice)[index:], (*slice)[index+1:])
	*slice = (*slice)[:len(*slice)-1]

	return nil

}

// slicesDelete removes the element at the given index from the slice
// using the standard library slices.Delete function.
// Returns an error if the slice is nil or the index is out of range.
func slicesDelete(slice *[]int, index int) error {

	if slice == nil {
		return fmt.Errorf("slicesDelete: slice is nil")
	} else if index < 0 || index >= len(*slice) {
		return fmt.Errorf("slicesDelete: index out of range")
	}

	*slice = slices.Delete(*slice, index, index+1)

	return nil

}
