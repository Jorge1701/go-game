package game

import (
	"game/audio"
	"game/collision"
	"game/graphics"
	"math"
)

type Bullet struct {
	x float64
	y float64

	xa float64
	ya float64

	speed float64

	image *graphics.Image

	game *Game
}

func NewBullet(game *Game, x, y, dir float64) *Bullet {
	return &Bullet{
		x:     x,
		y:     y,
		xa:    math.Cos(dir),
		ya:    math.Sin(dir),
		speed: 10,
		image: graphics.AllImages["bullet"],
		game:  game,
	}
}

func (b *Bullet) Update(enemies []*Enemy) {
	for _, e := range enemies {
		if collision.CheckCollision(b, e) {
			audio.AllAudios["enemy_dead"].Play()
			b.game.deleteEnemy(e)
			b.game.deleteBullet(b)
			return
		}
	}

	b.x += b.xa * b.speed
	b.y += b.ya * b.speed
}

func (b *Bullet) GetX() float64 {
	return b.x - float64(b.GetWidth()/2)
}

func (b *Bullet) GetY() float64 {
	return b.y - float64(b.GetHeight()/2)
}

func (b *Bullet) GetImage() *graphics.Image {
	return b.image
}

func (b *Bullet) GetWidth() int {
	return b.image.Image.Bounds().Dx()
}

func (b *Bullet) GetHeight() int {
	return b.image.Image.Bounds().Dy()
}
