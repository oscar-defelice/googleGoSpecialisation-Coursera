/*
Write a program which prompts the user to first enter a name,
and then enter an address.
Your program should create a map and add the name and address
to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object
from the map, and then your program should print the JSON object.
*/

/*

Disclaimer: This version might be the required solution for the
course assignment.

Created by Oscar de Felice on 22/04/2020.
Copyright © 2020 Oscar de Felice.

<oscar.defelice@gmail.com>

*/

package main

// Compulsory package,
// the only one generating an executable

import (
	"bufio"         // Library to implement buffered I/O
	"encoding/json" // Package to implement encoding and decoding of JSON
	"fmt"           // Format library, including I/O methods
	"os"            // Interface to operating system functionality
	"strings"       // Library to manipulate strings
)

// Create the Person variable of type map[string]string
// To allocate "name" and "address".

// Person is a map with two keys "name" and "address"
type Person map[string]string

// RemoveSpecialReturn returns a copy of the input string
// with removed '\n' or '\r' special characters.
func RemoveSpecialReturn(s string) string {
	var newline = "\n"
	var cabret = "\r"

	if strings.HasSuffix(s, newline) {
		s = strings.TrimSuffix(s, newline)
	}
	if strings.HasSuffix(s, cabret) {
		s = strings.TrimSuffix(s, cabret)
	}

	return s
}

// inputPersonData function
// it takes care of the read part of the script.
// It calls a reader and memorise input variables into
// a person type, which is returned.
// Using buffer reader allows both name and
// address to contain SPACE characters.
func inputPersonData() Person {
	person := make(Person)
	stdinReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please enter the name:")
		if name, err := stdinReader.ReadString('\n'); err != nil {
			fmt.Println("\nError: ", err)
		} else {
			person["name"] = RemoveSpecialReturn(name)
			break
		}
	}

	for {
		fmt.Print("Please enter the address:")
		if address, err := stdinReader.ReadString('\n'); err != nil {
			fmt.Println("\nError: ", err)
		} else {
			person["address"] = RemoveSpecialReturn(address)
			break
		}
	}

	return person
}

// marshJSON function
// it takes care of the marshal part of the script.
// Takes data as argument and store them in a JSON.
// Finally the JSON is printed.
// no return statement.
func marshJSON(person Person) {
	if data, err := json.Marshal(person); err != nil {
		fmt.Println("\nError in serializing to JSON: ", err)
	} else {
		fmt.Println("\nJSON:", string(data))
	}
}

// main function.
// It calls the inputPersonData function to read
// data and hence store them in a JSON.
// It calls the marshJSON function to store and print data.
func main() {
	person := inputPersonData()

	marshJSON(person)
}
