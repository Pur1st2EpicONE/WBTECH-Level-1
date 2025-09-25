package main

import "fmt"

func main() {

	arr := []int{4, 8, 15, 16, 23, 42}

	fmt.Println(binarySearch(arr, 8))
	fmt.Println(binarySearch(arr, 23))
	fmt.Println(binarySearch(arr, 1))

}

// binarySearch performs an iterative binary search on a sorted slice of integers.
// If the target value exists in the slice, it returns its index; otherwise, it returns -1.
func binarySearch(arr []int, val int) int {
	if len(arr) > 0 {
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
	}
	return -1
}

/*
Output:
1
4
-1
*/
