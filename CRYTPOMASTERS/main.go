package main

import (
	"fmt"

	"github.com/mayura-andrew/GoLang/CRYTPOMASTERS/api"
)


func main(){
	//rate, err := api.GetRate("BTC")
	getCurrencyData("BTC")
	getCurrencyData("ETH")
	getCurrencyData("BCH")


}

func getCurrencyData(currency string){
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}

}