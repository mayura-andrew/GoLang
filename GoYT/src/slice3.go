package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// )

// func main() {
// 	slice := make([]int, 0, 3)

// 	for {
// 		fmt.Print("Enter an integer (X to quit): ")
// 		var input string
// 		fmt.Scanln(&input)

// 		if input == "X" {
// 			break
// 		}

// 		num, err := strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("Invalid input. Please enter a valid integer or 'X' to quit.")
// 			continue
// 		}

// 		slice = append(slice, num)
// 		sort.Ints(slice)

// 		fmt.Println("Sorted slice:", slice)
// 	}
// }
