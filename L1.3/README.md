## L1.3

This Go snippet demonstrates a common pattern of distributing work to multiple concurrent goroutines and coordinating their execution using channels and a WaitGroup.

The program continuously generates integer values in the main goroutine and sends them into a channel. A set of worker goroutines, created at startup according to a specified number, read values from this channel and process them independently.

The WaitGroup ensures that all workers have finished processing before the program exits. Closing the channel acts as a signal to the workers that no more data will be sent, allowing them to terminate gracefully. Using this pattern, the program avoids race conditions, guarantees that no data is lost, and cleanly handles shutdown on receiving an interrupt signal (Ctrl+C).
