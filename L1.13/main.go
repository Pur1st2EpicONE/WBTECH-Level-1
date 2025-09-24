package main

import "fmt"

func main() {

	num1, num2 := 1, 2
	fmt.Println(num1, num2)

	num1, num2 = num2, num1
	fmt.Println(num1, num2)

	num1 = num1 ^ num2
	num2 = num1 ^ num2
	num1 = num1 ^ num2
	fmt.Println(num1, num2)

	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Println(num1, num2)

}

/*
Output:
1 2
2 1
1 2
2 1
*/
