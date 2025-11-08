package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	num1, num2 := 1, 2
	fmt.Println(num1, num2)

	if err := swapXOR(&num1, &num2); err != nil {
		fmt.Fprintf(os.Stderr, "did not swap: %v\n", err)
	} else {
		fmt.Println(num1, num2)
	}

	if err := swapArithmetic(&num1, &num2); err != nil {
		fmt.Fprintf(os.Stderr, "did not swap: %v\n", err)
	} else {
		fmt.Println(num1, num2)
	}

}

// swapXOR swaps two integers using XOR operations.
// Does not use a temporary variable and works for all int values.
// Returns an error if either pointer is nil.
// Does nothing if both pointers refer to the same variable.
func swapXOR(num1 *int, num2 *int) error {

	if num1 == nil || num2 == nil {
		return fmt.Errorf("swapXOR: nil pointer passed, swap skipped")
	}
	if num1 == num2 || *num1 == *num2 {
		return nil
	}

	*num1 = *num1 ^ *num2
	*num2 = *num1 ^ *num2
	*num1 = *num1 ^ *num2

	return nil

}

// swapArithmetic swaps two integers using addition and subtraction.
// Avoids using a temporary variable but checks for potential overflow.
// Returns an error if either pointer is nil or if addition would overflow.
// Does nothing if both pointers refer to the same variable.
func swapArithmetic(num1 *int, num2 *int) error {

	if num1 == nil || num2 == nil {
		return fmt.Errorf("swapArithmetic: nil pointer passed, swap skipped")
	}
	if num1 == num2 || *num1 == *num2 {
		return nil
	}
	if (*num2 > 0 && *num1 > math.MaxInt-*num2) || (*num2 < 0 && *num1 < math.MinInt-*num2) {
		return fmt.Errorf("swapArithmetic: possible int overflow, swap skipped")
	}

	*num1 = *num1 + *num2
	*num2 = *num1 - *num2
	*num1 = *num1 - *num2

	return nil

}

/*
Output:
1 2
2 1
1 2
*/
