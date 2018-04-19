package main

import (
	"fmt"
)

func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	foundOre := make(chan string)
	minedOre := make(chan string)

	go func() {
		for _, item := range theMine {
			if item == "ore" {
				foundOre <- item
			}
		}
		close(foundOre)
	}()
	go func() {
		for ore := range foundOre {
			fmt.Printf("Got: %v\n", ore)
			minedOre <- "minedOre"
		}
		close(minedOre)
	}()

	for processedOre := range minedOre {
		fmt.Printf("Got: %v!\n", processedOre)
	}
}
