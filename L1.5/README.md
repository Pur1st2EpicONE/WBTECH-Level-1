## L1.5

This Go snippet demonstrates sending data to a channel while stopping after a fixed duration using time.After.

The sendData goroutine writes integers to dataCh every 500 milliseconds and exits gracefully by closing the channel when stopCh signals after 5 seconds.

The receiveData function reads from dataCh until it is closed, printing each received value. This pattern ensures a clean shutdown without leaving any goroutines blocked or data unprocessed.

Unlike using a context, this approach is simple and well-suited for a single, local timeout. stopCh does not require manual closing, and the shutdown logic remains straightforward.
