package main

import "fmt"

// return function
func GenDisplaceFn(a, vo, so float64)  func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func mainGen(){
	var a, vo, so, t float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&vo)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&so)
	fmt.Print("Enter time: ")
	fmt.Scan(&t)
	fn := GenDisplaceFn(a, vo, so)
	fmt.Println(fn(t))
}