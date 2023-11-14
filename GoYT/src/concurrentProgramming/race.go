package main

import (
	"fmt"
)

var x = 10

func first() {
	fmt.Print("first")
	x = x + 1
}

func second() {
	fmt.Print("second")
	x = x * 2
}

func printX() {
	fmt.Print(x)
}

func mainRace() {
	go first()
	go second()
	go printX()
}

// A race condition is a problem that can happen as a result of interleaving that you have to consider.
// Race condition is usually defined as a program that runs, a problem where the outcome of the program
// depends on the interleaving
