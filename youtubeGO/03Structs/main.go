package main

import "fmt"

func main(){
	fmt.Println("Structs in golang")
	// no inheritance in go lang, No super or parent, child
	hitesh := Person{"Mayura", "dev@mayuraandrew.tech", true, 16}
	fmt.Printf("details: %+v\n", hitesh)
	fmt.Printf(" Name is %v and email is %v.\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()

	// control flow in go

	if num :=23; num < 10 {
		fmt.Println("Less")
	} else {
		fmt.Println("High")
	}

	result, msg := proAdder(3,4,5,5)
	fmt.Println("Result ", result)
	fmt.Println("message", msg)

}


func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hello"
}


type Person struct {
	Name string
	Email string
	Status bool
	Age int
}

func (p Person) GetStatus(){
	fmt.Println("Is this person active: ", p.Status)
}
