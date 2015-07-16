package ants

import (
	. "util"
)

type Ant struct {
	Direction int
	IsAlive   bool
}

type Ants map[Postion]*Ant

func (as *Ants) AddAnt(a *Ant, p Position) {
	as[p] = a
}

func (as *Ants) FindAnt(p Postion) *Ant {
	return as[p]
}

func (as *Ants) KillAnt(p Position) {
	as.FindAnt(p).IsAlive = false
}

func (a *Ant) ChangeDirection(d int) {
	a.Direction = d
}
