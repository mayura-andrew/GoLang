package main

import (
	"fmt"
	"strings"
)

// Animal is a struct to hold information about each animal
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat prints the animal's food
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move prints the animal's locomotion
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak prints the animal's spoken sound
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	// Define animals with their associated data
	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	// Main program loop
	for {
		// Prompt for user input
		fmt.Print("> ")

		// Read user input
		var animalName, info string
		_, err := fmt.Scan(&animalName, &info)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Convert animal name to lowercase for case-insensitivity
		animalName = strings.ToLower(animalName)

		// Validate the input animal name
		if _, ok := animals[animalName]; !ok {
			fmt.Println("Invalid animal name. Please enter 'cow', 'bird', or 'snake'.")
			continue
		}

		// Process user request and call the appropriate method
		switch info {
		case "eat":
			animals[animalName].Eat()
		case "move":
			animals[animalName].Move()
		case "speak":
			animals[animalName].Speak()
		default:
			fmt.Println("Invalid information request. Please enter 'eat', 'move', or 'speak'.")
		}
	}
}
