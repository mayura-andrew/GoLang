package main

import (
	"fmt"
	"net/http"
	
)


func main(){
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")

	http.ListenAndServe(":4000", r)

}
