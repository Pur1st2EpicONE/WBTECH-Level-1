package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	status1 := make(chan struct{})
	status2 := make(chan struct{})
	status3 := make(chan struct{})

	wg.Add(6)

	go condtnExample(&wg)
	go signalExample(sigCh, cancel, &wg)
	go contexExample(ctx, status1, &wg)
	go closedExample(status1, status2, &wg)
	go goexitExample(status2, status3, &wg)
	go timersExample(status3, &wg)

	wg.Wait()

}

// Demonstrates a simple conditional return; prints a prompt once and exits the goroutine.
// Serves as the initial trigger before the interrupt signal is handled.
func condtnExample(wg *sync.WaitGroup) {
	defer wg.Done()
	cnt := 0
	for {
		if cnt > 0 {
			fmt.Println("\n1st return (Would you kindly press Ctrl+C to send OS interrupt?)")
			return
		}
		cnt += 1
	}
}

// Listens for an OS interrupt signal (Ctrl+C)
func signalExample(sigCh chan os.Signal, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-sigCh:
			fmt.Println("\n2st return")
			cancel()
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

// Waits for the context to be canceled
func contexExample(ctx context.Context, status chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("3nd return")
			close(status)
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

// Waits for status channel to close
func closedExample(status, status2 chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-status:
			fmt.Println("4th return")
			close(status2)
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

// Waits for status2 channel to close, exits via Goexit
func goexitExample(status2, status3 chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-status2:
			fmt.Println("5th return")
			close(status3)
			runtime.Goexit()
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

// Waits for status3 to close, then launches two timer-based goroutines
func timersExample(status3 chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	wg.Add(2)
	for {
		select {
		case <-status3:
			timer := time.NewTimer(1 * time.Second)
			go timerExample(timer, wg)
			time.Sleep(2 * time.Second)
			go ctxTimeoutExample(ctx, wg)
			time.Sleep(1 * time.Second)
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

// Waits for a timer to expire
func timerExample(timer *time.Timer, wg *sync.WaitGroup) {
	defer wg.Done()
	<-timer.C
	fmt.Println("6th return")
}

// Waits for context timeout
func ctxTimeoutExample(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ctx.Done()
	fmt.Println("7th return")
}

/*
Output:
1st return (Would you kindly press Ctrl+C to send OS interrupt?)
^C
2st return
3nd return
4th return
5th return
6th return
7th return
*/
