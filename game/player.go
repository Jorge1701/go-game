package game

import (
	"game/render"
	"game/utils"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	x float64
	y float64

	speed float64

	health int

	texture *render.Texture

	game *Game
}

func NewPlayer(game *Game, x, y float64, texture *render.Texture) (*Player, error) {
	return &Player{
		x:       x,
		y:       y,
		speed:   2,
		health:  5,
		texture: texture,
		game:    game,
	}, nil
}

func (p *Player) Update() {
	keys := sdl.GetKeyboardState()

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

	if p.health <= 0 {
		p.game.GameOver()
	}
}

func (p *Player) GetX() float64 {
	return p.x
}

func (p *Player) GetY() float64 {
	return p.y
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
