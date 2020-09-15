/*
Write a program which prompts the user to enter a floating point
number and prints the integer which is a truncated version of the
floating point number that was entered.

Truncation is the process of removing the digits to the right of
the decimal place.

Example:
trunc(10.9) = 10
*/

package main

// Compulsory package,
// the only one generating an executable

import (
	"fmt" // Format library, including I/O methods
)

// Script to read a floating point number,
// to calculate its truncation and print it on screen.
func main() {

	var number int

	fmt.Println("Please, enter a floating point number: ")
	fmt.Scan(&number) // number is declared to be int. Truncation will work automatically
	fmt.Println("Truncated integer: ", number)

}
