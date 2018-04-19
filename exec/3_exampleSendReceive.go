package main

import "fmt"

func main() {
	myFirstChannel := make(chan string)
	go func() {
		// send on channel
		myFirstChannel <- "hello"
	}()
	message := <-myFirstChannel
	fmt.Printf("Received: %v", message)
}
