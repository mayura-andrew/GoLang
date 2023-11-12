package main

import (
	"fmt"
)


type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
	food string
	locomotion string
	noise string
}

func (c Cow) Eat(){
	fmt.Println(c.food)
}

func (c Cow) Move(){
	fmt.Println(c.locomotion)
}

func (c Cow) Speak(){
	fmt.Println(c.noise)
}

type Bird struct {
	name string
	food string
	locomotion string
	noise string
}

func (b Bird) Eat(){
	fmt.Println(b.food)
}

func (b Bird) Move(){
	fmt.Println(b.locomotion)
}

func (b Bird) Speak(){
	fmt.Println(b.noise)
}


type Snake struct {
	name string
	food string
	locomotion string
	noise string
}

func (s Snake) Eat(){
	fmt.Println(s.food)
}

func (s Snake) Move(){
	fmt.Println(s.locomotion)
}

func (s Snake) Speak(){
	fmt.Println(s.noise)
}

func main(){
	animals := make(map[string]Animal)

	for {
		var command, name, animalType string
		fmt.Print("> ")
		fmt.Scan(&command, &name, &animalType)

		switch command {
		case "newanimal":
			switch animalType{
			case "cow":
				animals[name] = Cow{name, "grass", "walk", "moo"}
			case "bird":
				animals[name] = Bird{name, "worms", "fly", "peep"}
			case "snake":
				animals[name] = Snake{name, "mice", "slither", "hsss"}
			default:
				fmt.Println("Invalid animal type")
			}
			fmt.Println("Created it!")
		case "query":
			animal, ok := animals[name]

			if !ok {
				fmt.Println("invalid animal name")
				continue
			}

			switch animalType{
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid animal info")
			}
		default:
			fmt.Println("Invalid command")
		}
	}
}