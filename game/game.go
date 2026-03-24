package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/probeldev/mysupergame/config"
	"github.com/probeldev/mysupergame/scope"
	"github.com/probeldev/mysupergame/screen"
)

type CurrentScreenInterfacer interface {
	Draw(screen *ebiten.Image)
	Update() error
}

type Game struct {
	Scope          scope.Scope
	gameOver       bool
	gameOverScreen CurrentScreenInterfacer
	gameScreen     CurrentScreenInterfacer
	start          bool
	startScreen    CurrentScreenInterfacer
}

func (g *Game) Update() error {

	if g.gameOver {

		g.gameOverScreen.Update()

		return nil

	}

	if g.start {

		g.startScreen.Update()

		return nil

	}

	g.gameScreen.Update()

	return nil
}

func (g *Game) gameOverFunc() {
	g.gameOver = true
}

func (g *Game) reset(isStartScreen bool) {

	g.Scope = scope.Scope{}
	g.gameOver = false
	g.start = isStartScreen

	g.gameOverScreen =
		screen.NewGameOverScreen(g.reset, &g.Scope)

	g.gameScreen = screen.NewGameScreen(g.gameOverFunc, &g.Scope)
	g.startScreen = screen.NewStartScreen(g.reset)

}

func NewGame() *Game {
	g := Game{}

	g.reset(true)

	return &g
}

func (g *Game) Draw(screenH *ebiten.Image) {
	if g.start {
		g.startScreen.Draw(screenH)
		return
	}

	if g.gameOver {
		g.gameOverScreen.Draw(screenH)
		return
	}

	g.gameScreen.Draw(screenH)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.WindowWidth, config.WindowHeight
}
