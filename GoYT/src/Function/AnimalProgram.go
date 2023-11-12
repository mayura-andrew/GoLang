 package main

// import "fmt"


// type Animal struct {
// 	food string
// 	locomotion string
// 	noise string
// }

// func (a Animal) Eat(){
// 	fmt.Println(a.food)
// }

// func (a Animal) Move(){
// 	fmt.Println(a.locomotion)
// }

// func (a Animal) Speak(){
// 	fmt.Println(a.noise)
// }

// func main(){
// 	animals := map[string] Animal{
// 		"cow": Animal{"grass", "walk", "moo"},
// 		"bird": Animal{"worms", "fly", "peep"},
// 		"snake": Animal{"mice", "slither", "hssss"},
// 	} 

// 	for {
// 		var animalName, infoRequested string
// 		fmt.Print("> ")
// 		fmt.Scan(&animalName, &infoRequested)

// 		animal, ok := animals[animalName]

// 		if !ok {
// 			fmt.Println("invalid animal name ")
// 			continue
// 		}  
// 		switch infoRequested {
// 		case "eat":
// 			animal.Eat()
// 		case "move":
// 			animal.Move()
// 		case "speak":
// 			animal.Speak()
// 		default:
// 			fmt.Println("Invalid info requested")
// 		}
// 	}
// }

