package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file "

	file, err := os.Create("./myText.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./myText.txt")
}


func readFile(filename string){
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("data in this file : ", string(databyte))

}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}

}