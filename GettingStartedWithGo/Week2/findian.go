/*
Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’,
‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the
character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter if the
characters are upper-case or lower-case.
*/

package main

import(
    "fmt"
    "strings"
  )

func main() {
	// Script to read a string from command line.
	// It searches for the substrings 'i', 'a', 'n'
  // The script prints 'Found!' if the strings
  // starts with 'i', ends with 'n' and contains 'a',
  // otherwise it prints 'Not Found!'
  // The script is not case-sensitive.
	var str string
  var prefix, suffix string = "i", "n"
  var inner string = "a"

	fmt.Println("Please, enter a string: ")
	fmt.Scan(&str)

  str = strings.ToLower(str) // transform the string all in lowercase.

  condPre := strings.HasPrefix(str, prefix)
  condSuf := strings.HasSuffix(str, suffix)
  condCon := strings.Contains(str, inner)

  if condPre && condSuf && condCon {
    fmt.Println("Found!")
  } else {
    fmt.Println("Not Found!")
  }
}
