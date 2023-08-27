package game

import (
	"fmt"
	"game/audio"
	"game/configuration"
	"game/engine"
	"game/fonts"
	"game/image"
	"image/color"
	"math/rand"
	"slices"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	terrain *Terrain

	fontManager  *fonts.FontManager
	imageManager *image.ImageManager
	audioPlayer  *audio.AudioPlayer
	player       *Player
	enemies      []*Enemy
	bullets      []*Bullet

	stage         int
	killedEnemies int
	maxEnemyCount int

	IsGameOver bool

	lastUpdate int64
}

func NewGame() (*Game, error) {
	// Create font manager
	fontManager, err := fonts.NewFontManager()
	if err != nil {
		return nil, err
	}

	// Create image manager
	imageManager, err := image.NewImageManager()
	if err != nil {
		return nil, err
	}

	// Create audio player
	audioPlayer, err := audio.NewAudioPlayer()
	if err != nil {
		return nil, err
	}

	// Creating game object
	game := &Game{
		fontManager:   fontManager,
		imageManager:  imageManager,
		audioPlayer:   audioPlayer,
		enemies:       []*Enemy{},
		bullets:       []*Bullet{},
		stage:         1,
		killedEnemies: 0,
		maxEnemyCount: configuration.StartingEnemyCount,
		IsGameOver:    false,
	}

	game.terrain = NewTerrain(game)

	game.player = NewPlayer(
		game,
		configuration.Width/2,
		configuration.Height/2,
	)

	return game, nil
}

func (g *Game) Update() error {
	// Calculate deltaTime time since last update
	currT := time.Now().UnixMilli()
	deltaTime := currT - g.lastUpdate
	g.lastUpdate = currT

	if g.IsGameOver {
		return nil
	}

	g.player.Update()

	for _, e := range g.enemies {
		e.Update()
	}

	for _, b := range g.bullets {
		b.Update(deltaTime, g.enemies)
	}

	for len(g.enemies) < g.maxEnemyCount {
		g.createEnemy()
	}

	fmt.Printf("stage %d - enemies %d - killed %d - next %d - bullets %d\n",
		g.stage,
		len(g.enemies),
		g.killedEnemies,
		g.nextStage(),
		len(g.bullets),
	)

	g.imageManager.XOffSet = g.player.drawable.Rect.Position.X - configuration.Width/2
	g.imageManager.YOffSet = g.player.drawable.Rect.Position.Y - configuration.Height/2

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	g.terrain.Draw(screen)

	g.imageManager.Draw(screen, g.player.drawable)

	for _, e := range g.enemies {
		g.imageManager.Draw(screen, e.drawable)
	}

	for _, b := range g.bullets {
		g.imageManager.Draw(screen, b.drawable)
	}

	if g.IsGameOver {
		g.fontManager.ShowText(screen, "Game Over!",
			"principal",
			configuration.Width/2,
			configuration.Height/2,
			color.Black,
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return configuration.Width, configuration.Height
}

func (g *Game) createBullet(position *engine.Point, dir float64) {
	g.bullets = append(g.bullets, NewBullet(g, position.X, position.Y, dir))
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

	e, err := NewEnemy(g, x+g.imageManager.XOffSet, y+g.imageManager.YOffSet)

	if err != nil {
		panic(fmt.Sprintf("Error generating enemy: %v", err))
	}

	g.enemies = append(g.enemies, e)
}

func (g *Game) nextStage() int {
	return configuration.StartingEnemyCount*g.stage + g.stage
}

func (g *Game) deleteEnemy(enemyToDelete *Enemy) {
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

func (g *Game) deleteBullet(bulletToDelete *Bullet) {
	for i, b := range g.bullets {
		if b == bulletToDelete {
			g.bullets = slices.Delete(g.bullets, i, i+1)
			return
		}
	}
}

func (g *Game) GameOver() {
	g.IsGameOver = true
	g.audioPlayer.PlayFromBytes("game_over")
}
