## L1.9

This Go snippet demonstrates building a simple number processing pipeline with goroutines and channels.

The program defines two stages:

sendNums writes numbers from an array into the first channel.

multiply reads those numbers, doubles each value, and forwards the results into the second channel.

The main function launches both stages as goroutines and then reads the processed values from the second channel, printing them to stdout.

Proper channel closing ensures the pipeline terminates gracefully once all numbers are processed.
