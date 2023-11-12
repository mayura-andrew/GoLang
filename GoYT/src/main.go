package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn generates a displacement function based on the provided parameters.
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + vo*t + so
	}
}

func mainMain() {
	// Prompt the user to enter values for acceleration, initial velocity, and initial displacement.
	var acceleration, initialVelocity, initialDisplacement float64

	fmt.Print("Enter acceleration: ")
	fmt.Scan(&acceleration)

	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&initialVelocity)

	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&initialDisplacement)

	// Generate the displacement function.
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	// Prompt the user to enter a value for time.
	var time float64
	fmt.Print("Enter time: ")
	fmt.Scan(&time)

	// Compute and print the displacement after the entered time.
	displacement := fn(time)
	fmt.Printf("Displacement after %.2f seconds: %.2f\n", time, displacement)
}
