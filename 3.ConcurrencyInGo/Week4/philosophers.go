/*
Write a program to implement the dining philosopher’s problem with the following
constraints/modifications.

1. There should be 5 philosophers sharing chopsticks, with one chopstick between
each adjacent pair of philosophers.
2. Each philosopher should eat only 3 times (not in an infinite loop as we did
in lecture).
3. The philosophers pick up the chopsticks in any order, not lowest-numbered
first (which we did in lecture).
4. In order to eat, a philosopher must get permission from a host which executes
in its own goroutine.
5. The host allows no more than 2 philosophers to eat concurrently.
6. Each philosopher is numbered, 1 through 5.
7. When a philosopher starts eating (after it has obtained necessary locks) it
prints “starting to eat <number>” on a line by itself, where <number> is the
number of the philosopher.
8. When a philosopher finishes eating (before it has released its locks) it
prints “finishing eating <number>” on a line by itself, where <number> is the
number of the philosopher.

A detailed explanation of the problem can be found at this link
https://en.wikipedia.org/wiki/Dining_philosophers_problem.

Here we implement the so-called arbitrator solution of the problem, introducing
a central authority (the host) to determine the chopstick picking order.
*/

/*

Disclaimer: This version might not be the required solution for the
course assignment.

Created by Oscar de Felice on 09/09/2020.
Copyright © 2020 Oscar de Felice.

<oscar.defelice@gmail.com>

*/

// Script to implement the dining philosopher's problem.
// We make use of a philosopher structure, and its eat method is called by a
// goroutine.
// The host is implemented through a channel structure.
// We implemented all as parametric as possible, so we only need to give a list
// of names and change a set of global variables to change the number and the
// identity of philosophers.
package main
// Compulsory package,
// the only one generating an executable

import (
  "fmt"       // Format library, including I/O methods
  "sync"      // Library providing basic synchronization primitives
  "time"      // Package providing functionality for handling time.
)

// Host structure
// The host is the one giving permission to eat.
// In this implementation is the arbitrator choosing the picking order of
// chopsticks.
// It has two attributes: a channel to communicate with philosophers and the
// channel buffer.
// It has just one method: Authorise.
type Host struct {
  channel   chan *Philosopher
  chan_buff int
}

// Authorise method
// Host method to determine the chopsticks picking order.
// The function transmit on the channel waiting for some philosopher to receive.
func (host *Host) Authorise(wg *sync.WaitGroup) {
	for {
		if len(host.channel) == host.chan_buff {
			<- host.channel
			<- host.channel

			// Add delay
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// Chopstick structure
// It is a wrapper for a mutex, since picking a chopstick
// would cause some semaphore locking.
type Chopstick struct {
  sync.Mutex
}

// Philosopher structure
// It contains five attributes, the name (only for aesthetic purposes), the id
// (position at table), left and right chopsticks and a counter, keepnig track
// of the number of times the philosopher has already eaten.
type Philosopher struct {
  name            string
  id              int
  leftCS, rightCS *Chopstick
  eatCount        int
}

// Eat method for philosophers
// It is written such that it has to be called in a goroutine.
// First, it checks whether the host gave permission to eat.
// It makes the philosopher grab the two chopsticks and increases the counter of
// eaten times.
func (philosopher *Philosopher) Eat (host *Host, wg *sync.WaitGroup) {
  for i := 0; i < 3; i++ {
		host.channel <- philosopher

		if philosopher.eatCount < 3 {
			philosopher.leftCS.Lock()
			philosopher.rightCS.Lock()

			fmt.Printf("%v starting to eat \n", philosopher.name)
			philosopher.eatCount++ // Increase counter
			fmt.Printf("%v finishing eating \n", philosopher.name)

			philosopher.leftCS.Unlock()
			philosopher.rightCS.Unlock()

			wg.Done()
		}
  }
}

// A collection of philosophers names
var names []string = []string{"Kant", "Heidegger", "Wittgenstein",
		"Locke", "Descartes", "Newton", "Hume", "Leibniz"}

// The host channel buffer
var chan_buff int = 2

// The max number of times philosophers are allowed to eat
var max_meals int = 3

// Script to model the philosophers' dining problem.
// Here, we implemented the arbitrator solution, with a central authority,
// the host and five hungry philosophers.
func main () {
  var wg sync.WaitGroup // WaitGroup object, to wait for routines to complete
                        // tasks
  var host Host
  host.chan_buff = chan_buff
  host.channel = make(chan *Philosopher, chan_buff)

  chopsticks := make([]*Chopstick, len(names))
  for i := 0; i < len(names); i++ {
		chopsticks[i] = new(Chopstick)
	}

  philosophers := make([]*Philosopher, len(names))
  for i := 0; i < len(names); i++ {
		philosophers[i] = &Philosopher{ names[i], i + 1, chopsticks[i],
                                    chopsticks[(i+1)%len(names)], 0}
	}
  fmt.Printf("There are %v philosophers sitting at a table.\n",
    len(philosophers))

  // Add semaphore counter values.
  wg.Add(max_meals*len(philosophers))

  // Authorisation of the host
  go host.Authorise(&wg)

  // Each Philosopher checks the authorisation and eats if he is allowed to.
  for i := 0; i < len(philosophers); i++ {
		go philosophers[i].Eat(&host, &wg)
	}

  // Be sure the main goroutine waits the others to complete.
  wg.Wait()

}
