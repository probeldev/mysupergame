package main

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/probeldev/mysupergame/config"
	"github.com/probeldev/mysupergame/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts" // стандартный шрифт из примеров ebiten
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed image/coin.png
var CoinPNG []byte

func init() {
	// Загружаем шрифт один раз при старте
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	config.GameFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	coinImage, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(CoinPNG))
	if err != nil {
		log.Fatal(err)
	}
	config.CoinImage = coinImage
}

func main() {
	ebiten.SetWindowSize(config.WindowWidth, config.WindowHeight)
	ebiten.SetWindowTitle("My Super Game")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
