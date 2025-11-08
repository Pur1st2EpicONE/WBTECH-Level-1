## L1.16

This Go snippet demonstrates an in-place quicksort implementation with optimizations for both pivot selection and recursion depth.

The quickSort function acts as the entry point, calling qSort with the full slice. It returns the sorted slice after in-place sorting.

The qSort function uses a hybrid of loop and recursion to minimize stack depth. After partitioning the slice around a pivot, it always recurses on the smaller subarray first and continues iterating on the larger subarray within the loop. This prevents deep recursion on large subarrays, helping avoid stack overflow in worst-case scenarios.

The medianOfThree function selects the pivot as the median of the left, middle, and right elements. This improves performance on partially sorted arrays by reducing the likelihood of poor pivot choices.

For small partitions (size ≤ 10), qSort switches to insertion sort (insertSort), which is more efficient for small arrays due to lower overhead.

Elements are partitioned in-place: values less than the pivot are moved to the left, values greater than the pivot to the right, and equal values remain in the middle. The algorithm handles duplicates gracefully and preserves stability within partitions as much as possible for small arrays.
