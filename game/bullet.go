package game

import (
	"game/collision"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	x      float64
	y      float64
	width  int
	height int

	xa float64
	ya float64

	speed float64

	game *Game
}

func NewBullet(game *Game, x, y, dir float64) *Bullet {
	return &Bullet{
		x:      x,
		y:      y,
		width:  4,
		height: 4,
		xa:     math.Cos(dir),
		ya:     math.Sin(dir),
		speed:  10,
		game:   game,
	}
}

func (b *Bullet) Update(enemies []*Enemy) {
	for _, e := range enemies {
		if collision.CheckCollision(b, e) {
			b.game.audioPlayer.PlayFromBytes("enemy_dead")
			b.game.deleteEnemy(e)
			b.game.deleteBullet(b)
			return
		}
	}

	b.x += b.xa * b.speed
	b.y += b.ya * b.speed
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	b.game.imageManager.Draw(screen, "bullet", b.GetX(), b.GetY())
}

func (b *Bullet) GetX() float64 {
	return b.x - float64(b.GetWidth()/2)
}

func (b *Bullet) GetY() float64 {
	return b.y - float64(b.GetHeight()/2)
}

func (b *Bullet) GetWidth() int {
	return b.width
}

func (b *Bullet) GetHeight() int {
	return b.height
}
