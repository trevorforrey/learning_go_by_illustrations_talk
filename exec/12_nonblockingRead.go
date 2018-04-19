package main

import "fmt"

func main() {
	myChannel := make(chan string)
	go func() {
		myChannel <- "message received!\n"
	}()

	select {
	case msg := <-myChannel:
		fmt.Printf("Received: %v", msg)
	default:
		fmt.Printf("no message\n")
	}

	select {
	case msg := <-myChannel:
		fmt.Printf("Received: %v", msg)
	default:
		fmt.Printf("no message")
	}
}
