package engine

func CheckCollision(b1 *Rectangle, b2 *Rectangle) bool {
	return int(b1.Position.X) < int(b2.Position.X)+b2.Width &&
		int(b1.Position.X)+b1.Width > int(b2.Position.X) &&
		int(b1.Position.Y) < int(b2.Position.Y)+b2.Height &&
		int(b1.Position.Y)+b1.Height > int(b2.Position.Y)
}
