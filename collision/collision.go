package collision

type Rectangle struct {
	X      float64
	Y      float64
	Width  int
	Height int
}

type Boundary interface {
	GetX() float64
	GetY() float64
	GetWidth() int
	GetHeight() int
}

func CheckCollision(b1 Boundary, b2 Boundary) bool {
	return int(b1.GetX()) < int(b2.GetX())+b2.GetWidth() &&
		int(b1.GetX())+b1.GetWidth() > int(b2.GetX()) &&
		int(b1.GetY()) < int(b2.GetY())+b2.GetHeight() &&
		int(b1.GetY())+b1.GetHeight() > int(b2.GetY())
}

func (r *Rectangle) GetX() float64 {
	return r.X
}

func (r *Rectangle) GetY() float64 {
	return r.Y
}

func (r *Rectangle) GetWidth() int {
	return r.Width
}

func (r *Rectangle) GetHeight() int {
	return r.Height
}
