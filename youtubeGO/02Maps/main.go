package main

import "fmt"


func main(){
	languages := make(map[string]string)

	languages["JS"] = "javaScript"
	languages["RB"] = "Ruby"
	languages["Java"] = "Java"

	fmt.Println(languages)
	for key, value := range languages{
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	delete(languages, "JS")


	fmt.Println(languages)



}

