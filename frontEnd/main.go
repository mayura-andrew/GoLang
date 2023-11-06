package main

import "fmt"

func birthDay(pointerAge *int){
	fmt.Printf("The pointer (memory address) is %v and the value is %v\n", pointerAge, *pointerAge)
	*pointerAge++
}

func main() {
	age := 2
	birthDay(&age)
	fmt.Println(age)
	
	var  x int = 32
	var p *int = &x // Storing the memory address of x in ptr
	fmt.Println("Value of x:", *p) // Dereferencing ptr to get the value stored at the memory address

}
