package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var number int64 = 5
	bitPos := checkArgs()

	runExample(number, bitPos)

}

// checkArgs parses the first command-line argument as a bit position
func checkArgs() uint64 {
	var bitPos uint64 = 0
	if len(os.Args) > 1 {
		newBitPos, err := strconv.Atoi(os.Args[1])
		if err != nil || newBitPos < 0 || newBitPos > 63 {
			fmt.Println("Failed to parse bit position — defaulting to 0.")
		} else {
			bitPos = uint64(newBitPos)
		}
	} else {
		fmt.Println("No bit position specified — defaulting to 0.")
	}
	return bitPos
}

// runExample demonstrates a round-trip toggle of a single bit
func runExample(number int64, bitPos uint64) {
	fmt.Printf("Initial number: [%064b] %d\n", number, number)
	switchBit(&number, bitPos)
	fmt.Printf("Bit switched:   [%064b] %d\n", number, number)
	switchBit(&number, bitPos)
	fmt.Printf("Bit restored:   [%064b] %d\n", number, number)
}

// switchBit flips the bit at the given position.
func switchBit(number *int64, bit uint64) {
	*number ^= 1 << bit
}

/*
Output:
No bit position specified — defaulting to 0.
Initial number: [0000000000000000000000000000000000000000000000000000000000000101] 5
Bit switched:   [0000000000000000000000000000000000000000000000000000000000000100] 4
Bit restored:   [0000000000000000000000000000000000000000000000000000000000000101] 5
*/
