package main

func main() {
	myFirstChannel := make(chan string)
	myFirstChannel <- "hello"
}
