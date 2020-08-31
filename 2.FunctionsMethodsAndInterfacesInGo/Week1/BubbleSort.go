/*
Write a program sorting integers using Bubble Sorting algorithm.
In order to implement the algoritmh, you have to create a routine
that sorts a series of ten integers using the bubble sort algorithm.
*/

/*

Disclaimer: This version might not be the required solution for the
course assignement.

Created by Oscar de Felice on 31/08/2020.
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
  "strconv" // Conversion type from/to strings
  "strings" // Library to manipulate strings
)

// Script which prompts the user to enter integers and stores
// the integers in a slice.
// The program sorts slice and prints on screen the sorted slice.
func main () {
  fmt.Println("Please input numbers(separate with space):")
  br := bufio.NewReader(os.Stdin)
  a, _, _ := br.ReadLine() // ReadLine returns
                           // the str containing the line,
                           // the isPrefix attribute will be false when
                           // returning the last fragment of the line.
                           // an error.
  ns := strings.Split(string(a), " ")
  var values []int
  for _, s := range(ns) {
    n, _ := strconv.Atoi(s)
    values = append(values, n)
  }
  BubbleSort(values)
  fmt.Println("Sorted list of numbers:")
  fmt.Println(values)
}

// BubbleSort function.
// It implements the Bubble Sort algorithm.
// It takes a slice as parameter and transform it in its sorted version.
func BubbleSort (a []int) {
  for i := 0; i < len(a); i++ {
    for j := 0; j < len(a) - i - 1; j++ {
      if a[j+1] < a[j] {
        Swap(a, j)
      }
    }
  }
}

func Swap (a []int, j int) {
  a[j], a[j+1] = a[j+1], a[j]
}
