package screen

import (
	"image/color"
	"math"
	"math/rand"
	"strconv"

	"github.com/probeldev/mysupergame/config"
	"github.com/probeldev/mysupergame/model"
	"github.com/probeldev/mysupergame/scope"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type gameScreen struct {
	Point model.Point
	Enemy model.Point
	timer int
	Coins []model.Point
	scope *scope.Scope

	changeScreenFunc func(config.ScreenType)
}

func NewGameScreen(
	changeScreenFunc func(config.ScreenType),
	scope *scope.Scope,
) *gameScreen {
	gs := &gameScreen{}
	gs.Enemy.X = config.CountPointX - 1
	gs.Enemy.Y = config.CountPointY - 1

	gs.Point.X = 0
	gs.Point.Y = 0
	gs.fillingMap()

	gs.changeScreenFunc = changeScreenFunc

	gs.scope = scope
	return gs
}

func (gs *gameScreen) Update() error {

	if gs.needsToMoveEnemy() {
		gs.moveEnemy()
	}

	if gs.isStopGame() {
		gs.changeScreenFunc(config.ScreenTypeGameOver)
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) ||
		inpututil.IsKeyJustPressed(ebiten.KeyH) {
		gs.Point.Left()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) ||
		inpututil.IsKeyJustPressed(ebiten.KeyL) {
		gs.Point.Right()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) ||
		inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		gs.Point.Down()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) ||
		inpututil.IsKeyJustPressed(ebiten.KeyK) {
		gs.Point.Up()
	}

	gs.grabCoin()

	gs.timer++
	return nil
}

func (gs *gameScreen) Draw(
	screenH *ebiten.Image,
) {
	// Обычный HUD с нормальным шрифтом

	screenH.Fill(color.RGBA{0x00, 0x33, 0x00, 0xFF})

	pointSize := float32(config.PointSize)

	scaleX := float64(config.PointSize) / float64(config.CoinImage.Bounds().Dx())
	scaleY := float64(config.PointSize) / float64(config.CoinImage.Bounds().Dy())

	for _, m := range gs.Coins {
		op := &ebiten.DrawImageOptions{}
		// Масштабируем, затем позиционируем
		op.GeoM.Scale(scaleX, scaleY)
		op.GeoM.Translate(float64(m.X)*float64(config.PointSize), float64(m.Y)*float64(config.PointSize))
		screenH.DrawImage(config.CoinImage, op)
	}

	vector.FillRect(screenH, float32(gs.Point.X)*pointSize, float32(gs.Point.Y)*pointSize, pointSize, pointSize, color.RGBA{0x00, 0xFF, 0x00, 0xff}, false)

	vector.FillRect(screenH, float32(gs.Enemy.X)*pointSize, float32(gs.Enemy.Y)*pointSize, pointSize, pointSize, color.RGBA{0xFF, 0x00, 0x00, 0xff}, false)

	scoreText := "Score: " + strconv.Itoa(gs.scope.Value)
	text.Draw(screenH, scoreText, config.GameFont, 10, 30, color.White)
}

func (gs *gameScreen) needsToMoveEnemy() bool {
	return gs.timer%config.MoveTime == 0
}

func (gs *gameScreen) isStopGame() bool {
	deltaX := gs.Point.X - gs.Enemy.X
	deltaY := gs.Point.Y - gs.Enemy.Y

	if deltaX == 0 && deltaY == 0 {
		return true
	}

	return false
}

func (gs *gameScreen) grabCoin() {

	isGrab := false
	index := 0
	for i, c := range gs.Coins {
		if c.X == gs.Point.X && c.Y == gs.Point.Y {
			isGrab = true
			index = i
			break
		}
	}

	if !isGrab {
		return
	}

	gs.scope.Value++

	newCoins := []model.Point{}
	for i, c := range gs.Coins {
		if i == index {
			continue
		}

		newCoins = append(newCoins, c)
	}

	gs.Coins = newCoins

	gs.addRandomCoin()
}

func (gs *gameScreen) moveEnemy() {

	deltaX := gs.Point.X - gs.Enemy.X
	deltaY := gs.Point.Y - gs.Enemy.Y

	if math.Abs(float64(deltaX)) > math.Abs(float64(deltaY)) {
		if deltaX > 0 {
			gs.Enemy.Right()
		} else if deltaX < 0 {
			gs.Enemy.Left()
		}
	} else {
		if deltaY > 0 {
			gs.Enemy.Down()
		} else if deltaY < 0 {
			gs.Enemy.Up()
		}
	}

}

func (gs *gameScreen) addRandomCoin() {
	minX := 0
	maxX := config.CountPointX - 1
	x := rand.Intn(maxX-minX+1) + minX

	minY := 0
	maxY := config.CountPointY - 1
	y := rand.Intn(maxY-minY+1) + minY

	if gs.Point.X == x && gs.Point.Y == y {
		gs.addRandomCoin()
		return
	}

	for _, coin := range gs.Coins {
		if coin.X == x && coin.Y == y {
			gs.addRandomCoin()
			return
		}
	}

	gs.Coins = append(gs.Coins, model.Point{X: x, Y: y})
}

func (gs *gameScreen) fillingMap() {
	gs.Coins = []model.Point{}
	for range 10 {
		gs.addRandomCoin()
	}

}
