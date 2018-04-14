package main

import "fmt"

func main() {
	// Anonymous go routine
	go func() {
		fmt.Println("I'm running in my own go routine")
	}()
}
