package grids

import (
	"ants"
	. "util"
)

var length int
var width int

func setLW(l, w int) {
	//Set the size of the grids
	length = l
	width = w
	SetTop(l, w)
}

type Grids map[Position]*Grid

func CreateGrids(l, w int) Grids {
	//Set length and width of the world
	//Full the world with grids
	//Set all the grids with white color in init
	setLW(l, w)
	gs := make(map[Position]*Grid, l*w)
	for x := 0; x < l; x++ {
		for y := 0; y < w; y++ {
			k := Position{x, y}
			gs[k] = &Grid{White, nil}
		}
	}
	return gs
}

func (gs Grids) GetGrid(p Position) *Grid {
	return gs[p]
}

func (gs *Grids) PutAnt(a ants.Ant, p Position) {
	gs.GetGrid(p).AntIn(&a)
}

func (gs *Grids) MoveAnt(fromP Position) {
	fromG := gs.GetGrid(fromP)
	var toG *Grid
	switch fromG.Color {
	case Black:
		//Get new position
		toP := fromP.GetRight(gs.GetGrid(fromP).Ant.Direction)
		toG = gs.GetGrid(toP)
		//Get new direction
		nd := fromP.GetDirection(toP)
		//Set new direction
		fromG.Ant.ChangeDirection(nd)
	case White:
		//Get new position
		toP := fromP.GetLeft(gs.GetGrid(fromP).Ant.Direction)
		toG = gs.GetGrid(toP)
		//Get new direction
		nd := fromP.GetDirection(toP)
		//Set new direction
		fromG.Ant.ChangeDirection(nd)
	}
	//In new grid
	toG.AntIn(fromG.Ant)
	//Left grid
	fromG.AntLeft()
}

func (gs *Grids) KillAnt(p Position) {
	gs.GetGrid(p).KillAnt()
}
