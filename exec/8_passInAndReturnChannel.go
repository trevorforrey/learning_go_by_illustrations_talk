package main

import (
	"fmt"
)

func main() {
	baseData := [5]string{"rock", "ore", "ore", "rock", "ore"}
	generatedChannel := gen(baseData)
	outputChannel := finder(generatedChannel)
	for ore := range outputChannel {
		fmt.Printf("%v found!\n", ore)
	}
}

func gen(mine [5]string) <-chan string {
	out := make(chan string)
	go func() {
		for _, item := range mine {
			out <- item
		}
		close(out)
	}()
	return out
}

// Finds all 'ore' in the mine
func finder(mineChan <-chan string) <-chan string {
	foundOreChan := make(chan string)
	go func() {
		for item := range mineChan {
			if item == "ore" {
				foundOreChan <- item
				fmt.Println("Finder found ore")
			}
		}
		close(foundOreChan)
	}()
	return foundOreChan
}
