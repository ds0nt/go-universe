package universe

import "fmt"

type Position struct {
	X float64
	Y float64
	Z float64
}

func (p *Position) String() string {
	return fmt.Sprintf("(%f, %f, %f)", p.X, p.Y, p.Z)
}

func (p *Position) SetValue(p2 Position) {
	p.X = p2.X
	p.Y = p2.Y
	p.Z = p2.Z
}

func newPosition() *Position {
	return &Position{}
}
