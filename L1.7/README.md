## L1.7

This Go snippet demonstrates concurrency-safe map operations implemented in two different ways.

The first approach uses sync.Map, which provides built-in support for concurrent access without explicit locking.

The second approach relies on a regular map protected by a sync.Mutex, ensuring that only one goroutine writes at a time.
