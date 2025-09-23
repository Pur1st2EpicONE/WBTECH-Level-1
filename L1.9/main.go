package main

import "fmt"

func main() {

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go sendNums(ch1, arr)
	go multiply(ch1, ch2)

	for val := range ch2 {
		fmt.Println(val)
	}

}

// sendNums writes all numbers from the array into the channel and then closes it
func sendNums(ch chan<- int, arr [10]int) {
	for _, val := range arr {
		ch <- val
	}
	close(ch)
}

// multiply reads numbers from ch1, doubles them, and sends them to ch2
func multiply(ch1 <-chan int, ch2 chan<- int) {
	for val := range ch1 {
		ch2 <- val * 2
	}
	close(ch2)
}

/*
Output:
2
4
6
8
10
12
14
16
18
20
*/
