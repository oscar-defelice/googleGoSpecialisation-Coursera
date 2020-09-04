/*
Write a program which prompts the user to enter integers and stores
the integers in a sorted slice.
The program should be written as a loop.
Before entering the loop, the program should create an empty integer slice
of size (length) 3.
During each pass through the loop, the program prompts the user to enter
an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints
the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers
which the user decides to enter.
The program should only quit (exiting the loop) when the user enters
the character ‘X’ instead of an integer.
*/

/*

Disclaimer: This version might be the required solution for the
course assignment.

Created by Oscar de Felice on 08/04/2020.
Copyright © 2020 Oscar de Felice.

<oscar.defelice@gmail.com>

*/

package main
// Compulsory package,
// the only one generating an executable

import (
	"fmt"     // Format library, including I/O methods
	"sort"    // Sort library provides primitives for sorting slices.
	"strconv" // Conversion type from/to strings
)

// Script which prompts the user to enter integers and stores
// the integers in a sorted slice.
// The program prints on screen the sorted slice.
// It keeps going until the user enters X to exit.
func main() {
	storage := make([]int, 3)  // create an empty integer slice of size (length) 3
	var input string
	counter := 0
	for {
		fmt.Printf("Enter integer: ")
		fmt.Scan(&input)

		if input == "X" {
			break
		}

		if value, err := strconv.Atoi(input); err == nil {
			if counter < 3 {
				storage[counter] = value
			} else {
				storage = append(storage, value)
			}
			sorted := make([]int, len(storage))
			copy(sorted, storage)
			sort.Ints(sorted) // sort the slice
			fmt.Println(sorted)
			counter++
		}
	}
}
