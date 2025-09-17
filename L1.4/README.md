## L1.4

This Go snippet is functionally identical to L1.3.

signal.NotifyContext internally creates a channel subscribed to os.Signal and automatically cancels the associated context.Context when a signal is received. In other words, it hides the boilerplate of manually wiring up a signal channel and passing it around, while still providing the same behavior.

By using this abstraction, the code stays clean and idiomatic: channels are only for data flow, while shutdown is handled through a context. 

NotifyContext is often seen as the ultimate sugar for handling interrupts, as it keeps shutdown logic consistent with other uses of context.Context.
