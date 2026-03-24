package screen

import (
	"image/color"
	"os"
	"strconv"
	"test/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type gameOverScreen struct {
	gameOverMenuIndexSelected int
	resetFunc                 func()
}

func NewGameOverScreen(
	resetFunc func(),
) *gameOverScreen {
	gs := &gameOverScreen{}
	gs.gameOverMenuIndexSelected = 0
	gs.resetFunc = resetFunc

	return gs
}

var gameOverMenu = []string{
	"Restart",
	"Exit",
}

func (gs *gameOverScreen) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) ||
		inpututil.IsKeyJustPressed(ebiten.KeyJ) {

		gs.gameOverMenuIndexSelected++

		if gs.gameOverMenuIndexSelected > len(gameOverMenu)-1 {
			gs.gameOverMenuIndexSelected = 0
		}

	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) ||
		inpututil.IsKeyJustPressed(ebiten.KeyK) {

		gs.gameOverMenuIndexSelected--

		if gs.gameOverMenuIndexSelected < 0 {
			gs.gameOverMenuIndexSelected = len(gameOverMenu) - 1
		}

	} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if gs.gameOverMenuIndexSelected == 0 {

			gs.resetFunc()
			return nil
		}
		if gs.gameOverMenuIndexSelected == 1 {
			os.Exit(0)
		}
	}

	return nil
}

func (gs *gameOverScreen) Draw(
	screen *ebiten.Image,
	scope int,
) {
	screen.Fill(color.RGBA{0x00, 0x33, 0x00, 0xFF})

	// Центрируем текст GAME OVER
	msg := "GAME OVER"
	bounds := text.BoundString(config.GameFont, msg)
	x := (config.WindowWidth - bounds.Dx()) / 2
	y := config.WindowHeigh/2 - 30
	text.Draw(screen, msg, config.GameFont, x, y, color.White)

	// Score
	scoreMsg := "Score: " + strconv.Itoa(scope)
	bounds = text.BoundString(config.GameFont, scoreMsg)
	x = (config.WindowWidth - bounds.Dx()) / 2
	y = config.WindowHeigh / 2
	text.Draw(screen, scoreMsg, config.GameFont, x, y, color.White)

	defaultColor := color.RGBA{0x80, 0x80, 0x80, 0xFF}
	selectedColor := color.RGBA{0x80, 0x50, 0x80, 0xFF}

	for i, m := range gameOverMenu {
		bounds = text.BoundString(config.GameFont, m)

		color := defaultColor

		if i == gs.gameOverMenuIndexSelected {
			color = selectedColor
		}

		x = (config.WindowWidth - bounds.Dx()) / 2
		y = y + 30
		text.Draw(screen, m, config.GameFont, x, y, color)
	}
	return

}
