## L1.19

This Go snippet demonstrates how to reverse a string while correctly handling Unicode characters.

The reverse function first converts the input string to a slice of runes. This ensures that multi-byte characters, such as Cyrillic letters or emoji, are treated as single elements rather than individual bytes.

It then uses a single-loop approach with an index i iterating from the start to the middle of the slice. In each iteration, the function swaps the element at i with the element at the corresponding position from the end (len(runes)-i-1). The loop stops when all elements up to the middle have been swapped.

Finally, the rune slice is converted back into a string and returned. This approach ensures that the reversal works for any Unicode string, not just ASCII.
