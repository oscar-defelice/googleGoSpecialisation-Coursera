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
// We make the program print two strings one character per time.
// We will have a final shuffle of the two strings.
//
// To test this, one can run this application in loop.
// E.g. use the following command in terminal (bash):
// for i in (1, 1, 100) do go run racecondition.go
//
// My test were showing it would print different shuffles of the two strings.
// This randomness in output proves there is a race condition.

package main

// Compulsory package,
// the only one generating an executable

import (
	"fmt"  // Format library, including I/O methods
	"time" // Package providing functionality for handling time.
)

// f function
// The function prints elements of the strings separated by blank spaces.
// Since we plan to use the function in a goroutine, we added a sleep time
// to prevent program termination before goroutines return.
func f(s string) {
	for _, c := range s {
		fmt.Print(string(c), " ")
		time.Sleep(10 * time.Millisecond)
	}
}

// Script containing two goroutines.
// One goroutine prints elements of one string.
// The other one prints another strings.
// This to show we will have a shuffle of the two strings.
func main() {
	// run two different goroutine
	go f("Hello")
	go f("World")

	// prevent program termination before goroutines return
	time.Sleep(1 * time.Second)
}
