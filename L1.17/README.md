## L1.17

This Go snippet demonstrates an iterative binary search on a sorted slice of integers.

The binarySearch function repeatedly halves the search interval to locate a target value. It uses mid := left + (right-left)/2 to calculate the midpoint safely, preventing potential integer overflow that could occur with (left + right) / 2 when the indices are large.

If the target value exists in the slice, the function returns its index. Otherwise, it returns -1. This iterative approach avoids recursion and provides an efficient O(log n) search time, making it suitable for large slices.
