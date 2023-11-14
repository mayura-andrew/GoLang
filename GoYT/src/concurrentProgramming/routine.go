package main

import(
	"fmt"
)

var counter = 10

func increment(){
	counter++
}

func decrement(){
	counter--
}

func mainRoutine() { 
	i := 0
	for i < 1000 {
		go increment()
		go decrement()
		i++
	}

	fmt.Println(counter)
	
	//Explanation
	//Because there a two goroutines running and performing different operation on a shared resource, there is
	//no deterministic order in which there are executed. So for instance, the increment operations will be interleaving
	//with the decrement operation in no specific order. This causes a race condition as no one can predict the value
	//of the counter in each run of the program. Counter could be any value as it could be increment or decreased in
	//any order
	
}