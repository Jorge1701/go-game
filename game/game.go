package game

import (
	"fmt"
	"game/configuration"
	"game/render"
	"math/rand"
	"slices"
)

type Game struct {
	renderer *render.Renderer

	player  *Player
	enemies []*Enemy

	stage         int
	killedEnemies int
	maxEnemyCount int

	IsGameOver bool
}

func NewGame(renderer *render.Renderer) (g *Game, err error) {
	g = &Game{
		renderer:      renderer,
		enemies:       []*Enemy{},
		stage:         1,
		killedEnemies: 0,
		maxEnemyCount: configuration.StartingEnemyCount,
		IsGameOver:    false,
	}

	p, err := NewPlayer(
		g,
		configuration.Width/2,
		configuration.Height/2,
		render.AllTextures["player"],
	)
	if err != nil {
		return &Game{}, fmt.Errorf("Error creating new player: %v", err)
	}
	g.player = p

	return g, err
}

func (g *Game) Update() {
	g.player.Update()

	for _, e := range g.enemies {
		e.Update()
	}

	for len(g.enemies) < g.maxEnemyCount {
		g.createEnemy()
	}

	fmt.Printf("stage %d - enemies %d - killed %d - next %d\n",
		g.stage,
		len(g.enemies),
		g.killedEnemies,
		g.nextStage(),
	)
}

func (g *Game) Render() {
	g.renderer.RenderDrawable(g.player)

	for _, e := range g.enemies {
		g.renderer.RenderDrawable(e)
	}
}

func (g *Game) createEnemy() {
	var x, y float64

	side := rand.Intn(4)

	switch side {
	case 0: // top
		x = float64(rand.Intn(configuration.Width))
		y = -50
	case 1: // right
		x = configuration.Width + 50
		y = float64(rand.Intn(configuration.Height))
	case 2: // bottom
		x = float64(rand.Intn(configuration.Width))
		y = configuration.Height + 50
	case 3: // left
		x = -50
		y = float64(rand.Intn(configuration.Height))
	}

	e, err := NewEnemy(g, x, y, render.AllTextures["enemy"])

	if err != nil {
		panic(fmt.Sprintf("Error generating enemy: %v", err))
	}

	g.enemies = append(g.enemies, e)
}

func (g *Game) nextStage() int {
	return configuration.StartingEnemyCount*g.stage + g.stage
}

func (g *Game) DeleteEnemy(enemyToDelete *Enemy) {
	for i, e := range g.enemies {
		if e == enemyToDelete {
			g.enemies = slices.Delete(g.enemies, i, i+1)

			g.killedEnemies++
			if g.killedEnemies%g.nextStage() == 0 {
				g.stage++
				g.maxEnemyCount++
			}

			return
		}
	}
}

func (g *Game) GameOver() {
	g.IsGameOver = true
}
