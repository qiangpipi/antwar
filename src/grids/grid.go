package grids

import (
	. "util"
)

type Grid struct {
	Color int
	Ant   *ants.Ant
}

func (g *Grid) SetColor(c int) {
	g.Color = c
}

func (g *Grid) setWhite() {
	g.SetColor(White)
}

func (g *Grid) setBlack() {
	g.SetColor(Black)
}

func (g Grid) GetColor() (c int) {
	return g.Color
}

func (g *Grid) AntLeft() {
	g.Ant = nil
	if g.Color == White {
		g.setBlack()
		//Get direction of right
	} else {
		g.setWhite()
		//Get direction of left
	}
}

func (g *Grid) AntIn(a *ants.Ant) {
	g.Ant = a
}
