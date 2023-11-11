package main

import (
	"encoding/json"
	"fmt"
)

func makeJson(){

	details := make(map[string]string)

	var name string
	var addr string

	fmt.Println("Enter your name: ")
	_, err := fmt.Scan(&name)

	if err != nil {
		fmt.Print("Error : %s\n", err)
		return
	}

	fmt.Println("Enter your address")
	_, errA := fmt.Scan(&addr)

	if errA != nil {
		fmt.Print("Error : %s", errA)
		return
	}

	details["name"] = name
	details["address"] = addr

	jsonStr, err := json.Marshal(details)

	if err != nil {
		fmt.Printf("Error : %s\n", err)
		return
	}

	fmt.Println(string(jsonStr))



}

// func main(){
// 	makeJson()
// }