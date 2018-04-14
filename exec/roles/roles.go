package roles

import (
	"bytes"
	"fmt"
)

// Finds all 'ore' in the mine
func finder(mine [5]string) [3]string {
	var foundOre [3]string
	arrayIndex := 0
	for _, item := range mine {
		if item == "ore" {
			foundOre[arrayIndex] = item
			arrayIndex++
		}
	}
	fmt.Println("From Finder: ", foundOre)
	return foundOre
}

// Turns all found 'ore' into 'minedOre'
func miner(foundOre [3]string) [3]string {
	var minedOre [3]string
	for index, _ := range foundOre {
		minedOre[index] = "minedOre"
	}
	fmt.Println("From Miner: ", minedOre)
	return minedOre
}

// Smelts all 'minedOre' into 1 mixture
func smelter(minedOre [3]string) string {
	var buffer bytes.Buffer
	for _, minedOre := range minedOre {
		buffer.WriteString(minedOre + "-")
	}
	fmt.Println("From Smelter: " + buffer.String())
	return buffer.String()
}
