/*

Disclaimer: This version might be the required solution for the
course assignment.

Created by Oscar de Felice on 07/09/2020.
Copyright © 2020 Oscar de Felice.

<oscar.defelice@gmail.com>

*/

// main() executes two anonymous functions. Keyword 'go' makes them run
// concurrently, in goroutines.
// Go runtime scheduler schedules goroutines in a non-deterministic way so
// there are multiple possible interleavings - each time we run the application
// scheduler might change the order of their execution.
// Sometimes function which increments x will be executed first while sometimes
// the function which prints out its value.
// Because both functions in goroutines are communicating through variable x
// there is a chance of race condition - the order of accessing variable x
// is non-deterministic.
// If we run the application multiple times, it will sometimes print 0
// and sometimes 1.
//
// To test this, one can run this application in loop.
// Please use the following command in terminal (bash):
// for i in (1, 1, 100) do go run racecondition.go
//
// My test were showing it would print alternatively 1 or 0.
// This randomness in output proves there is a race condition.

package main
// Compulsory package,
// the only one generating an executable

import (
  "fmt"               // Format library, including I/O methods
  "time"              // Package providing functionality for handling time.
)

var x int

// Script containing two goroutines.
// One goroutine implements an increment of the variable.
// The other one prints the value of the variable on screen.
func main () {
  go func() {
    x += 1
  }()

  go func() {
    fmt.Println("Value of x: ", x)
  }()

  // prevent program termination before goroutines return
	time.Sleep(1 * time.Second)
}
