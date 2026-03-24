package screen

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/probeldev/mysupergame/config"
)

type startScreen struct {
	startMenuIndexSelected int
	resetFunc              func(bool)
}

func NewStartScreen(
	// TODO: стоит использовать не resetFunc, а сделать отдельную функцию для запуска игры
	resetFunc func(bool),
) *startScreen {
	gs := &startScreen{}
	gs.startMenuIndexSelected = 0
	gs.resetFunc = resetFunc

	return gs
}

var startMenu = []string{
	"New Game",
	"Exit",
}

func (gs *startScreen) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) ||
		inpututil.IsKeyJustPressed(ebiten.KeyJ) {

		gs.startMenuIndexSelected++

		if gs.startMenuIndexSelected > len(startMenu)-1 {
			gs.startMenuIndexSelected = 0
		}

	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) ||
		inpututil.IsKeyJustPressed(ebiten.KeyK) {

		gs.startMenuIndexSelected--

		if gs.startMenuIndexSelected < 0 {
			gs.startMenuIndexSelected = len(startMenu) - 1
		}

	} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if gs.startMenuIndexSelected == 0 {

			gs.resetFunc(false)
			return nil
		}
		if gs.startMenuIndexSelected == 1 {
			os.Exit(0)
		}
	}

	return nil
}

func (gs *startScreen) Draw(
	screen *ebiten.Image,
) {
	screen.Fill(color.RGBA{0x00, 0x1a, 0x33, 0xFF})

	msg := "My Super Game"
	bounds := text.BoundString(config.GameFont, msg)
	x := (config.WindowWidth - bounds.Dx()) / 2
	y := config.WindowHeight/2 - 30
	text.Draw(screen, msg, config.GameFont, x, y, color.White)

	defaultColor := color.RGBA{0x80, 0x80, 0x80, 0xFF}
	selectedColor := color.RGBA{0x80, 0x50, 0x80, 0xFF}

	for i, m := range startMenu {
		bounds = text.BoundString(config.GameFont, m)

		color := defaultColor

		if i == gs.startMenuIndexSelected {
			color = selectedColor
		}

		x = (config.WindowWidth - bounds.Dx()) / 2
		y = y + 30
		text.Draw(screen, m, config.GameFont, x, y, color)
	}

}
