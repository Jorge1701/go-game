package game

import (
	"game/engine"
	"game/graphics"
	"math"
)

type Bullet struct {
	drawable *graphics.Drawable

	xa float64
	ya float64

	speed float64

	game *Game
}

func NewBullet(game *Game, x, y, dir float64) *Bullet {
	return &Bullet{
		drawable: &graphics.Drawable{
			Rect: &engine.Rectangle{
				Position: &engine.Point{
					X: x,
					Y: y,
				},
				Width:  4,
				Height: 4,
			},
			ImageAlias: "bullet",
		},
		xa:    math.Cos(dir),
		ya:    math.Sin(dir),
		speed: 10,
		game:  game,
	}
}

func (b *Bullet) Update(enemies []*Enemy) {
	for _, e := range enemies {
		if engine.CheckCollision(b.drawable.Rect, e.drawable.Rect) {
			b.game.audioPlayer.PlayFromBytes("enemy_dead")
			b.game.deleteEnemy(e)
			b.game.deleteBullet(b)
			return
		}
	}

	b.drawable.Rect.Position.X += b.xa * b.speed
	b.drawable.Rect.Position.Y += b.ya * b.speed
}
