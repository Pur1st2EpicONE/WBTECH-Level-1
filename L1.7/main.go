package main

import (
	"sync"
)

func main() {

	sm := new(sync.Map)
	hm := make(map[int]int)

	var m sync.Mutex
	var wg sync.WaitGroup

	for i := range 100 {
		wg.Add(2)
		go fillSyncMap(sm, i, i*2, &wg)
		go fillHashMap(hm, i, i*2, &wg, &m)
	}

	wg.Wait()

}

// Stores a key-value pair into a sync.Map (no explicit locking needed)
func fillSyncMap(sm *sync.Map, key int, val int, wg *sync.WaitGroup) {
	defer wg.Done()
	sm.Store(key, val)
}

// Stores a key-value pair into a regular map (requires explicit locking via Mutex)
func fillHashMap(hm map[int]int, key int, val int, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	hm[key] = val
	m.Unlock()
}
