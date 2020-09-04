/*
Write a program which prompts the user to enter a floating point
number and prints the integer which is a truncated version of the
floating point number that was entered.

Truncation is the process of removing the digits to the right of
the decimal place.

Example:
trunc(10.9) = 10
*/

/*

Disclaimer: This version is NOT the required solution for the
course assignment.

I just wrote it to exercise myself on a slightly more complicated
problem, involving multiple functions and libraries.

Created by Oscar de Felice on 08/04/2020.
Copyright © 2020 Oscar de Felice.

<oscar.defelice@gmail.com>

*/

package main
// Compulsory package,
// the only one generating an executable

import (
	"bufio"	  // Library to implement buffered I/O
  "fmt" 		// Format library, including I/O methods
  "log"     // Library to handle logging messages
  "os"      // Interface to operating system functionality
  "regexp"  // Regular expressions library
  "strconv" // Conversion type from/to strings
  "strings" // Library to manipulate strings
)

// The functions return a string as readed from the console
// input.
// It eventually returns an error, trough an error code.
func ReadTextFromConsole(message string) (string, error) {
  reader := bufio.NewReader(os.Stdin)

  fmt.Print(message)
  text, err := reader.ReadString('\n')

  return text, err
}

var (
    floatingNumberREGEX = regexp.MustCompile(`^(\d+)\.?(\d*)$`)
) // floating number. There is a check on the number.
  // if the value is not a float this will cause an error.

// Truncate. Parse a FloatingNumber from a string and get the integer part.
func Truncate(number string) (int64, error) {
	matched := floatingNumberREGEX.FindStringSubmatch(number)

	if len(matched) == 0 { // check on the lenght of the parsed number
			 return 0, fmt.Errorf("Error parsing: %s is not a Float", number)
	 }

	 return strconv.ParseInt(matched[1], 10, 64)
}

// Script to read a floating point number,
// to calculate its truncation and print it on screen.
func main() {
  var err error

	number, err := ReadTextFromConsole("floating point number: ")

	if err != nil {
  	log.Fatal("Error reading text from console: ", err)
  } else {
      var integer, err = Truncate(strings.Trim(number, "\n"))

      if err != nil {
          log.Fatal("Error parsing FloatingNumber: ", err)
      } else {
          fmt.Println(fmt.Sprintf("Truncated Integer: %d", integer))
      }
  }
}
