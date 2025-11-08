## L1.23

This Go snippet demonstrates two ways to remove an element from a slice: manually by shifting elements and using the standard library function slices.Delete.

The deleteElement function removes the element at the given index by shifting all elements after it to the left and reducing the slice length by one. It safely handles nil slices and out-of-range indices without panicking.

The slicesDelete function uses slices.Delete from the Go standard library, which performs the same operation in a concise, built-in way while also safely handling invalid indices.
