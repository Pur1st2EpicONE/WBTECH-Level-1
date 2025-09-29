## L1.20

This Go snippet demonstrates how to reverse the order of words in a string.

The reverseWords function first converts the input string to a slice of runes. This ensures that multi-byte characters, such as Cyrillic letters, are treated as single elements rather than individual bytes.

It then reverses the entire rune slice using the two-pointer approach: one pointer starts at the beginning (i) and the other at the end (j) of the slice. In each iteration, the elements at i and j are swapped, and the pointers move towards the center (i++, j--). This step inverts the order of all characters in the string.

After that, the function iterates over the slice to locate word boundaries (spaces) and reverses each word individually using the same two-pointer method. The last word is reversed after the loop.

Finally, the rune slice is converted back into a string and returned. This approach ensures that the word order is reversed correctly for any Unicode string, without allocating additional slices for individual words.
