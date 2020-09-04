/*
Write a program which allows the user to get information about a predefined set
 of animals.
 Three animals are predefined
  cow,
  bird,
  snake.

Each animal can
  eat,
  move,
  speak.

The user can issue a request to find out one of three things about an animal:
  1) the food that it eats,
  2) its method of locomotion,
  3) the sound it makes when it speaks.

The following table contains the three animals and their associated data
which should be hard-coded into your program.

Animal  | Food eaten  |	Locomotion method |	Spoken sound
=========================================================
cow     |	grass	      | walk	            |  moo
bird	  | worms	      | fly	              |  peep
snake	  | mice	      | slither	          |  hsss

*/

/*

Disclaimer: This version might not be the required solution for the
course assignment.

Created by Oscar de Felice on 01/09/2020.
Copyright © 2020 Oscar de Felice.

<oscar.defelice@gmail.com>

*/

package main
// Compulsory package,
// the only one generating an executable

import (
  "bufio"             // Library to implement buffered I/O
  "errors"            // User defined erros package
  "fmt"               // Format library, including I/O methods
  "os"                // Interface to operating system functionality
  "strings"           // Library to manipulate strings
)

// Interface for animals
type animalInterface interface {
	Eat()
	Move()
	Speak()
}

// Structure collecting attributes about animals.
// The implemented attributes are food, locomotion and noise.
type Animal struct {
  name       string
  food       string
  locomotion string
  noise      string
}

// Eat function
// It prints on screen the animal food.
func (animal *Animal) Eat() {
  fmt.Printf("The %[1]s eats %[2]s", animal.name, animal.food)
}

// Move function
// It prints on screen the animal locomotion method.
func (animal *Animal) Move() {
  fmt.Printf("The %[1]s moves by %[2]s", animal.name, animal.locomotion)
}

// Speak function
// It prints on screen the animal noise.
func (animal *Animal) Speak() {
  fmt.Printf("The %[1]s noise is %[2]s", animal.name, animal.noise)
}

// Declare animal variables and a map
var cow, bird, snake Animal = Animal{"cow", "grass", "walk", "moo"},
                              Animal{"bird", "worms", "fly", "peep"},
                              Animal{"snake", "mice", "slither", "hsss"}

var animals = map[string]Animal{
  "cow"   : cow,
  "bird"  : bird,
  "snake" : snake,
}

// getRequest function
// It reads users inputs and returns chosen animal name and chose attribute.
func getRequest() (string, string) {
  var animalChoice, infoChoice string

  for {
    animal, info, err := ReadValues(); if err != nil {
      fmt.Println("Input error:", err, ".Please try again.")
      continue
    } else {
      animalChoice = string(animal)
      infoChoice = string(info)
      break
    }
  }
  return animalChoice, infoChoice
}

// ReadValues function.
// Generic function returning the read values and an error.
func ReadValues() (animalChoice string, infoChoice string, err error) {
  fmt.Println("\n\nPlease input animal name and required info: ")
  fmt.Println("Type 'new animal' to create a new one. ")
  fmt.Printf("\n>")
  br := bufio.NewReader(os.Stdin)
  text, _, err := br.ReadLine() // ReadLine returns
                                // the str containing the line,
                                // the isPrefix attribute will be false when
                                // returning the last fragment of the line.
                                // an error.
  words := strings.Split(string(text), " ")

  return words[0], words[1], err
}

// assignAnimal function.
// It takes the animal name as input and returns the corresponding animal
// type if implemented.
func assignAnimal (input string) (animal Animal, err error) {

    if input == "new" {
      insertNewAnimal()
    } else {
      if animals[input].name != "" {
        animal = animals[input] // Easy case: input animal is in the map.
      } else {
          fmt.Printf("%s is not in our list of animals. Sorry." +
            " You can create it if you want.\n", input)
          return cow, errors.New("Try again")
      }
    }
  return
}

// insertNewAnimal function
// It scans values and insert new animal in the global map animals.
func insertNewAnimal () {
  var name, eat, move, speak string
  var animal Animal
  fmt.Println("Please, enter name, food, locomotion method and noise of " +
    "the new animal, separated by spaces.")
  fmt.Scan(&name, &eat, &move, &speak)
  animal = Animal{name, eat, move, speak}
  animals[name] = animal
}

// getInfo function
// It takes the animal as input and returns the corresponding animal
// info if implemented.
func getInfo (animal Animal, input string)  {
  switch input {
  case "eat":
    animal.Eat()
  case "move":
    animal.Move()
  case "speak":
    animal.Speak()
  case "animal":
    fmt.Printf("")
  default:
    fmt.Printf("Invalid request, sorry.")
  }
}

func main() {
  fmt.Println("Please type '[cow|bird|snake] [eat|move|speak]' after " +
    "the prompt \">\" to discover animals. 'new animal' " +
    "to create a new one.")

  for {
    animalInput, featureInput := getRequest()

    animal, err := assignAnimal(animalInput); if err != nil {
        fmt.Println("Error. ", err)
        continue
    } else {
      getInfo(animal, featureInput)
    }
  }

}
