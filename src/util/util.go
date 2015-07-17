package util

type Position struct {
	X, Y int
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

var Pole = &Position{0, 0}
var Top = &Position{0, 0}

func SetTop(l, w int) {
	Top = &Position{l, w}
}

func (p Position) GetEast() Position {
	if p.X+1 > Top.X {
		return p.GetWest()
	} else {
		return Position{p.X + 1, p.Y}
	}
}

func (p Position) GetSouth() Position {
	if p.Y+1 > Top.Y {
		return p.GetNorth()
	} else {
		return Position{p.X, p.Y + 1}
	}
}

func (p Position) GetWest() Position {
	if p.X-1 < 0 {
		return p.GetEast()
	} else {
		return Position{p.X - 1, p.Y}
	}
}

func (p Position) GetNorth() Position {
	if p.Y-1 < 0 {
		return p.GetSouth()
	} else {
		return Position{p.X, p.Y - 1}
	}
}

func (p Position) GetLeft(d int) Position {
	var t Position
	switch d {
	case East:
		t = p.GetSouth()
	case South:
		t = p.GetWest()
	case West:
		t = p.GetNorth()
	case North:
		t = p.GetEast()
	}
	return t
}

func (p Position) GetRight(d int) Position {
	var t Position
	switch d {
	case East:
		t = p.GetNorth()
	case South:
		t = p.GetEast()
	case West:
		t = p.GetSouth()
	case North:
		t = p.GetWest()
	}
	return t
}

func (fromP Position) GetDirection(toP Position) int {
	v := Position{toP.X - fromP.X, toP.Y - fromP.Y}
	var direction int
	switch v {
	case Position{1, 0}:
		direction = East
	case Position{0, 1}:
		direction = South
	case Position{-1, 0}:
		direction = West
	case Position{0, -1}:
		direction = North
	}
	return direction
}
