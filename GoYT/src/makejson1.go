package main

import (
	"encoding/json"
	"fmt"
)

func main7() {
	// Create a struct to hold the name and address.
	type Person struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	// Prompt the user to enter name and address.
	var name, address string
	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter Address: ")
	fmt.Scanln(&address)

	// Create a Person struct with the entered values.
	person := Person{Name: name, Address: address}

	// Marshal the struct into JSON.
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON object.
	fmt.Println(string(jsonData))
}
