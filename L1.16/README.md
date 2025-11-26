## L1.16

This Go snippet demonstrates a quicksort implementation with optimizations for pivot selection and recursion depth.

The quickSort function serves as the entry point, calling qSort to sort the entire slice. It returns the sorted slice while keeping the original array unchanged.

The qSort function uses recursion for partitioning the slice around a pivot. After partitioning, it always recurses on the smaller subarray first, while iterating over the larger subarray in a loop. This helps avoid deep recursion on large subarrays, minimizing the risk of stack overflow in worst-case scenarios.

The medianOfThree function selects the pivot as the median of the left, middle, and right elements. This improves performance on partially sorted data by reducing the chances of poor pivot choices.

For small partitions (size ≤ 10), qSort switches to insertion sort, which is more efficient for small arrays due to its lower overhead.

Elements are partitioned in place: values less than the pivot are moved to the left, values greater than the pivot to the right, and equal values remain in the middle. The algorithm handles duplicates efficiently and tries to minimize unnecessary swaps, helping maintain stability for small subarrays.