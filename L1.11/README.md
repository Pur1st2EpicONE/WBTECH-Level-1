## L1.11

This Go snippet demonstrates finding the intersection of two slices, removing duplicates.

The program defines a generic function, intersection, which accepts two slices of any comparable type.

It first iterates over the first slice and stores each element in a map for efficient lookup.

Then it iterates over the second slice and checks whether each element exists in the map. If it does, the element is appended to the result slice, and it is removed from the map to ensure uniqueness.
