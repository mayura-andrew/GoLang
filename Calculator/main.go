package main

import "fmt"

func main(){
	var operation string
	var number1, number2 int
	fmt.Println("<--- SIMPLE CALCULATOR 0.1 -->")
	fmt.Println("==============================")
	fmt.Println("---       Operations       --- \n 1. [+]Addition \n 2. [-]Substract \n 3. [*]Multiply \n 4. [/]Divide")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter your first number : ")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter your second number : ")
	fmt.Scanf("%d", &number2)
	switch operation {
	case "1":
		fmt.Println("Addition ->", number1 + number2)
	case "2":
		fmt.Println("Substraction ->", number1 - number2)
	case "3":
		fmt.Println("Multiplication ->", number1 * number2)
	case "4":
		fmt.Println("Divition ->", number1 / number2)
	}


}