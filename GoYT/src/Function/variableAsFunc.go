package main

import "fmt"

//varible as func

// var funcVar func(int) int 
// func incFn (x int) int {
// 	return x +1
// }

// function as arguments

func applyIt(aFunc func (int) int, val int) int {
	return aFunc(val)
}

func incFn(x int) int {return x +1}
func decFn(x int) int {return x -1}

func main(){
	//funcVar = incFn
	//fmt.Print(funcVar(1))
	fmt.Println(applyIt(incFn,2))
	fmt.Println(applyIt(decFn,2))

	// anonymouse functions
	v := applyIt(func (x int) int {return x+1},2)
	fmt.Println(v)


}