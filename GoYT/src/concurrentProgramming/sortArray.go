package main

import (
	"fmt"
	"sort"
	"sync"
)

func mainArry() {
	var n int
	fmt.Print("Enter the number of integers: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Printf("Enter %d integers:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Input array:", arr)

	// Create a channel to communicate between goroutines
	partChan := make(chan []int, 4)

	// Use a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Divide the array into 4 parts and sort each part concurrently
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go sortPart(&wg, arr[i*n/4:(i+1)*n/4], partChan)
	}

	// Close the channel after all parts are sent
	go func() {
		wg.Wait()
		close(partChan)
	}()

	// Merge the sorted subarrays
	mergedArr := merge(partChan)

	fmt.Println("Sorted array:", mergedArr)
}

func sortPart(wg *sync.WaitGroup, part []int, ch chan<- []int) {
	defer wg.Done()
	sort.Ints(part)
	fmt.Println("Sorted part:", part)
	ch <- part
}

func merge(ch <-chan []int) []int {
	// Collect sorted parts from the channel and merge them
	var mergedArr []int

	for part := range ch {
		mergedArr = mergeSortedArrays(mergedArr, part)
	}

	return mergedArr
}

func mergeSortedArrays(arr1, arr2 []int) []int {
	// Merge two sorted arrays into one
	var merged []int
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	// Append the remaining elements
	merged = append(merged, arr1[i:]...)
	merged = append(merged, arr2[j:]...)

	return merged
}
