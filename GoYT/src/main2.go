package main

import (
	"fmt"
	//"strconv"
)

func truncMy(){
	var floatNum float64
	var intNum int64
	fmt.Print("Enter a floating point number = ")
	_, err := fmt.Scan(&floatNum)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	intNum = int64(floatNum) // Typecasting 
	fmt.Printf("%d\n", intNum)
}
func main2(){
	// fmt.Println("Hello, world!")
	// var name string

	// fmt.Scan(&name)

	// parsed_number, _ := strconv.ParseFloat(name, 64)

	// fmt.Printf("%d", int(parsed_number))
	// truncMy()
	var number float32
	fmt.Print("Enter a float number in order to truncate it: ")
	fmt.Scan(&number)
	truncatedNumber := int(number)
	fmt.Print("Truncated value is: ", truncatedNumber)

}