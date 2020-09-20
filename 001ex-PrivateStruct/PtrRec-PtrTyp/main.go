package main

import "fmt"

// Position defines necessary functions for a position.
type Position interface {
	GetX() float64
	GetY() float64
	GetZ() float64
}

type position struct {
	x, y, z float64
}

// NewPosition returns an initial position
func NewPosition(pos ...float64) Position {
	if len(pos) == 3 {
		return &position{pos[0], pos[1], pos[2]}
	}
	return &position{0.0, 0.0, 0.0}
}

func (p *position) GetX() float64 {
	return p.x
}
func (p *position) GetY() float64 {
	return p.y
}
func (p *position) GetZ() float64 {
	return p.z
}

func main() {
	fmt.Println(NewPosition(1.0, 2.0, 3.0)) // {1 2 3}
	fmt.Println(NewPosition())              // {0 0 0}
}
