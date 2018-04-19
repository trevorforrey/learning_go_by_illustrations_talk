package main

import "fmt"

// Example code not meant to run

func main() {
	// Anonymous go routine
	go func() {
		fmt.Println("I'm running in my own go routine")
	}()
}
