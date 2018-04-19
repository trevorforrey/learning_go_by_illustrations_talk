package main

import (
	"fmt"
	"time"
)

func main() {
	dataArray := [7]string{"1", "2", "3", "4", "5", "6", "7"}
	myFirstChannel := make(chan string, 3)
	go func() {
		for _, item := range dataArray {
			myFirstChannel <- item
			fmt.Printf("Sent: %v\n", item)
		}
		close(myFirstChannel)
	}()
	for item := range myFirstChannel {
		<-time.After(1 * time.Second)
		fmt.Printf("Received: %v\n", item)
	}
}
