## Week 3 Quiz

1. Suppose you want to start a goroutine which executes a function called `test1()`. What code would create this goroutine?

![](https://user-images.githubusercontent.com/49638680/92456008-a7d62480-f1c2-11ea-8ed2-5d11e46b5aed.png)

For concurrent execution, Go has a keyword `go`.
This indicates the program to send the code to Go runtime scheduler in order to schedule a concurrent run.

2. When does a goroutine complete?

![](https://user-images.githubusercontent.com/49638680/92456013-a99fe800-f1c2-11ea-8e28-834d50753aa4.png)

Goroutines complete when their code is completely executed and when their parent routine has completed.

3. Synchronisation is useful for what purpose?

![](https://user-images.githubusercontent.com/49638680/92456017-aa387e80-f1c2-11ea-999c-baaaba50bac5.png)

Synchronisation is a way to force some interleaves and exclude others.
This is useful to avoid race conditions, that is illegal interleaves, and to force a certain order of execution of different goroutines.

4. If a goroutine `g1` is using a WaitGroup `wg` to wait until another goroutine `g2` completes a task, what method of the the WaitGroup should be called when `g2` has finished the task?

![](https://user-images.githubusercontent.com/49638680/92456019-aad11500-f1c2-11ea-8028-911430652d84.png)

A WaitGroup uses a “semaphore” counter to authorise the running of code.
The counter is increased by the number of goroutines to be executed.
In order to decrease the counter once the goroutine has finished its task, we call the WaitGroup method `Done()`.

5. If a goroutine `g1` is using a WaitGroup `wg` to wait until another goroutine `g2` completes a task, what method of the the WaitGroup should be called before `g2` starts its task?

![](https://user-images.githubusercontent.com/49638680/92456021-ab69ab80-f1c2-11ea-871e-2d23ba359a8d.png)

As for the previous question, we have to refer to the “semaphore” counter of the WaitGroup.
Here we need to increase the counter in order to make it wait for the `g2` routine to decrease it.
Hence, we need to call the method `Add()`.

6. How might you write code to allow a goroutine to receive data from a channel `c`?

![](https://user-images.githubusercontent.com/49638680/92456023-ac024200-f1c2-11ea-8476-026df6935c04.png)

In order to send data on a channel and to receive data from it, the operator is the same.
We use arrow operator `<-` with the channel on the left hand side to send data on that channel.
If the channel is on the right hand side the operator is to receive data from it.
We also have to insert an equal sign in order to store data from the channel in a variable.

7. What is the difference between a buffered channel and an unbuffered channel?

![](https://user-images.githubusercontent.com/49638680/92456024-ac024200-f1c2-11ea-9566-3628f6de0e63.png)

By default, channels in Go are unbuffered, meaning they have no memory to store values.
This makes every sending/receiving instruction a wait point for the program, since every time you send data on the channel, you wait for another piece of code in another goroutine to receive them.
With buffered channels, _i.e._ you allow channel to store variables in an accessory memory, so you are not forced to wait another routine to free the channel in order to continue execution.
