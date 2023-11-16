package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)


func main(){
	PerformGETRequest()
	PerformPOSTRequest()

}

const GETURL = "http://localhost:3000/"
const POSTURL = "http://localhost:3000/items"

func PerformGETRequest(){

	response, err := http.Get(GETURL)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Statue code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	// Method 1
//	fmt.Println(content)
//	fmt.Println(string(content))


// method 2

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is : ", byteCount)
	fmt.Println(responseString.String())

}

func PerformPOSTRequest(){

	requestBody := strings.NewReader(`
	{
		"id": 6,
		"name": "Item 6",
		"description": "Description for Item 1"
	}
	`)

	response, err := http.Post(POSTURL, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()	

	fmt.Println(string(content))

}


func PerformPostFormRequest(){
	
}