package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type mutexCounter struct {
	cnt int
	mu  sync.Mutex
}

type atomicCounter struct {
	cnt atomic.Int64
}

func main() {

	var wg sync.WaitGroup
	var mc mutexCounter
	var ac atomicCounter

	mutexExample(&mc, &wg)
	atomicExample(&ac, &wg)

	wg.Wait()

	fmt.Printf("mutex value — %d\n", mc.cnt)
	fmt.Printf("atomic value — %d\n", ac.cnt.Load())

}

// mutexExample increments the counter using a mutex to protect access.
func mutexExample(mc *mutexCounter, wg *sync.WaitGroup) {
	for range 10 {
		wg.Go(func() {
			mc.mu.Lock()
			mc.cnt++
			mc.mu.Unlock()
		})
	}
}

// atomicExample increments the counter using atomic operations.
func atomicExample(ac *atomicCounter, wg *sync.WaitGroup) {
	for range 20 {
		wg.Go(func() {
			ac.cnt.Add(1)
		})
	}
}

/*
Output:
mutex value — 10
atomic value — 20
*/
