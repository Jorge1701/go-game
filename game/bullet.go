package game

import (
	"game/collision"
	"game/render"
	"math"
)

type Bullet struct {
	x float64
	y float64

	xa float64
	ya float64

	speed float64

	texture *render.Texture

	game *Game
}

func NewBullet(game *Game, x, y, dir float64) *Bullet {
	return &Bullet{
		x:       x,
		y:       y,
		xa:      math.Cos(dir),
		ya:      math.Sin(dir),
		speed:   10,
		texture: render.AllTextures["bullet"],
		game:    game,
	}
}

func (b *Bullet) Update(enemies []*Enemy) {
	for _, e := range enemies {
		if collision.CheckCollision(b, e) {
			b.game.deleteEnemy(e)
			b.game.deleteBullet(b)
			return
		}
	}

	b.x += b.xa * b.speed
	b.y += b.ya * b.speed
}

func (b *Bullet) GetX() float64 {
	return b.x - b.GetWidth()/2
}

func (b *Bullet) GetY() float64 {
	return b.y - b.GetHeight()/2
}

func (b *Bullet) GetTexture() *render.Texture {
	return b.texture
}

func (b *Bullet) GetWidth() float64 {
	return b.texture.Width
}

func (b *Bullet) GetHeight() float64 {
	return b.texture.Height
}