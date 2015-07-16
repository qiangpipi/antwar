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

func (g Grids) GetGrid(p Position) *Grid {
	return g[p]
}

func (g *Grids) PutAnt(a ants.Ant, p Position) {
	g.GetGrid(p).AntIn(&a)
}

func (g *Grids) MoveAnt(fromP Position) {
	//Get new position
	var toP Position
	switch g.GetGrid(p).Color {
	case Black:
		toP = fromP.GetLeft(g.GetGrid(p).Ant.Direction)
		//Get new direction
		//Set new direction
	case White:
		toP = fromP.GetRight(g.GetGrid(p).Ant.Direction)
		//Get new direction
		//Set new direction
	}
	//Left grid
	//In new grid
}
