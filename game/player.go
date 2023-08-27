package game

import (
	"game/utils"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var lastFire = int64(0)

type Player struct {
	x      float64
	y      float64
	width  int
	height int

	speed    float64
	fireRate int64

	health int

	game *Game
}

func NewPlayer(game *Game, x, y float64) *Player {
	return &Player{
		x:        x,
		y:        y,
		width:    19,
		height:   40,
		speed:    1,
		fireRate: 500,
		health:   5,
		game:     game,
	}
}

func (p *Player) Update() {
	// Check mouse input
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		mouseX, mouseY := ebiten.CursorPosition()
		currT := time.Now().UnixMilli()

		if currT-lastFire > p.fireRate {
			lastFire = currT
			dirToMouse := utils.Direction(&utils.Point{X: p.x, Y: p.y}, &utils.Point{X: float64(mouseX), Y: float64(mouseY)})
			p.game.createBullet(p.x, p.y, dirToMouse)
			p.game.audioPlayer.PlayFromBytes("shot")
		}
	}

	// Checkout keyboard input
	keys := inpututil.AppendPressedKeys([]ebiten.Key{})

	movePoint := &utils.Point{}

	for _, key := range keys {
		if key == ebiten.KeyA {
			movePoint.X--
		}
		if key == ebiten.KeyD {
			movePoint.X++
		}
		if key == ebiten.KeyW {
			movePoint.Y--
		}
		if key == ebiten.KeyS {
			movePoint.Y++
		}
	}

	if movePoint.X != 0 || movePoint.Y != 0 {
		moveDirection := utils.DirectionTo(movePoint)

		p.x += math.Cos(moveDirection) * p.speed
		p.y += math.Sin(moveDirection) * p.speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.game.imageManager.Draw(screen, "player", p.GetX(), p.GetY())
}

func (p *Player) GetHit() {
	p.health--
	p.game.audioPlayer.PlayFromBytes("player_hit")

	if p.health <= 0 {
		p.game.GameOver()
	}
}

func (p *Player) GetX() float64 {
	return p.x - float64(p.GetWidth()/2)
}

func (p *Player) GetY() float64 {
	return p.y - float64(p.GetHeight()/2)
}

func (p *Player) GetWidth() int {
	return p.width
}

func (p *Player) GetHeight() int {
	return p.height
}
