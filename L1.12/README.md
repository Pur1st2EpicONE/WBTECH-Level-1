## L1.12

This Go snippet demonstrates extracting the unique elements from a slice, effectively converting it into a set.

The program defines a generic function, makeSet, which accepts a slice of any comparable type.

It iterates over the input slice and stores each element in a map to eliminate duplicates.

Then it collects all map keys into a new slice, which is returned as the set of unique elements.
