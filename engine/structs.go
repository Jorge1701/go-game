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

func (r *Rectangle) CenterX() float64 {
	return r.Position.X - float64(r.Width)/2
}

func (r *Rectangle) CenterY() float64 {
	return r.Position.Y - float64(r.Height)/2
}
