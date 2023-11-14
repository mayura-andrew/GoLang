// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	resultNumbers := make([]int, 0)
// 	resultChan := make(chan []int, 4)
// 	numRoutines := 4

// 	initialNumbers := getNumbersByPrompt()

// 	chunkSize := len(initialNumbers) / numRoutines
// 	chunks := getChunks(initialNumbers, chunkSize)

// 	for i := 0; i < numRoutines; i++ {
// 		go sortFragment(chunks[i], resultChan)
// 	}

// 	for i := 0; i < numRoutines; i++ {
// 		resultNumbers = append(resultNumbers, <-resultChan...)
// 	}
// 	sort.Ints(resultNumbers)
// 	fmt.Printf("\nOrdered Array : %v", resultNumbers)
// }

// func getNumbersByPrompt() []int {
// 	var numbers []int
// 	prompt := ""

// 	fmt.Println("Enter array of numbers like: [1,3,8,9]")
// 	fmt.Scanln(&prompt)

// 	err := json.Unmarshal([]byte(prompt), &numbers)
// 	if err != nil {
// 		panic("Error: Not an array")
// 	}

// 	return numbers
// }

// func sortFragment(values []int, channel chan []int) {
// 	fmt.Printf("\nWill sort this fragment: %v", values)
// 	sort.Ints(values)
// 	channel <- values
// }

// func getChunks(numbers []int, chunkSize int) [][]int {
// 	var chunks [][]int
// 	mod := len(numbers) % 4
// 	originalChunkSize := chunkSize

// 	for i := 0; i < len(numbers); i += chunkSize {
// 		chunkSize = originalChunkSize

// 		if mod != 0 {
// 			mod--
// 			chunkSize++
// 		}

// 		end := i + chunkSize

// 		if end > len(numbers) {
// 			end = len(numbers)
// 		}

// 		chunks = append(chunks, numbers[i:end])
// 	}
// 	return chunks
// }
