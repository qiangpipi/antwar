package main

import (
	"ants"
	"fmt"
	"grids"
	. "util"
)

func main() {
	grids := grids.CreateGrids(31, 31)
	for k, v := range grids {
		fmt.Println(k, v)
	}
	ant := ants.Ant{1, true}
	grids.PutAnt(ant, Position{10, 10})
	grids[Position{10, 10}].SetColor(Black)
	fmt.Println(grids[Position{10, 10}])
}
