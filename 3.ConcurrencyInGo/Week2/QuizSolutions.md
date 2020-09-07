## Week 2 Quiz

1. What unique resources does a process have?

![](https://user-images.githubusercontent.com/49638680/92366019-e69da900-f0f4-11ea-912e-0c2e95cb4d5b.png)

A process is identified by unique resources it has.
In particular, it has a code, a program counter value, some virtual address space. These are resources it owns uniquely. Furthermore it may share some libraries with other processes.

2. What does an operating system do to enable concurrency on a single processor machine?

![](https://user-images.githubusercontent.com/49638680/92366024-e7ced600-f0f4-11ea-82fd-b85d5c23c727.png)

Roughly speaking, the operating system main role is to schedule processes execution.
In order to enable concurrency it is crucial to have a schedule of execution of the different instructions of different threads and processes. This operation is called _interleaving_.

3. What is the “context” that is referred to in the term “context switch”?

![](https://user-images.githubusercontent.com/49638680/92366029-e9000300-f0f4-11ea-8655-44166e753eca.png)

Each process comes with a _context_, _i.e._ a set of memory and register values unique to that process, describing the state of the process at that time.
_Context switch_ is the operation of storing state of a process in memory, in order to stop the executing process and recover the to-be-executed process state.

4. What is the difference between a thread and a process?

![](https://user-images.githubusercontent.com/49638680/92366030-e9989980-f0f4-11ea-8539-dc04c52890c3.png)

A thread is the unit of execution within a process. A process can have anywhere from just one thread to many threads. Thus the unique context of a process contains elements that can be shared by its threads (memory values, memory addresses).

Indeed, processes are typically independent of each other, while threads exist as the subset of a process. Threads can communicate with each other more easily than processes can, however, and for this reason threads are more vulnerable to problems caused by other threads in the same process.

5. What is the main function of the Go runtime scheduler?

![](https://user-images.githubusercontent.com/49638680/92366033-e9989980-f0f4-11ea-8353-d9d4b8b7dba6.png)

Goroutines are the way we handle concurrency in Go.
Typically the Go runtime scheduler is executed in a single operating system thread.
It manages the interleaves of goroutines instructions inside this single thread.

6. Suppose that there are two goroutines executing, g1 and g2. Assume that g1 requires 1 second to complete when it is executed alone, and g2 requires 2 seconds to complete when it is executed alone. Assume that there is no synchronization between g1 and g2. Assume that g1 and g2 are executed concurrently and that g1 begins executing first. What do we know about the relative completion times of g1 and g2?

![](https://user-images.githubusercontent.com/49638680/92366035-ea313000-f0f4-11ea-9e72-e16e3e4778e5.png)

Simply the interleaves is left to Go runtime scheduler and it depends on variable initial states. Thus, we know nothing about the completion times of the two routines.

7. Assume that two goroutines are executed concurrently. Which of the following statements is true?

![](https://user-images.githubusercontent.com/49638680/92366039-eac9c680-f0f4-11ea-8613-e000c0e39303.png)

As the Go runtime scheduler handles interleaves differently everytime the program is executed, the relative order of the instructions changes everytime.
