package ants

import (
	. "util"
)

type Ant struct {
	Direction int
	IsAlive   bool
}

type Ants map[Position]*Ant

func (as Ants) FindAnt(p Position) *Ant {
	return as[p]
}

func (as *Ants) KillAnt(p Position) {
	as.FindAnt(p).IsAlive = false
}

func (a *Ant) ChangeDirection(d int) {
	a.Direction = d
}
