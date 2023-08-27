package utils

import (
	"game/engine"
	"math"
)

func Distance(p1 *engine.Point, p2 *engine.Point) float64 {
	return math.Sqrt(
		math.Pow(p1.X-p2.X, 2) +
			math.Pow(p1.Y-p2.Y, 2))
}

func Direction(from *engine.Point, to *engine.Point) float64 {
	return math.Atan2(to.Y-from.Y, to.X-from.X)
}

func DirectionTo(to *engine.Point) float64 {
	return math.Atan2(to.Y, to.X)
}
