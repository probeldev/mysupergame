package main

import (
	"log"
	"test/config"
	"test/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts" // стандартный шрифт из примеров ebiten
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Глобальная переменная для шрифта

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
}

func main() {
	ebiten.SetWindowSize(config.WindowWidth, config.WindowHeigh)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
