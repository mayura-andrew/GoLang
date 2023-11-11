 package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// )

// func main() {
// 	// Create an empty integer slice of size 3
// 	slice := make([]int, 0, 3)

// 	for {
// 		// Prompt user for input
// 		fmt.Print("Enter an integer (or 'X' to exit): ")
// 		var input string
// 		_, err := fmt.Scan(&input)

// 		// Check for input error
// 		if err != nil {
// 			fmt.Println("Error reading input:", err)
// 			return
// 		}

// 		// Check if the user wants to exit
// 		if input == "X" || input == "x" {
// 			break
// 		}

// 		// Convert input to an integer
// 		num, err := strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("Invalid input. Please enter a valid integer or 'X' to exit.")
// 			continue
// 		}

// 		// Add the integer to the slice
// 		slice = append(slice, num)

// 		// Sort the slice
// 		sort.Ints(slice)

// 		// Print the sorted slice
// 		fmt.Println("Sorted Slice:", slice)
// 	}
// }
