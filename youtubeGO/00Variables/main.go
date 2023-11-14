package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const X = 10

func main(){
	fmt.Println("variables")
	fmt.Println(X)

	//var userName string = "Mayura"
//	fmt.Println(userName)
	// fmt.Printf("Variable is of type: %T \n", userName)
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter something")
	// //comma ok || err error
	// input, _ := reader.ReadString('\n')
	// fmt.Printf("Type of entered value %T ", input)
	// fmt.Println(input)

	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pixzza between 1 and 5")


	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Addredd 1 to your rating ", numRating +1)
	}


	// data time package

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	

}