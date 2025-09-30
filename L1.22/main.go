package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := big.NewInt(123_456_789)
	b := big.NewInt(987_654_321)

	sum := big.NewInt(0).Add(a, b)
	sub := big.NewInt(0).Sub(b, a)
	mul := big.NewInt(0).Mul(a, b)
	div := big.NewInt(0).Div(b, a)

	fmt.Printf("%d + %d = %d\n", a, b, sum)
	fmt.Printf("%d - %d = %d\n", b, a, sub)
	fmt.Printf("%d * %d = %d\n", a, b, mul)
	fmt.Printf("%d / %d = %d\n", b, a, div)

}

/*
Output:
123456789 + 987654321 = 1111111110
987654321 - 123456789 = 864197532
123456789 * 987654321 = 121932631112635269
987654321 / 123456789 = 8
*/
