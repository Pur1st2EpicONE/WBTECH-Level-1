## L1.12

This Go snippet demonstrates two approaches for extracting the unique elements from a slice, effectively converting it into a set. The program defines generic functions that accept slices of any comparable type. Both approaches iterate over the input slice and use maps to eliminate duplicates by storing each element as a map key, then collecting the unique keys into a new slice.

The first approach uses two passes through the data - first populating the map with all elements, then extracting the map keys into a result slice. This method is efficient but returns elements in random order based on map iteration.

The second approach uses a single pass, checking the map during iteration and immediately appending new elements to the result slice. This preserves the original order of first occurrence while still leveraging the map for duplicate detection, providing more predictable output at the cost of some performance.