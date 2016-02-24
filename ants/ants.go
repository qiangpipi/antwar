package ants

import (
	//	"math/rand"
	. "antwar/util"
)

type Ant struct {
	Direction int
	IsAlive   bool
}

type Ants map[Position]*Ant

func NewAnt() *Ant {
	a := &Ant{West, true}
	return a
}

func (a *Ant) ChangeDirection(d int) {
	a.Direction = d
}

func (a *Ant) Killed() {
	a.IsAlive = false
}
