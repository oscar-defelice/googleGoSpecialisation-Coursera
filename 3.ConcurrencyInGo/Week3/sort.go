/*
Write a program to sort an array of integers.
The program should partition the array into 4 parts, each of which is sorted by
a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large
sorted array.
The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that
it will sort.
When sorting is complete, the main goroutine should print the entire sorted
list.
*/

/*

Disclaimer: This version might not be the required solution for the
course assignment.

Created by Oscar de Felice on 08/09/2020.
Copyright © 2020 Oscar de Felice.

<oscar.defelice@gmail.com>

*/

// Script to sort an array of integers.
// We split the array in four partitions and sort each partition in parallel.
//
// If executing the applicaton multiple times with the same sequence of numbers
// (with more than 4 numbers) observe how each goroutine would take the same
// partition and on each running they would be run in different order
// (their order is undeterminstic and depends on Go runtime scheduler) and their
// sequences would be print in different order.
//
// The splitting is done sequentially. We could have splitted the partitions
// in a random way.

package main

// Compulsory package,
// the only one generating an executable

import (
	"bufio"     // Library to implement buffered I/O"
	"fmt"       // Format library, including I/O methods
	"log"       // Library to handle logging messages
	"math/rand" // Package to implement pseudo-random number generators.
	"os"        // Interface to operating system functionality
	"strconv"   // Conversion type from/to strings
	"strings"   // Library to manipulate strings
	"sync"      // Library providing basic synchronization primitives
)

// ReadValues function
// It reads integer values from console and store them into a slice.
// Returns an error if something went wrong.
func ReadValues() (numbers []int, err error) {
	fmt.Println("Please input numbers(separate with space): ")
	br := bufio.NewReader(os.Stdin)
	a, _, err := br.ReadLine() // ReadLine returns
	// the string containing the line,
	// the isPrefix attribute will be false when
	// returning the last fragment of the line.
	// An error occurs if something goes wrong.
	ns := strings.Split(string(a), " ")
	for _, s := range ns {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	return
}

// QuickSort function.
// It takes a slice as parameter and transform it in its sorted version.
// Defined as a recursive function.
func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1 // Start from the extrema

	pivot := rand.Int() % len(a) // Choose a random pivot index

	a[pivot], a[right] = a[right], a[pivot] // Put the pivot element
	// in the right place

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSort(a[:left])
	QuickSort(a[left+1:])
}

// Sort function
// It is a wrapper of the QuickSort sorting function to sort integers.
// Takes a WaitGroup as argument, since it will be called in a goroutine.
func Sort(numbers []int, wg *sync.WaitGroup) {
	QuickSort(numbers)

	if wg != nil {
		wg.Done()
	}
}

// CreatePartitions function
// It takes a slice and an int k as arguments and returns k partitions of the
// original slice, stored in a slice of slices.
func CreatePartitions(numbers []int, k int) (partitions [][]int) {
	var n = len(numbers)
	partitions = make([][]int, k)

	lenPartition := n / k

	for j := range partitions {
		partitions[j] = numbers[j*lenPartition : (j+1)*lenPartition]
	}

	return
}

// MergePartitions function
// It takes two sorted slices as arguments and returns a sorted merge of
// the two.
func MergePartitions(s1 []int, s2 []int) (merged []int) {
	var n1, n2 int = len(s1), len(s2)

	if n1 < n2 {
		s1, s2 = s2, s1 // Ensure to have s1 always as the longer slice.
		n1, n2 = n2, n1
	}

	merged = make([]int, n1+n2)

	i, j, k := 0, 0, 0
	for {
		if s1[i] < s2[j] {
			merged[k] = s1[i]
			i++
			k++
		} else if s1[i] == s2[j] {
			merged[k] = s1[i]
			i++
			k++

			merged[k] = s2[j]
			j++
			k++
		} else if s1[i] > s2[j] {
			merged[k] = s2[j]
			j++
			k++
		}

		if j == n2 {
			copy(merged[k:], s1[i:])
			k = n1 + n2 // exit condition
		}

		if i == n1 {
			copy(merged[k:], s2[j:])
			k = n1 + n2 // exit condition
		}

		if k == (n1 + n2) {
			break // exit condition
		}
		fmt.Println(merged)
	}

	return
}

// Merge function
// It calls recursively the MergePartitions on couples of slices and store the
// result in one single slice.
func Merge(partitions [][]int, k int) (sorted []int) {

	var psorted [][]int

	for i := 0; i <= k/2; i += 2 {
		psorted = append(psorted, MergePartitions(partitions[i], partitions[i+1]))
	}

	if len(psorted) == 2 {
		sorted = MergePartitions(psorted[0], psorted[1])
	} else {
		Merge(psorted, k)
	}
	return
}

// main goroutine.
// The script sorts a list of integers.
// It acts through four subroutines in parallel, splitting the list into four
// parts, sorting them singularly and merging them only at the end.
func main() {
	var k = 4             // number of partitions and of concurrent subroutines.
	var wg sync.WaitGroup // WaitGroup object, to wait for routines to complete
	// tasks

	// Read numbers from console
	if numbers, err := ReadValues(); err != nil {
		log.Fatal("Error reading text from console: ", err)
	} else {
		nTotal := len(numbers)

		// With k numbers or less, we do not parallelise the computing.
		if nTotal <= k {
			QuickSort(numbers)
			fmt.Printf("Sorted numbers: %v", numbers)
		} else {
			// Create partitions.
			partitions := CreatePartitions(numbers, k)

			// add the number of partitions to the semaphore counter.
			wg.Add(k)
			for i := 0; i < k; i++ {
				go Sort(partitions[i], &wg)
			}
			wg.Wait()

			sorted := Merge(partitions, k)

			fmt.Printf("Sorted numbers: %v", sorted)
		}
	}
}
