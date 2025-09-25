package main

import "fmt"

func main() {

	arr := []int{5, 7, 9, 2, 4, 5, 0, 1, 3, 5, 6}
	fmt.Println(quickSort(arr))

}

// quickSort is the entry point that sorts a slice of integers in-place
func quickSort(nums []int) []int {
	qSort(nums, 0, len(nums)-1)
	return nums
}

// qSort implements in-place quicksort. Since Go lacks tail-call optimization,
// it recurses on the smaller partition and loops on the larger one to minimize stack depth.
func qSort(nums []int, left int, right int) {
	for left < right {
		pivot := medianOfThree(nums, left, right)
		nums[pivot], nums[right] = nums[right], nums[pivot]
		pivotEndPos := partition(nums, left, right)
		if pivotEndPos-left < right-pivotEndPos {
			qSort(nums, left, pivotEndPos-1)
			left = pivotEndPos + 1
		} else {
			qSort(nums, pivotEndPos+1, right)
			right = pivotEndPos - 1
		}
	}
}

// medianOfThree selects the pivot as the median of the first, middle, and last elements
// This helps improve performance on partially sorted arrays
func medianOfThree(nums []int, left int, right int) int {
	mid := left + (right-left)/2
	if nums[left] > nums[right] {
		nums[left], nums[right] = nums[right], nums[left]
	}
	if nums[left] > nums[mid] {
		nums[left], nums[mid] = nums[mid], nums[left]
	}
	if nums[mid] > nums[right] {
		nums[mid], nums[right] = nums[right], nums[mid]
	}
	return mid
}

// partition rearranges elements in-place so that elements <= pivot are left
// and elements > pivot are right, then returns the pivot's final index
func partition(nums []int, left int, right int) int {
	i := left
	for j := left; j < right; j++ {
		if nums[j] <= nums[right] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[right], nums[i] = nums[i], nums[right]
	return i
}

/*
Output:
[0 1 2 3 4 5 5 5 6 7 9]
*/
