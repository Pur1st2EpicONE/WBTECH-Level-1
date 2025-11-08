## L1.6

This Go snippet demonstrates several common techniques for gracefully stopping goroutines.
Each method is implemented as a separate example, chained together to show how different stop mechanisms can interact.

The program includes the following approaches:

1) Conditional return — the condtnExample goroutine exits after a simple condition check.

2) OS signal — the signalExample goroutine listens for SIGINT (Ctrl+C) and returns when the signal is received.

3) Context cancellation — the contexExample goroutine stops when its context is canceled.

4) Channel closure — the closedExample goroutine terminates when it detects that a channel has been closed.

5) Runtime exit — the goexitExample goroutine calls runtime.Goexit to stop itself immediately.

6) Timer-based stop — the timerExample goroutine waits for a time.Timer to expire.

7) Context timeout — the ctxTimeoutExample goroutine exits automatically when its context times out.

Each approach demonstrates a clean and safe way to stop concurrent execution, avoiding goroutine leaks and ensuring all resources are released properly. The examples also illustrate how multiple termination mechanisms can be combined to coordinate complex shutdown sequences.