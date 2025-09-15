package main

import (
	"fmt"
	"sync"
)

func main() {

	arr := [5]int{2, 4, 6, 8, 10}

	wgExample(arr)
	chExample(arr)

}

// wgExample uses a WaitGroup to synchronize goroutines and prevent premature goroutine completion.
// No mutex is required since each goroutine handles a separate value.
func wgExample(arr [5]int) {
	var wg sync.WaitGroup
	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Println(val * val)
		}(val)
	}
	wg.Wait()
	fmt.Println()
}

// chExample uses a channel to collect results from concurrent goroutines.
// The channel itself acts as a synchronization point for all goroutines.
func chExample(arr [5]int) {
	res := make(chan int, len(arr))
	defer close(res)
	for _, val := range arr {
		go func(val int) {
			res <- val * val
		}(val)
	}
	for range arr {
		fmt.Println(<-res)
	}
}

/*
Output:
100
36
64
4
16

100
36
64
4
16
*/
