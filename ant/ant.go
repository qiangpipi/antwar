package ant

import (
	//	"math/rand"
	. "antwar/util"
)

type Ant struct {
	Direction int
	Pos       Position
	IsAlive   bool
}

type Ants map[Position]*Ant

func NewAnt(p Position) *Ant {
	a := &Ant{West, p, true}
	return a
}

func (a *Ant) ChangeDirection(d int) {
	a.Direction = d
}

func (a *Ant) Killed() {
	a.IsAlive = false
}
