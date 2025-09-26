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
	var mutex mutexCounter
	var atomic atomicCounter

	mutexExample(&mutex, &wg)
	atomicExample(&atomic, &wg)

	wg.Wait()

	fmt.Printf("mutex value — %d\n", mutex.cnt)
	fmt.Printf("atomic value — %d\n", atomic.cnt.Load())

}

// mutexExample increments the counter using a mutex to protect access.
func mutexExample(mutex *mutexCounter, wg *sync.WaitGroup) {
	for range 10 {
		wg.Go(func() {
			mutex.mu.Lock()
			mutex.cnt++
			mutex.mu.Unlock()
		})
	}
}

// atomicExample increments the counter using atomic operations.
func atomicExample(atomic *atomicCounter, wg *sync.WaitGroup) {
	for range 20 {
		wg.Go(func() {
			atomic.cnt.Add(1)
		})
	}
}

/*
Output:
mutex value — 10
atomic value — 20
*/
