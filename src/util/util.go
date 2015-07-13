package util

type Position struct {
	x, y int
}

const (
	White = iota
	Black
)

const (
	East = iota
	South
	West
	North
)
