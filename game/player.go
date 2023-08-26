package game

import (
	"game/audio"
	"game/render"
	"game/utils"
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var lastFire = int64(0)

type Player struct {
	x float64
	y float64

	speed    float64
	fireRate int64

	health int

	texture *render.Texture

	game *Game
}

func NewPlayer(game *Game, x, y float64) *Player {
	return &Player{
		x:        x,
		y:        y,
		speed:    1,
		fireRate: 500,
		health:   5,
		texture:  render.AllTextures["player"],
		game:     game,
	}
}

func (p *Player) Update() {
	keys := sdl.GetKeyboardState()

	mouseX, mouseY, mouseState := sdl.GetMouseState()

	if mouseState == sdl.ButtonLMask {
		currT := time.Now().UnixMilli()
		if currT-lastFire > p.fireRate {
			lastFire = currT
			dirToMouse := utils.Direction(&utils.Point{X: p.x, Y: p.y}, &utils.Point{X: float64(mouseX), Y: float64(mouseY)})
			p.game.createBullet(p.x, p.y, dirToMouse)
			audio.AllAudios["shot"].Play()
		}
	}

	movePoint := &utils.Point{}

	if keys[sdl.SCANCODE_A] == 1 {
		movePoint.X--
	}
	if keys[sdl.SCANCODE_D] == 1 {
		movePoint.X++
	}
	if keys[sdl.SCANCODE_W] == 1 {
		movePoint.Y--
	}
	if keys[sdl.SCANCODE_S] == 1 {
		movePoint.Y++
	}

	if movePoint.X != 0 || movePoint.Y != 0 {
		moveDirection := utils.DirectionTo(movePoint)

		p.x += math.Cos(moveDirection) * p.speed
		p.y += math.Sin(moveDirection) * p.speed
	}
}

func (p *Player) GetHit() {
	p.health--
	audio.AllAudios["player_hit"].Play()

	if p.health <= 0 {
		p.game.GameOver()
	}
}

func (p *Player) GetX() float64 {
	return p.x - p.GetWidth()/2
}

func (p *Player) GetY() float64 {
	return p.y - p.GetHeight()/2
}

func (p *Player) GetTexture() *render.Texture {
	return p.texture
}

func (p *Player) GetWidth() float64 {
	return p.texture.Width
}

func (p *Player) GetHeight() float64 {
	return p.texture.Height
}
