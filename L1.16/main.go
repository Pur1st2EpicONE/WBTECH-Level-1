package main

import "fmt"

func main() {

	arr := []int{5, 7, 9, 2, 4, 5, 0, 1, 3, 5, 6}
	fmt.Println(quickSort(arr))

}

// quickSort sorts a slice of integers in ascending order using
// the quicksort algorithm. It returns a sorted copy of the input slice,
// leaving the original slice unchanged.
func quickSort(nums []int) []int {

	if nums == nil {
		return nil
	}

	res := make([]int, len(nums))
	copy(res, nums)

	if len(res) <= 1 {
		return res
	}

	qSort(res, 0, len(res)-1)

	return res

}

// qSort recursively sorts the subarray nums[left:right] using
// the quicksort algorithm. It switches to insertion sort
// for partitions of size 10 or less for better performance.
func qSort(nums []int, left int, right int) {

	for left < right {

		if right-left <= 10 {
			insertSort(nums, left, right)
			return
		}

		median := medianOfThree(nums, left, right)
		nums[left], nums[median] = nums[median], nums[left]
		pivot := nums[left]

		lessThan := left
		i := left + 1
		greaterThan := right

		for i <= greaterThan {
			if nums[i] < pivot {
				nums[lessThan], nums[i] = nums[i], nums[lessThan]
				lessThan++
				i++
			} else if nums[i] > pivot {
				nums[i], nums[greaterThan] = nums[greaterThan], nums[i]
				greaterThan--
			} else {
				i++
			}
		}

		if lessThan-left < right-greaterThan { // Recurse on the smaller partition to limit stack depth
			qSort(nums, left, lessThan-1)
			left = greaterThan + 1
		} else {
			qSort(nums, greaterThan+1, right)
			right = lessThan - 1
		}

	}

}

// medianOfThree returns the index of the median value among nums[left],
// nums[mid], and nums[right]. This helps to improve pivot selection
// and avoid worst-case performance on sorted inputs.
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

// insertSort sorts the subarray nums[left:right] using insertion sort.
// This is efficient for small partitions.
func insertSort(nums []int, left int, right int) {
	for i := left + 1; i <= right; i++ {
		key := nums[i]
		j := i - 1
		for j >= left && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

/*
Output:
[0 1 2 3 4 5 5 5 6 7 9]
*/
