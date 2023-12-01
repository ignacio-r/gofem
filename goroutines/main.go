package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}
	channel <- "Done!"
}

func main() {

	var channel chan string

	//go printMessage("Go is great!")
	go printMessage("Frontend Masters Rocks!", channel)
	//go printMessage("We miss classes!")

	response := <-channel

	fmt.Println(response)
	//deadlock!!
}
