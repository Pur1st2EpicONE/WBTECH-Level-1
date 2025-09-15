## L1.2

This Go snippet demonstrates two common ways to coordinate concurrent goroutines and collect their results.

In the first case, wgExample launches multiple goroutines, each calculating the square of an element from an array, and waits for all of them to finish before continuing. This ensures that the program does not exit prematurely while goroutines are still running, and no mutex is required because each goroutine operates on its own independent value.

In the second case, chExample uses a buffered channel to gather results from concurrent goroutines, where each goroutine sends its calculated square into the channel. Receiving from the channel acts as a synchronization point, ensuring that all goroutines have completed and their results are collected.
