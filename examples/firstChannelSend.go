package main

func main() {
	myFirstChannel := make(chan string)
	myFirstChannel <- "hello"      // Send
	myVariable := <-myFirstChannel // Receive
}
