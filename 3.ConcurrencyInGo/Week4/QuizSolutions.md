## Week 4 Quiz

1. What line of code could be used to define a loop which iteratively reads from a channel `ch1`?

![](https://user-images.githubusercontent.com/49638680/92599875-1abaca80-f2ab-11ea-8772-b74b2cd75c33.png)

`range` is a construct used to iterate over basic data structures. We can also use this syntax to iterate over values received from a channel.
The `range` in the answer iterates over each element as it’s received from `ch1`.

2. What does the `select` keyword do?

![](https://user-images.githubusercontent.com/49638680/92599878-1b536100-f2ab-11ea-9183-028619eed185.png)

The `select` statement lets a goroutine wait on multiple communication operations.
A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
One condition, fon instance, can be receiving from a channel.

3. What is the meaning of the `default` clause inside a `select`?

![](https://user-images.githubusercontent.com/49638680/92599879-1bebf780-f2ab-11ea-85c1-af1f15f13eac.png)

The default case in a `select` is run if no other case is ready. Hence, if all other cases are blocked, the default gets executed.

4. Suppose that there are two goroutines, `g1` and `g2`, which share a variable `x`. `x` is initialized to `0`. The only instruction executed by `g1` is `x = 4`. The only instruction executed by `g2` is `x = x + 1`. What is a possible value for `x` after both goroutines are complete?

![](https://user-images.githubusercontent.com/49638680/92599880-1bebf780-f2ab-11ea-862b-f0b12a594132.png)

In order to answer this question, we need to recall that interleavings happen at machine code level, not at source code one.
Hence, it might happen that `x = 4` block of instructions happens after the _read the value of `x` from memory_ part of `x = x + 1`, hence in this case `x` would be equal `1` eventually.
Other cases are simply `g1` executing first (hence `x = 5`), or `g2` executing first (hence `x = 4`).

5. What is _mutual exclusion_?

![](https://user-images.githubusercontent.com/49638680/92599881-1c848e00-f2ab-11ea-8f28-5666cfc14726.png)

Mutual execution is a way to avoid conflicts on global variables and so to avoid race condition.
Indeed, with such construct, one can define a block of code to be executed in mutual exclusion, _i.e._ only one goroutine accesses a variable at a time.

6. What is true about deadlock?

![](https://user-images.githubusercontent.com/49638680/92599882-1c848e00-f2ab-11ea-9dbf-926c305eeac5.png)

Deadlock is a state in which each member of a group is waiting for another member, including itself, to take action, such as sending a message or more commonly releasing a lock.
Roughly speaking, this happens when there are mutually exclusive resources (hold by some process), and the unlock condition of such resources depends on other processes that in turn are dependent on the initial process.

7. What is the method of the sync.mutex type which must be called at the beginning of a critical region?

![](https://user-images.githubusercontent.com/49638680/92599883-1c848e00-f2ab-11ea-8bfa-b5d144f8feb1.png)

In Go, a Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point of time to prevent race condition from happening.
There are two methods defined on Mutex namely `Lock` and `Unlock`. Any code that is present between a call to Lock and Unlock will be executed by only one Goroutine, thus avoiding race condition.

_e.g._
```go
mutex.Lock()  
x = x + 1  
mutex.Unlock()
```
