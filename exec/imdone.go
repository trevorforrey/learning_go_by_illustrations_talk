package main

import "fmt"

func main() {
	doneChan := make(chan string)
	go func() {
		fmt.Printf("Hi there!\n")
		fmt.Printf("Hello?!")
		doneChan <- "I’m all done!"
	}()

	<-doneChan // block until go routine signals work is done
}
