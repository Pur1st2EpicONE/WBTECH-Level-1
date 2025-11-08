## L1.13

This Go snippet demonstrates two safe ways to swap the values of two integers without using a third variable. The program first initializes two integer variables and prints their initial values.

It then swaps the values using the XOR bitwise operation. This method does not require a temporary variable and works for all int values. It safely handles nil pointers and the case when both pointers refer to the same variable.

Next, it swaps the values using arithmetic operations (addition and subtraction). This method also avoids an extra variable but includes checks for nil pointers, identical variables, and potential integer overflow to prevent runtime errors.

Each function returns an error if the swap cannot be safely performed; otherwise, it modifies the values in place. The output shows the values before and after each swap.
