## L1.22

This Go snippet demonstrates how to perform arithmetic operations—addition, subtraction, multiplication, and division—on arbitrarily large integers using the math/big package. It ensures correctness even for numbers far exceeding the range of standard integer types.

The program first converts input strings representing numbers into *big.Int values via the newBigNumber function. This function validates the input and returns an error if the string does not represent a valid integer.

Each arithmetic operation is implemented as a separate function (add, sub, mul, div) that takes two *big.Int arguments. These functions first check for nil inputs to prevent runtime panics. The division function also checks for division by zero. Operations are performed using big.Int methods such as Add, Sub, Mul, and Div, which handle arbitrarily large numbers safely.
