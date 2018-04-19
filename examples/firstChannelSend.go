package main

// Example code not meant to run
// This program will immediately deadlock if ran
// Due to the main goroutine sending on a channel without a routine reading from it

func main() {
	myFirstChannel := make(chan string)

	myFirstChannel <- "hello"      // Send
	myVariable := <-myFirstChannel // Receive
	<-myFirstChannel               // Receive and discard result
}
