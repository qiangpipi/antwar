package main

import (
	"antwar/ants"
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
	ant := ants.NewAnt()
	//Put ants in the grids
	world.PutAnts(ants)
}
