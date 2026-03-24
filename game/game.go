package game

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"
	"test/config"
	"test/model"
	"test/screen"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type CurrentScreenInterfacer interface {
	Draw(screen *ebiten.Image, scope int)
	Update() error
}

type Game struct {
	Point          model.Point
	Enemy          model.Point
	timer          int
	Coins          []model.Point
	Scope          int
	gameOver       bool
	gameOverScreen CurrentScreenInterfacer
}

func (g *Game) Update() error {

	if g.gameOver {

		g.gameOverScreen.Update()

		return nil

	}

	if g.needsToMoveEnemy() {
		g.moveEnemy()
	}

	if g.isStopGame() {
		g.gameOver = true
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) ||
		inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.Point.Left()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) ||
		inpututil.IsKeyJustPressed(ebiten.KeyL) {
		g.Point.Right()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) ||
		inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		g.Point.Down()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) ||
		inpututil.IsKeyJustPressed(ebiten.KeyK) {
		g.Point.Up()
	}

	g.grabCoin()

	g.timer++

	return nil
}

func (g *Game) needsToMoveEnemy() bool {
	return g.timer%config.MoveTime == 0
}

func (g *Game) isStopGame() bool {
	deltaX := g.Point.X - g.Enemy.X
	deltaY := g.Point.Y - g.Enemy.Y

	if deltaX == 0 && deltaY == 0 {
		return true
	}

	return false
}

func (g *Game) grabCoin() {

	isGrab := false
	index := 0
	for i, c := range g.Coins {
		if c.X == g.Point.X && c.Y == g.Point.Y {
			isGrab = true
			index = i
			break
		}
	}

	if !isGrab {
		return
	}

	g.Scope++

	log.Println("scope", g.Scope)

	newCoins := []model.Point{}
	for i, c := range g.Coins {
		if i == index {
			continue
		}

		newCoins = append(newCoins, c)
	}

	g.Coins = newCoins

	g.addRandomCoin()
}

func (g *Game) moveEnemy() {

	deltaX := g.Point.X - g.Enemy.X
	deltaY := g.Point.Y - g.Enemy.Y

	if math.Abs(float64(deltaX)) > math.Abs(float64(deltaY)) {
		if deltaX > 0 {
			g.Enemy.Right()
		} else if deltaX < 0 {
			g.Enemy.Left()
		}
	} else {
		if deltaY > 0 {
			g.Enemy.Down()
		} else if deltaY < 0 {
			g.Enemy.Up()
		}
	}

}

func (g *Game) addRandomCoin() {
	minX := 0
	maxX := config.CountPointX - 1
	x := rand.Intn(maxX-minX+1) + minX

	minY := 0
	maxY := config.CountPointY - 1
	y := rand.Intn(maxY-minY+1) + minY

	if g.Point.X == x && g.Point.Y == y {
		g.addRandomCoin()
		return
	}

	g.Coins = append(g.Coins, model.Point{X: x, Y: y})
}

func (g *Game) fillingMap() {
	g.Coins = []model.Point{}
	for range 10 {
		g.addRandomCoin()
	}

}

func (g *Game) reset() {

	g.Enemy.X = config.CountPointX - 1
	g.Enemy.Y = config.CountPointY - 1

	g.Point.X = 0
	g.Point.Y = 0

	g.Scope = 0
	g.gameOver = false

	g.gameOverScreen =
		screen.NewGameOverScreen(g.reset)

	g.fillingMap()
}

func NewGame() *Game {
	g := Game{}

	g.reset()

	return &g
}

func (g *Game) Draw(screenH *ebiten.Image) {
	if g.gameOver {
		g.gameOverScreen.Draw(screenH, g.Scope)
		return
	}

	// Обычный HUD с нормальным шрифтом
	scoreText := "Score: " + strconv.Itoa(g.Scope)
	text.Draw(screenH, scoreText, config.GameFont, 10, 30, color.White)

	pointSize := float32(config.PointSize)

	for _, m := range g.Coins {
		vector.FillRect(screenH, float32(m.X)*pointSize, float32(m.Y)*pointSize, pointSize, pointSize, color.RGBA{0xFF, 0xFF, 0x00, 0xff}, false)
	}

	vector.FillRect(screenH, float32(g.Point.X)*pointSize, float32(g.Point.Y)*pointSize, pointSize, pointSize, color.RGBA{0x00, 0xFF, 0x00, 0xff}, false)

	vector.FillRect(screenH, float32(g.Enemy.X)*pointSize, float32(g.Enemy.Y)*pointSize, pointSize, pointSize, color.RGBA{0xFF, 0x00, 0x00, 0xff}, false)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.WindowWidth, config.WindowHeigh
}
