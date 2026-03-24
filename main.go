package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	pointSize    = 30
	windowWidth  = 640
	windowHeight = 480
	moveTime     = 15

	countPointX = windowWidth / pointSize
	countPointY = windowHeight / pointSize
)

type Point struct {
	X int
	Y int
}

func (p *Point) Left() {
	if p.X == 0 {
		return
	}
	p.X--
}

func (p *Point) Right() {
	if p.X == countPointX-1 {
		return
	}
	p.X++
}

func (p *Point) Up() {
	if p.Y == 0 {
		return
	}
	p.Y--
}

func (p *Point) Down() {

	if p.Y == countPointY-1 {

		return
	}
	p.Y++
}

type Game struct {
	Point Point
	Enemy Point
	timer int
	Coins []Point
	Scope int
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.reset()
	}

	if g.needsToMoveEnemy() {
		g.moveEnemy()
	}

	if g.isStopGame() {
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
	return g.timer%moveTime == 0
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

	newCoins := []Point{}
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
	maxX := countPointX - 1
	x := rand.Intn(maxX-minX+1) + minX

	minY := 0
	maxY := countPointY - 1
	y := rand.Intn(maxY-minY+1) + minY

	if g.Point.X == x && g.Point.Y == y {
		g.addRandomCoin()
		return
	}

	g.Coins = append(g.Coins, Point{X: x, Y: y})
}

func (g *Game) fillingMap() {
	g.Coins = []Point{}
	for range 10 {
		g.addRandomCoin()
	}

}

func (g *Game) reset() {

	g.Enemy.X = countPointX - 1
	g.Enemy.Y = countPointY - 1

	g.Point.X = 0
	g.Point.Y = 0

	g.fillingMap()
}

func newGame() *Game {
	g := Game{}

	g.reset()

	return &g
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Scope: "+strconv.Itoa(g.Scope))

	for _, m := range g.Coins {
		vector.FillRect(screen, float32(m.X*pointSize), float32(m.Y*pointSize), pointSize, pointSize, color.RGBA{0xFF, 0xFF, 0x00, 0xff}, false)
	}

	vector.FillRect(screen, float32(g.Point.X*pointSize), float32(g.Point.Y*pointSize), pointSize, pointSize, color.RGBA{0x00, 0xFF, 0x00, 0xff}, false)

	vector.FillRect(screen, float32(g.Enemy.X*pointSize), float32(g.Enemy.Y*pointSize), pointSize, pointSize, color.RGBA{0xFF, 0x00, 0x00, 0xff}, false)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}
