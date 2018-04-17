package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan string)
	go func() {
		msg := "message"
		select {
		case myChannel <- msg:
			fmt.Printf("sent: %v\n", msg)
		default:
			fmt.Printf("no one home to receive\n")
		}
		close(myChannel)
	}()

	<-time.After(time.Second * 2)
}
