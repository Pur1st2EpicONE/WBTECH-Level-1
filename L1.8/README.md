## L1.8

This Go snippet demonstrates manipulating individual bits in an integer using bit masks.

The switchBit function toggles a bit at a specified position using the XOR (^) operator. Applying XOR with a mask flips the target bit: if it was 0, it becomes 1; if it was 1, it becomes 0.

The runExample function shows a round-trip toggle by calling switchBit twice, effectively switching the bit there and back.

Command-line arguments allow specifying which bit to toggle; if omitted or invalid, the program defaults to bit 0.
