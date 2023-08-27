package engine

type Point struct {
	X float64
	Y float64
}

type Rectangle struct {
	Position *Point
	Width    int
	Height   int
}
