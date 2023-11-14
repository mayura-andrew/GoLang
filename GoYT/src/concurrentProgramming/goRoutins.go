package main

import (
	"fmt"
	"sync"
)

var counter = 0
//var wg sync.WaitGroup

func increment(){
	for i := 0; i<10000; i++ {
		counter++
	}
	wg.Done()
}

func mainGo(){
	wg.Add(2)

	go increment()
	go increment()

	wg.Wait()

	fmt.Println("Counter: ", counter)
}


// Explanation:

// In this program, we have a shared variable counter that is accessed and modified by two goroutines (increment function) concurrently. The sync.WaitGroup is used to wait for both goroutines to finish.

// The race condition occurs because the counter++ operation is not atomic. It involves reading the current value of counter, incrementing it, and then storing it back. If these steps are interleaved with another goroutine, the final result might not be what is expected.

// Race conditions happen when multiple goroutines access shared data concurrently, and at least one of them modifies the data. In this case, if Goroutine 1 reads the counter value, and before it can increment and update it, Goroutine 2 also reads and updates the counter, both goroutines are unaware of each other's changes, leading to a final counter value that is less than the expected sum of increments.

// To run this program and detect race conditions, you can use the -race flag with the go run command:

// bash

// go run -race filename.go

// The output will indicate if any race conditions were detected during the execution.

// Remember, race conditions can lead to unpredictable and erroneous behavior in concurrent programs. It's essential to use synchronization mechanisms, such as mutexes, to protect shared data and prevent race conditions.