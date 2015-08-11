package ants

import (
	"math/rand"
	. "util"
)

type Ant struct {
	Direction int
	IsAlive   bool
}

type Ants map[Position]*Ant

func NewAnt() Ant {
	a := Ant{rand.Intn(3), true}
	return a
}

func (as Ants) FindAnt(p Position) *Ant {
	return as[p]
}

func (as *Ants) Kill(p Position) {
	as.FindAnt(p).IsAlive = false
}

func (a *Ant) ChangeDirection(d int) {
	a.Direction = d
}

func (a *Ant) Killed() {
	a.IsAlive = false
}
