package grids

import (
	"ants"
	. "util"
)

var length int
var width int

func setLW(l, w int) {
	length = l
	width = w
	SetTop(l, w)
}

type Grids map[Position]*Grid

func CreateGrids(l, w int) Grids {
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
	//Get new position
	toP := fromP.GetLeft(gs.GetGrid(fromP).Ant.Direction)
	g := gs.GetGrid(fromP)
	switch g.Color {
	case Black:
		//Get new direction
		nd := fromP.GetDirection(toP)
		//Set new direction
		g.Ant.ChangeDirection(nd)
	case White:
		//Get new direction
		nd := fromP.GetDirection(toP)
		//Set new direction
		g.Ant.ChangeDirection(nd)
	}
	//Left grid
	//In new grid
}
