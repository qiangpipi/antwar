package grids

import "util"

const Zero = Position(0, 0)

type Grid struct {
	P     Position
	color int
}

func (g *Grid) SetColor(c int) {
	g.color = c
}

func (g *Grid) SetWhite() {
	g.SetColor(util.White)
}

func (g *Grid) SetBlack() {
	g.SetColor(util.Black)
}

func (g Grid) GetColor() (c int) {
	return g.color
}
