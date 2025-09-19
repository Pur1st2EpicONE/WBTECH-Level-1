## L1.6

This Go snippet demonstrates several common ways to stop goroutines in Go, each implemented in a separate example.

The program includes the following approaches:

1) OS signal — the signalExample goroutine listens for SIGINT (Ctrl+C) and returns when a signal is received.

2) Context cancellation — the contexExample goroutine exits when its context is canceled.

3) Channel closure — the closedExample goroutine waits for a channel to close before terminating.

4) Runtime exit — the goexitExample goroutine calls runtime.Goexit to stop itself.

5) Timer-based stop — the timerExample goroutine waits for a time.Timer to expire.

6) Context timeout — the ctxTimeoutExample goroutine exits when its context times out.

Each approach ensures that goroutines terminate cleanly, without leaving blocked operations or unprocessed data. The snippet also illustrates how these methods can be combined to orchestrate graceful shutdown sequences.