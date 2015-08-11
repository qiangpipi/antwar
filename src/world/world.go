package world

import (
	//	"ants"
	"grids"
	//	. "util"
)

//New grids create
func NewWorld(length, width int) grids.Grids {
	return grids.CreateGrids(length, width)
}

//Put ants in grids
//All ants are living and random direction

//World runing
