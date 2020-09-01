/*
Let us assume the following formula for displacement s as a function of time t,
acceleration a, initial velocity vo, and initial displacement so.

s = 1/2 a t^2 + v0 t + s0

This program first prompts the user to enter values for acceleration,
initial velocity, and initial displacement.
Then it prompts the user to enter a value for time and computes
the displacement after the entered time.

We define and use a function called GenDisplaceFn() which takes three float64
arguments:
  acceleration a,
  initial velocity v0,
  initial displacement s0.

GenDisplaceFn() returns a function which computes displacement as
a function of time, assuming the given values acceleration, initial velocity,
and initial displacement.
The function returned by GenDisplaceFn() takes one float64 argument t, r
epresenting time, and return one float64 argument which is the displacement
travelled after time t.

Example:
assume the following initial values:
  acceleration: a = 10,
  initial velocity: v0 = 2,
  initial displacement: s0 = 1.

The following statement can be used to call GenDisplaceFn()
to generate a function fn which will compute displacement as a function of time.

  fn := GenDisplaceFn(10, 2, 1)

Then to print the displacement after 3 seconds.

  fmt.Println(fn(3))

And hence the following statement to print the displacement after 5 seconds.

  fmt.Println(fn(5))
*/

/*

Disclaimer: This version might be the required solution for the
course assignement.

Created by Oscar de Felice on 01/09/2020.
Copyright © 2020 Oscar de Felice.

<oscar.defelice@gmail.com>

*/

package main
// Compulsory package,
// the only one generating an executable

import (
  "bufio"   // Library to implement buffered I/O
  "fmt"     // Format library, including I/O methods
  "math"    // Package to access mathematical operations
  "os"      // Interface to operating system functionality
  "strconv" // Conversion type from/to strings
)

// GenDisplaceFn function.
// A function that returns a function which computes displacement as
// a function of time, assuming the given values acceleration, initial velocity,
// and initial displacement.
// The function returned by GenDisplaceFn() takes one float64 argument t,
// representing time, and return one float64 argument which is the displacement
// travelled after time t.
func GenDisplaceFn (a, v0, s0 float64) func (float64) float64 {
  f := func (t float64) (s float64) {
    s = .5*a*math.Pow(t, 2) + v0*t + s0
    return
  }
  return f
}

// ReadValue function.
// Generic function returning the read value and an error.
func ReadValue(var_name string) (value float64) {

  var temp string
  var valid int

  read_input := bufio.NewScanner(os.Stdin)

  temp = var_name
  valid = 0
  for valid == 0 {
      fmt.Printf("\nPlease enter %s as a floating point: ", var_name)
      read_input.Scan()
      temp = read_input.Text()
      value, err := strconv.ParseFloat(temp, 64)
      if err == nil{
          valid = 1
          return value
      } else {
          fmt.Println("Invalid input. Floating point number required.")
      }
  }
  return value
}

// Script which first prompts the user to enter values for acceleration,
// initial velocity, and initial displacement.
// Then it prompts the user to enter a value for time and computes
// the displacement after the entered time.
func main () {

  var a, v0, s0, t float64

  // Read values from command line
  a = ReadValue("acceleration")
  v0 = ReadValue("initial velocity")
  s0 = ReadValue("initial displacement")
  t = ReadValue("time interval")

  f := GenDisplaceFn(a, v0, s0)

  fmt.Printf("\nDisplacement after %[1]v seconds is %[2]v.", t, f(t)) 

}
