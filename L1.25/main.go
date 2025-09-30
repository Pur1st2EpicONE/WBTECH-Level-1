package main

import (
	"time"
)

func main() {

	sleep(5 * time.Second)
	restlessSleep(5 * time.Second)

}

// sleep pauses the current goroutine for the given duration
// using a channel receive from time.After (efficient, non-blocking to the CPU)
func sleep(duration time.Duration) {
	<-time.After(duration)
}

// restlessSleep pauses the current goroutine for the given duration
// using a busy-wait loop (CPU-intensive and inefficient)
func restlessSleep(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
	}
}
