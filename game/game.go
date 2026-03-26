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
	Scope         scope.Scope
	CurrentScreen CurrentScreenInterfacer
}

func (g *Game) Update() error {

	return g.CurrentScreen.Update()
}

func (g *Game) ChangeScreen(newScreen config.ScreenType) {
	switch newScreen {
	case config.ScreenTypeStart:
		g.CurrentScreen = screen.NewStartScreen(
			g.ChangeScreen,
		)
	case config.ScreenTypeGame:
		g.Scope.Value = 0
		g.CurrentScreen = screen.NewGameScreen(
			g.ChangeScreen,
			&g.Scope,
		)

	case config.ScreenTypeGameOver:
		g.CurrentScreen = screen.NewGameOverScreen(

			g.ChangeScreen,
			&g.Scope,
		)
	}
}

func NewGame() *Game {
	g := Game{}
	g.Scope = scope.Scope{}
	g.ChangeScreen(config.ScreenTypeStart)

	return &g
}

func (g *Game) Draw(screenH *ebiten.Image) {
	g.CurrentScreen.Draw(screenH)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.WindowWidth, config.WindowHeight
}
