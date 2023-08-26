package utils

import "math"

type Point struct {
	X float64
	Y float64
}

type Position interface {
	GetX() float64
	GetY() float64
}

func Distance(p1 Position, p2 Position) float64 {
	return math.Sqrt(
		math.Pow(p1.GetX()-p2.GetX(), 2) +
			math.Pow(p1.GetY()-p2.GetY(), 2))
}

func Direction(from Position, to Position) float64 {
	return math.Atan2(to.GetY()-from.GetY(), to.GetX()-from.GetX())
}

func DirectionTo(to *Point) float64 {
	return math.Atan2(to.Y, to.X)
}

func (p *Point) GetX() float64 {
	return p.X
}

func (p *Point) GetY() float64 {
	return p.Y
}
