## L1.19

This Go snippet demonstrates how to reverse a string while correctly handling Unicode characters.

The reverse function first converts the input string to a slice of runes. This ensures that multi-byte characters, such as Cyrillic letters, are treated as single elements rather than individual bytes.

It then uses a two-pointer approach: one index starts at the beginning (i) and the other at the end (j) of the slice. In each iteration, the function swaps the elements at i and j, then moves the pointers towards the center (i++, j--). The loop stops when the pointers meet or cross.

Finally, the rune slice is converted back into a string and returned. This approach ensures that the reversal works for any Unicode string, not just ASCII.
