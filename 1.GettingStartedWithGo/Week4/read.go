/*
Write a program which reads information from a file
and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name,
in that order, separated by a single space on the line.

Your program will define a name struct which has two fields,
fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file
and create a struct which contains the first and last names found
in the file.
Each struct created will be added to a slice, and after all lines
have been read from the file, your program will have a slice containing
one struct for each line in the file. After reading all lines from the file
 your program should iterate through your slice of structs
 and print the first and last names found in each struct.
*/

/*

Disclaimer: This version might be the required solution for the
course assignment.

Created by Oscar de Felice on 23/04/2020.
Copyright © 2020 Oscar de Felice.

<oscar.defelice@gmail.com>

*/

package main

// Compulsory package,
// the only one generating an executable

import (
	"bufio"   // Library to implement buffered I/O
	"fmt"     // Format library, including I/O methods
	"os"      // Interface to operating system functionality
	"regexp"  // Regular expressions library
	"strconv" // Conversion type from/to strings
	"strings" // Library to manipulate strings
)

// Define a name structure.
// The names read from file will be stored in a
// slice of name structures.
type name struct {
	fname string
	lname string
}

// integer number. There is a check on the number.
// if the value is not an integer this will cause an error.
var (
	integerNumberREGEX = regexp.MustCompile(`^(\d+)\.?(\d*)$`)
)

// ReadFromConsole function.
// Generic function taking as argument a message to show on screen
// and returning the read input and an error.
func ReadFromConsole(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)
	text, err := reader.ReadString('\n')

	return text, err
}

// inputFilename function
// It calls a reader and memorise input filename into
// a string, which is returned.
// Using buffer reader allows both name and
// address to contain SPACE characters.
func inputFilename() (filename string) {

	for {
		if name, err := ReadFromConsole("Please enter the filename: "); err != nil {
			fmt.Println("\nError: ", err)
		} else {
			filename = RemoveSpecialReturn(name)
			break
		}
	}

	return
}

func inputMaxLen() (int, error) {

	for {

		if text, err := ReadFromConsole("Please enter the max lenght for names: "); err != nil {
			fmt.Println("\nError: ", err)
		} else {

			maxLenStr := RemoveSpecialReturn(text)
			match := integerNumberREGEX.FindStringSubmatch(maxLenStr)

			if len(match) == 0 { // check on the lenght of the parsed number
				return 0, fmt.Errorf("Error parsing: %s is not an integer", maxLenStr)
			} else {
				maxLen, err := strconv.Atoi(match[1])
				return maxLen, err
			}
		}
	}
}

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

// Cut the string if the name is too long
func fixLongName(buffer string, maxLen int) (name string) {

	if len(buffer) > maxLen {
		name := string(buffer[0:maxLen])
	} else {
		name := buffer
	}
	return
}

// runeToName function.
// It takes first name, last name and a maxLen parameter and store them
// in a name structure.
// If first name and last name are longer then maxLen, they
// are truncated.
func runeToName(fname string, lname string, maxLen int) name {
	var nameFromLine name

	nameFromLine.fname = fixLongName(fname, maxLen)
	nameFromLine.lname = fixLongName(lname, maxLen)

	return nameFromLine
}

// ReadAndStore function
// This reads data from file and returns the slice of name structures.
func ReadAndStore(filename string, maxLen int) (names []name) {

	if file, err := os.Open(filename); err != nil {
		fmt.Println("Error while opening the file: ", err)
		os.Exit(1)
	} else {
		names = make([]name, 0)
		reader := bufio.NewReader(file)
		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			line := scanner.Text()
			words := strings.Split(line, " ")
			firstName := words[0]
			lastName := words[1]

			nameFromLine := runeToName(firstName, lastName, maxLen)

			names = append(names, nameFromLine)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

	}
	return names
}

func main() {
	filename := inputFilename()
	maxLen, err := inputMaxLen()

	names := ReadAndStore(filename, maxLen)

	if err != nil {
		fmt.Println("Error while reading maxLen: ", err)
	} else {
		fmt.Printf("\nNames slice (maxLen of %d characters)\n", maxLen)
	}

	for index := 0; index < len(names); index++ {
		fmt.Println("Name:", string(names[index].fname[:]))
		fmt.Println("Surname:", string(names[index].lname[:]))
		fmt.Println()
	}
}
