package main

import (
	"fmt"

	"github.com/mayura-andrew/GoLang/CRYTPOMASTERS/api"
)


func main(){
	rate, err := api.GetRate("BTC")
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}

}