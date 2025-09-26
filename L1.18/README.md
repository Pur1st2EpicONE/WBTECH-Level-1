## L1.18

This Go snippet demonstrates two ways to implement a concurrent counter: using a mutex and an atomic integer.

The mutexCounter uses sync.Mutex to ensure that increments to the counter are safe across multiple goroutines. Each goroutine locks the mutex before modifying the counter and unlocks it afterwards, preventing race conditions.

The atomicCounter uses atomic.Int64 from the sync/atomic package to perform lock-free atomic increments. Operations like Add and Load are guaranteed to be safe for concurrent access without requiring explicit locking.
