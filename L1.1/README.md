## L1.1

This Go snippet demonstrates struct embedding and method shadowing. We define a Human struct with fields Name and Age and a method handleApple(). Then we define an Action struct that embeds Human, which allows Action to inherit all of Human's methods automatically.

At the same time, Action defines its own handleApple() method, which shadows the embedded Human's method. Calling action.handleApple() invokes Action's version, while action.Human.handleApple() explicitly calls the original method from Human. 

This highlights how Go uses composition to simulate inheritance and how to handle method overriding when needed.