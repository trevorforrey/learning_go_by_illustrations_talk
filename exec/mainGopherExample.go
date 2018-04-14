package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Printf("Hi there!")
	}()
}
