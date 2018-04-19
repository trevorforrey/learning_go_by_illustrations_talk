package main

import "fmt"

func main() {
	myFirstChannel := make(chan string)
	message := <-myFirstChannel
	fmt.Printf("Message: %v", message)
}
