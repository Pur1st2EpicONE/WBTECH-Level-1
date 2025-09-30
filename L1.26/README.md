## L1.26

This Go snippet demonstrates how to check whether a string consists of unique characters.

The isUnique function iterates over each character in the string, converting it to lowercase with unicode.ToLower. It uses a map[rune]struct{} as a set to track which characters have been seen. If a character appears more than once, the function returns false. If the loop completes without finding duplicates, it returns true.
