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

func newAnt() Ant {
	a := Ant{rand.Intn(3), true}
	return a
}

func NewAnts(n int) Ants {
	as := make(map[Position]*Ant, n)
	return as
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
