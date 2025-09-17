package main

import (
	"fmt"
	"time"
)

func main() {

	dataCh := make(chan int)

	// stopCh will receive a signal after 5 seconds to stop sending data.
	// Alternatively, you could use context.WithDeadline / context.WithTimeout
	// if you needed to propagate cancellation through multiple functions or goroutines.
	// Another option is time.NewTimer, but the Go documentation notes:
	// "There is no reason to prefer NewTimer when After will do."
	stopCh := time.After(5 * time.Second)

	go sendData(dataCh, stopCh)
	receiveData(dataCh)

}

// sendData sends integers to dataCh every 500 milliseconds
// and stops when stopCh signals
func sendData(dataCh chan<- int, stopCh <-chan time.Time) {
	i := 1
	for {
		select {
		case <-stopCh:
			close(dataCh)
			return
		default:
			dataCh <- i
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// receiveData reads values from dataCh until it is closed
func receiveData(dataCh <-chan int) {
	for val := range dataCh {
		fmt.Printf("Received value from data channel — %d.\n", val)
	}
	fmt.Println("Data channel is closed, exiting.")
}

/*
Output:
Received value from data channel — 1.
Received value from data channel — 2.
Received value from data channel — 3.
Received value from data channel — 4.
Received value from data channel — 5.
Received value from data channel — 6.
Received value from data channel — 7.
Received value from data channel — 8.
Received value from data channel — 9.
Received value from data channel — 10.
Data channel is closed, exiting.
*/
