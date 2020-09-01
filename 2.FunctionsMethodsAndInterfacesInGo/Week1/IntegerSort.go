/*
Write a program sorting integers using a sorting algorithm.
In order to implement the algoritmh, you have to create a routine
that sorts a series of ten integers using the one of the implemented sorting
algorithms.
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
  "bufio"     // Library to implement buffered I/O
  "fmt"       // Format library, including I/O methods
  "math/rand" // Package to implement pseudo-random number generators.
  "os"        // Interface to operating system functionality
  "strconv"   // Conversion type from/to strings
  "strings"   // Library to manipulate strings
)

// The implemented algorithms are collected in a map[string]func([]int)
var (
  sortingAlgs = map[string]func([]int) {
    "BubbleSort": BubbleSort,
    "InsertionSort": InsertionSort,
    "QuickSort": QuickSort,
  }
)

// Swap function.
// It simply swaps the two elements in the slice at index j.
func Swap (a []int, j int) {
  a[j], a[j+1] = a[j+1], a[j]
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

// InsertionSort function.
// It takes a slice as parameter and transform it in its sorted version.
func InsertionSort (a []int) {
  for i := 0; i < len(a); i++ {
    j:=i
    for j>0 {
      if a[j-1] > a[j] {
        Swap(a, j-1)
      }
      j -= 1
    }
  }
}

// QuickSort function.
// It takes a slice as parameter and transform it in its sorted version.
// Defined as a recursive function.
func QuickSort (a []int) {
  if len(a) < 2 {
        return
    }
    left, right := 0, len(a)-1 // Start from the extrema

    pivot := rand.Int() % len(a) // Choose a random pivot index

    a[pivot], a[right] = a[right], a[pivot] // Put the pivot element
                                            // in the right place

    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }

    a[left], a[right] = a[right], a[left]

    QuickSort(a[:left])
    QuickSort(a[left+1:])
}

// Script which prompts the user to enter integers and stores
// the integers in a slice.
// The program sorts slice and prints on screen the sorted slice.
func main () {
  values, _ := ReadValues()

  fmt.Println("List of implemented algorithms: ")
  PrintKeys(sortingAlgs)
  algorithm := GetAlgorithm()

  algorithm(values)
  fmt.Println(algorithm)
  fmt.Println("Sorted list of numbers:")
  fmt.Println(values)
}

// ReadValues function.
// Generic function returning the read values and an error.
func ReadValues () (values []int, err error) {
  fmt.Println("Please input numbers(separate with space):")
  br := bufio.NewReader(os.Stdin)
  a, _, err := br.ReadLine() // ReadLine returns
                             // the str containing the line,
                             // the isPrefix attribute will be false when
                             // returning the last fragment of the line.
                             // an error.
  ns := strings.Split(string(a), " ")
  for _, s := range(ns) {
    n, _ := strconv.Atoi(s)
    values = append(values, n)
  }

  return values, err
}

// InputAlgorithm function.
// Generic function returning the string input by the user
// corresponding to the algorithm name and an error.
func InputAlgorithm () (string, error) {
  fmt.Println("Please input which sorting algorithm you want to use:")
  br := bufio.NewReader(os.Stdin)
  algorithm, _, err := br.ReadLine() // ReadLine returns
                                     // the str containing the line,
                                     // the isPrefix attribute is false when
                                     // returning the last fragment of the line.
                                     // Returnd eventually an error.

  return string(algorithm), err
}

// ChooseAlgorithm function.
// Generic function taking the sorting algorithm name as parameter
// and returning the corresponding function.
func ChooseAlgorithm(algorithm string) func([]int) {
  var result func([]int)
  if f, ok := sortingAlgs[algorithm]; ok != true {
			fmt.Printf(`Sorry, Algorithm not implemented!
        Choose an algorithm from the list.`)
			} else {
			result = f
			}
  return result
}

// GetAlgorithm function.
// Generic function reading the algorithm name
// and returning the corresponding function.
func GetAlgorithm() func([]int) {
  var function func([]int)
  if algorithm, err := InputAlgorithm(); err!= nil {
    fmt.Println("\nError: ", err)
  } else {
    function = ChooseAlgorithm(algorithm)
  }
  return function
}

// PrintKeys function.
// function to print keys of the map m given as argument.
func PrintKeys (m map[string]func([]int)) []string {
  keys := []string{}
  for k := range m {
    keys = append(keys, k)
  }
	fmt.Println(keys)
  return keys
}
