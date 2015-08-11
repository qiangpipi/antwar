package main

import (
	"ants"
	"fmt"
	"grids"
	//	. "util"
)

func main() {
	world := grids.CreateGrids(31, 31)
	for k, v := range world {
		fmt.Println(k, v)
	}
	//Create ants
	ants := ants.NewAnts(5)
	//Put ants in the grids
	world.PutAnts(ants)
}
