package main

import (
	//	"ants"
	"fmt"
	//	"grids"
	//	. "util"
	"world"
)

func main() {
	W := world.NewWorld(31, 31)
	for k, v := range W {
		fmt.Println(k, v)
	}
	//	ant := ants.Ant{1, true}
	//	grids.PutAnt(ant, Position{10, 10})
	//	grids[Position{10, 10}].SetColor(Black)
	//	fmt.Println(grids[Position{10, 10}])
}
