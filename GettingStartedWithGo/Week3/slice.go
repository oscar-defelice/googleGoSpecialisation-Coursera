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

package main
// Compulsory package,
// the only one generating an executable

import(
  "bufio"    // Library to implement buffered I/O
  "fmt"      // Format library, including I/O methods
  "os"       // Interface to operating system functionality
  "strconv"  // Conversion type from/to strings
  "strings"  // Library to manipulate strings
 )

// Function to print the slice on screen.
// It prints also the capacity and lentgth of the slice.
func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// The functions return a string as readed from the console
// input.
// It eventually returns an error, trough an error code.
func ReadTextFromConsole(message string) (string, error) {
  reader := bufio.NewReader(os.Stdin)

  fmt.Print(message)
  text, err := reader.ReadString('\n')

  return text, err
}

// RemoveSpecialReturn returns a copy of the input string
// with removed '\n' special character.
func RemoveSpecialReturn (s string) string {
  var trailing string = "\n"

  str := strings.ReplaceAll(s, trailing, "") // remove new line

  return str
}

// Insertion Sort algorithm implementation.
// We choose the insertion algorithm because
// it works well with small amount of data
// it works with online (continuous stream of) data
// the additional amount of memory at each step is always O(1).
// The function modifies the slice, hence there is no return statement.
func insertionSort(numbers []int) {
  for i := 1; i < len(numbers); i++ {
    tmp := numbers[i] // define a temporary variable
    j := i-1          // the sorting start from i=1, since i=0 already sorted.
    for (j >= 0) && (numbers[j] >= tmp) {
      numbers[j+1] = numbers[j] // insert the element in the right place
      j = j-1                   // swapping elements when necessary.
    }
    numbers[j+1] = tmp  // assign tmp to the next element
  }
}

// Insertion Sort algorithm implemented recursively.
// The function calls itself until the list is sorted.
// The function modifies the slice, hence there is no return statement.
func insertionSortR(numbers []int, n int){
  if n > 0 {
    insertionSortR(numbers, n-1)
    tmp := numbers[n]
    j := n-1
    for (j >= 0) && numbers[j] >= tmp {
      numbers[j+1] = numbers[j]
      j = j-1
    }
    tmp = numbers[j+1]
  }
}

// Script which prompts the user to enter integers and stores
// the integers in a sorted slice.
// The program prints on screen the sorted slice.
// It keeps going until the user enters X to exit.
func main() {
  slice := make([]int, 0, 3) // create an empty integer slice of size (length) 3
  var input string

  for input != "X" {
    if input, err := ReadTextFromConsole("Insert an integer or X to exit: "); err != nil {
  	   fmt.Println("Input error:", err, ".Please try again.")
  		 continue
  	} else {
  	   fmt.Println("Input string:", input)
       input = RemoveSpecialReturn(input)
  		 if input == "X" {
         fmt.Println("X entered, exit")
  		   break
  			}

        if number, err := strconv.Atoi(input); err != nil {
				      fmt.Println("Input string is not an integer number. Please try again.")
				      continue
			  } else {
				      fmt.Println("Input string is an integer number: ", number)
              slice = append(slice, number)
				      insertionSort(slice)
				      fmt.Println("Slice in the (increasing) order:", slice)
			    }
      }
    }
}
