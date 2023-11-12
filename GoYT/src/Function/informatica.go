package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface describes the methods of an animal
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow type
type Cow struct{}

// Eat method for Cow
func (c Cow) Eat() {
	fmt.Println("grass")
}

// Move method for Cow
func (c Cow) Move() {
	fmt.Println("walk")
}

// Speak method for Cow
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird type
type Bird struct{}

// Eat method for Bird
func (b Bird) Eat() {
	fmt.Println("worms")
}

// Move method for Bird
func (b Bird) Move() {
	fmt.Println("fly")
}

// Speak method for Bird
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake type
type Snake struct{}

// Eat method for Snake
func (s Snake) Eat() {
	fmt.Println("mice")
}

// Move method for Snake
func (s Snake) Move() {
	fmt.Println("slither")
}

// Speak method for Snake
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animalMap := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		command := scanner.Text()
		parts := strings.Fields(command)

		switch parts[0] {
		case "newanimal":
			if len(parts) != 3 {
				fmt.Println("Invalid command. Format: newanimal name type")
				continue
			}
			name := parts[1]
			animalType := parts[2]

			switch animalType {
			case "cow":
				animalMap[name] = Cow{}
			case "bird":
				animalMap[name] = Bird{}
			case "snake":
				animalMap[name] = Snake{}
			default:
				fmt.Println("Invalid animal type. Choose from cow, bird, or snake.")
			}
			fmt.Println("Created it!")
		case "query":
			if len(parts) != 3 {
				fmt.Println("Invalid command. Format: query name information")
				continue
			}
			name := parts[1]
			info := parts[2]

			animal, ok := animalMap[name]
			if !ok {
				fmt.Println("Animal not found.")
				continue
			}

			switch info {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid information request. Choose from eat, move, or speak.")
			}
		default:
			fmt.Println("Invalid command. Choose from newanimal or query.")
		}
	}
}
