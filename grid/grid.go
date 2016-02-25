package grid

import (
	"antwar/ant"
	. "antwar/util"
	//	"fmt"
)

type Grid struct {
	Pos   Position
	Color int
	Ant   *ant.Ant
}

func NewGrid(p Position) *Grid {
	g := &Grid{p, White, nil}
	return g
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

func (g *Grid) ChangeColor() {
	if g.Color == White {
		g.setBlack()
	} else if g.Color == Black {
		g.setWhite()
	}
}

func (g *Grid) AntLeave() {
	g.Ant = nil
	//Change color
	g.ChangeColor()
}

func (g *Grid) AntIn(a *ant.Ant) {
	g.Ant = a
}

func (g *Grid) KillAnt() {
	if g.Ant != nil {
		g.Ant.Killed()
	}
}
