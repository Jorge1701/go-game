package engine

func CheckCollision(b1 *Rectangle, b2 *Rectangle) bool {
	return int(b1.CenterX()) < int(b2.CenterX())+b2.Width &&
		int(b1.CenterX())+b1.Width > int(b2.CenterX()) &&
		int(b1.CenterY()) < int(b2.CenterY())+b2.Height &&
		int(b1.CenterY())+b1.Height > int(b2.CenterY())
}
