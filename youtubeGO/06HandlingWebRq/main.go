package main

import (
	"fmt"
	// "io/ioutil"
	// "net/http"
	"net/url"
)

const URL = "https://mayuraandrew.tech/whoami"
const newURL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghb456gph"


func main(){
	fmt.Println("mayuraandrew.tech")

	// response, err := http.Get(URL)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("response if of type %T\n", response)
	// defer response.Body.Close()

	// dataBytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// content := string(dataBytes)
	// fmt.Println(content)


	fmt.Println(newURL)
	// parsing

	result, _ := url.Parse(newURL)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}
	partOfUrl := &url.URL {
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partOfUrl.String()
	fmt.Println(anotherURL)


}