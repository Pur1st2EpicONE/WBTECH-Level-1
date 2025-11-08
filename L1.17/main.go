package main

import "fmt"

func main() {

	arr := []int{4, 8, 15, 16, 23, 42}

	fmt.Println(binarySearch(arr, 8))
	fmt.Println(binarySearch(arr, 23))
	fmt.Println(binarySearch(arr, 1))

}

// binarySearch performs an iterative binary search on a sorted slice of integers.
// It returns the index of the target value if it exists; otherwise, it returns -1.
// Note: if the slice contains duplicates, it may return any matching index.
func binarySearch(arr []int, val int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

/*
Output:
1
4
-1
*/
