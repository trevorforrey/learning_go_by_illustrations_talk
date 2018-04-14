package main

import (
	"github.com/trevorforrey/learning_go_by_illustrations_talk/exec/roles"
)

func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	foundOre := roles.finder(theMine)
	minedOre := roles.miner(foundOre)
	roles.smelter(minedOre)
}
