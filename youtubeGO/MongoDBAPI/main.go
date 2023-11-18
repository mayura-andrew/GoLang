package main

import (
	"fmt"
	"net/http"
	"github.com/mayura-andrew/mongodbapi/router"
)


func main(){
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")

	http.ListenAndServe(":4000", r)

}
