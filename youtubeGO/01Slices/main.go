package main

import (
	"fmt"
	"sort"
)

func main(){
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of the fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 223
	highScores[2] = 23
	highScores[3] = 43
	//highScores[4] = 43

	highScores = append(highScores, 44, 55, 666)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2

	// remove items

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}