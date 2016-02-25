package world

import (
	"antwar/ant"
	"antwar/grid"
	. "antwar/util"
	//	"fmt"
)

type World struct {
	Size  int
	Ants  []*ant.Ant
	Grids map[Position]*grid.Grid
}

var MyWorld *World

func CreateWorld(n int) *World {
	//Check n<5
	var antworld World
	antworld.Size = n * 11
	antworld.Ants = make([]*ant.Ant, n)
	antworld.Grids = make(map[Position]*grid.Grid)
	for i := 0; i < antworld.Size; i++ {
		for j := 0; j < antworld.Size; j++ {
			antworld.Grids[Position{i, j}] = grid.NewGrid(Position{i, j})
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			antworld.Ants[i] = ant.NewAnt(Position{i*11 + 5, j*11 + 5})
			antworld.Grids[antworld.Ants[i].Pos].AntIn(antworld.Ants[i])
		}
	}
	return &antworld
}

func (w *World) WorldRun()
