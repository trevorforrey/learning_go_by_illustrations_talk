package main

import (
	"fmt"
)

func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	mineChan := gen(theMine)
	foundOreChan := finder(mineChan)
	minedOreChan := miner(foundOreChan)
	smeltedOreChan := smelter(minedOreChan)
	for smeltedOre := range smeltedOreChan {
		fmt.Printf("%v processed\n", smeltedOre)
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

// Turns all found 'ore' into 'minedOre'
func miner(oreChan <-chan string) <-chan string {
	minedOreChan := make(chan string)
	go func() {
		for ore := range oreChan {
			minedOreChan <- ore
			fmt.Println("Miner mined ore: ")
		}
		close(minedOreChan)
	}()
	return minedOreChan
}

// Smelts all 'minedOre' into 1 mixture
func smelter(minedOreChan <-chan string) <-chan string {
	smeltedOreChan := make(chan string)
	go func() {
		for minedOre := range minedOreChan {
			smeltedOreChan <- minedOre
			fmt.Println("Smelter smelted ore")
		}
		close(smeltedOreChan)
	}()
	return smeltedOreChan
}
