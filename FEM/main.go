package main

import (
	"fmt"

	"github.com/mayura-andrew/GoLang/go/server/data"
)

func main() {
	max := data.Instructor{Id: 3, LastName: "Andrew"}
	max.FirstName = "Mayura"

	kyle := data.NewInstructor("kyle", "Simpson", 45)
	goCourse := data.Course{Id: 2, Name: "Go Fundamentals", Instructor: max}
	fmt.Printf("%v", goCourse)
	print(max.Print())
	print(kyle.Print())
	fmt.Printf("%v", goCourse.String()) // Call the String() method
	

}

