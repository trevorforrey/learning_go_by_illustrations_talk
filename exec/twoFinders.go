package main

import (
	"fmt"
	"time"
)

// import local package of finder functions

func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	go finder(theMine, 1)
	go finder(theMine, 2)
	<-time.After(time.Second * 5) //you can ignore this for now
}

// Finds all 'ore' in the mine
func finder(mine [5]string, number int) {
	for _, item := range mine {
		if item == "ore" {
			fmt.Printf("From Finder %v: %v\n", number, item)
		}
	}
	return
}
