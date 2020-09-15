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

/*

Disclaimer: This version is NOT the required solution for the
course assignment.

I just wrote it to exercise myself on a slightly more complicated
problem, involving multiple functions and libraries.

Created by Oscar de Felice on 09/04/2020.
Copyright © 2020 Oscar de Felice.

<oscar.defelice@gmail.com>

*/

package main

import (
	"bufio"   // Library to implement buffered I/O
	"fmt"     // Format library, including I/O methods
	"os"      // Interface to operating system functionality
	"strings" // Library to manipulate strings
)

// RemoveSpaceReturn returns a lowercase copy of the input string
// with removed spaces and '\n' special character.
func RemoveSpaceReturn(s string) string {
	var space, trailing string = " ", "\n"

	str := strings.ToLower(s) // transform the string all in lowercase

	str = strings.ReplaceAll(str, space, "")    // remove spaces
	str = strings.ReplaceAll(str, trailing, "") // remove new line

	return str
}

// IsRight returns true if string starts with character "i", has somewhere "a"
// and ends with "n" while ignoring spaces, the trailing invisible CR
// and/or LF characters.
func IsRight(s string) bool {
	var isRight bool
	var prefix, suffix string = "i", "n"
	var inner = "a"

	str := RemoveSpaceReturn(s)

	condPre := strings.HasPrefix(str, prefix)
	condSuf := strings.HasSuffix(str, suffix)
	condCon := strings.Contains(str, inner)

	if condPre && condSuf && condCon {
		isRight = true
	} else {
		isRight = false
	}

	return isRight

}

func main() {
	fmt.Println("Please enter a string: ")

	reader := bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Sorry, invalid input. Error: ", err)
	} else {
		if IsRight(inputString) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}
