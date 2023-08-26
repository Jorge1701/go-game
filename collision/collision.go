package collision

type Rectangle struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

type Boundary interface {
	GetX() float64
	GetY() float64
	GetWidth() float64
	GetHeight() float64
}

func CheckCollision(b1 Boundary, b2 Boundary) bool {
	return b1.GetX() < b2.GetX()+b2.GetWidth() &&
		b1.GetX()+b1.GetWidth() > b2.GetX() &&
		b1.GetY() < b2.GetY()+b2.GetHeight() &&
		b1.GetY()+b1.GetHeight() > b2.GetY()
}

func (r *Rectangle) GetX() float64 {
	return r.X
}

func (r *Rectangle) GetY() float64 {
	return r.Y
}

func (r *Rectangle) GetWidth() float64 {
	return r.Width
}

func (r *Rectangle) GetHeight() float64 {
	return r.Height
}
