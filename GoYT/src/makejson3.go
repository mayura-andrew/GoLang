package main

import (
	"encoding/json"
	"fmt"
)

func main6() {
	// Prompt user for name
	fmt.Print("Enter a name: ")
	var name string
	_, err := fmt.Scan(&name)

	// Check for input error
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	// Prompt user for address
	fmt.Print("Enter an address: ")
	var address string
	_, err = fmt.Scan(&address)

	// Check for input error
	if err != nil {
		fmt.Println("Error reading address:", err)
		return
	}

	// Create a map with name and address
	data := map[string]string{
		"name":    name,
		"address": address,
	}

	// Convert map to JSON
	jsonData, err := json.Marshal(data)

	// Check for JSON marshaling error
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		return
	}

	// Print the JSON object
	fmt.Println("JSON Object:", string(jsonData))
}
