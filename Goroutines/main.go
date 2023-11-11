package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Microsecond)
	}
	channel <- "Done!"
}
func main(){
	var channel chan string
	go printMessage("Hello world", channel)
	response := <-channel
	fmt.Println(response)
}