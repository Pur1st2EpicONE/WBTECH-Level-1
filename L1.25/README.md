## L1.25

This Go snippet demonstrates two ways to pause execution for a specified duration.

The sleep function uses <-time.After(duration) to suspend the current goroutine efficiently without consuming CPU resources. In contrast, restlessSleep uses a busy-wait loop that repeatedly checks the elapsed time, which is CPU-intensive and inefficient.
