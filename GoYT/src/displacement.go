package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getParameter(paramName string) float64 {

	// Gets the parameter and validates it is a float64

	var param string
	var value float64
	var err error

	for {
		fmt.Printf("Enter %s: ", paramName)
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			param = scanner.Text()
		}

		value, err = strconv.ParseFloat(param, 64)
		if err == nil {
			break
		}
	}
	return value
}

// func GenDisplaceFn(a, v, s float64) func(t float64) float64 {
// 	fn := func(t float64) float64 {
// 		return t*t*a*1/2 + v*t + s
// 	}
// 	return fn
// }

func Displacement() {

	var a, v, s, t float64

	a = getParameter("acceleration")
	v = getParameter("velocity")
	s = getParameter("initial displacement")
	t = getParameter("time")

	fn := GenDisplaceFn(a, v, s)
	fmt.Printf("Total displacement: %f", fn(t))
}
