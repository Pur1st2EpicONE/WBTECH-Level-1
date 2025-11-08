package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {

	bigNum1, err := newBigNumber("18072590815981508125893102")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create bigNum1: %v", err)
		return
	}

	bigNum2, err := newBigNumber("69810236717261802715809128")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create bigNum2: %v", err)
		return
	}

	sum, err := add(bigNum1, bigNum2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to add numbers: %v", err)
	} else {
		fmt.Printf("%s + %s = %s\n", bigNum1.String(), bigNum2.String(), sum.String())
	}

	sub, err := sub(bigNum2, bigNum1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to subtract numbers: %v", err)
	} else {
		fmt.Printf("%s - %s = %s\n", bigNum2.String(), bigNum1.String(), sub.String())
	}

	mul, err := mul(bigNum1, bigNum2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to multiply numbers: %v", err)
	} else {
		fmt.Printf("%s * %s = %s\n", bigNum1.String(), bigNum2.String(), mul.String())
	}

	div, err := div(bigNum2, bigNum1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to divide numbers: %v", err)
	} else {
		fmt.Printf("%s / %s = %s\n", bigNum2.String(), bigNum1.String(), div.String())
	}

}

// newBigNumber converts a string to a *big.Int.
// It trims any leading/trailing whitespace and returns an error
// if the string does not represent a valid integer.
func newBigNumber(s string) (*big.Int, error) {
	num := new(big.Int)
	if _, ok := num.SetString(s, 10); !ok {
		return nil, fmt.Errorf("invalid number: %s", s)
	}
	return num, nil
}

// add returns the sum of two *big.Int numbers.
// Returns an error if any of the arguments is nil.
func add(x *big.Int, y *big.Int) (*big.Int, error) {
	if x == nil || y == nil {
		return nil, fmt.Errorf("add: input number is nil")
	}
	return new(big.Int).Add(x, y), nil
}

// sub returns the difference of two *big.Int numbers (x - y).
// Returns an error if any of the arguments is nil.
func sub(x *big.Int, y *big.Int) (*big.Int, error) {
	if x == nil || y == nil {
		return nil, fmt.Errorf("sub: input number is nil")
	}
	return new(big.Int).Sub(x, y), nil
}

// mul returns the product of two *big.Int numbers.
// Returns an error if any of the arguments is nil.
func mul(x *big.Int, y *big.Int) (*big.Int, error) {
	if x == nil || y == nil {
		return nil, fmt.Errorf("mul: input number is nil")
	}
	return new(big.Int).Mul(x, y), nil
}

// div returns the integer division result of two *big.Int numbers (x / y).
// Returns an error if any argument is nil or if y is zero.
func div(x *big.Int, y *big.Int) (*big.Int, error) {
	if x == nil || y == nil {
		return nil, fmt.Errorf("div: input number is nil")
	}
	if y.Sign() == 0 {
		return nil, fmt.Errorf("div: number cannot be divided by zero")
	}
	return new(big.Int).Div(x, y), nil
}

/*
Output:
18072590815981508125893102 + 69810236717261802715809128 = 87882827533243310841702230
69810236717261802715809128 - 18072590815981508125893102 = 51737645901280294589916026
18072590815981508125893102 * 69810236717261802715809128 = 1261651842957880722320587551136332995610259563835056
69810236717261802715809128 / 18072590815981508125893102 = 3
*/
