package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func main() {

	workers := checkArgs()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	ch := make(chan int)
	var wg sync.WaitGroup

	for wokerID := range workers {
		wg.Add(1)
		go worker(ch, wokerID+1, &wg)
	}

	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Main — received signal interrupt, stopping workers.")
			close(ch)
			wg.Wait()
			return
		case ch <- i:
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// checkArgs parses the number of workers from command-line arguments.
// Returns the number of workers and defaults to 3 if argument is invalid or missing.
func checkArgs() int {
	amount := 3
	if len(os.Args) > 1 {
		newAmount, err := strconv.Atoi(os.Args[1])
		if err != nil || newAmount < 1 {
			fmt.Println("Failed to parse worker amount, using default value of 3.")
		} else {
			amount = newAmount
		}
	} else {
		fmt.Println("No worker amount specified, using default value of 3.")
	}
	return amount
}

// worker reads integers from the channel and processes them.
// Each worker runs in its own goroutine and stops when the channel is closed.
func worker(ch <-chan int, workerID int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("Worker %d — processing value %d.\n", workerID, val)
	}
	fmt.Printf("Worker %d — stopped.\n", workerID)
}

/*
Output:
go run main.go 5
Worker 5 — processing value 1.
Worker 1 — processing value 2.
Worker 5 — processing value 3.
Worker 4 — processing value 4.
Worker 2 — processing value 5.
^C
Main — received signal interrupt, stopping workers.
Worker 3 — stopped.
Worker 5 — stopped.
Worker 1 — stopped.
Worker 2 — stopped.
Worker 4 — stopped.
*/
