package main

import (
	//	"antwar/ant"
	//	"antwar/grid"
	"antwar/world"
	"fmt"
)

func main() {
	myworld := world.CreateWorld(2)
	for k, v := range myworld.Grids {
		fmt.Println(k, v)
	}
}
