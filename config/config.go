package config

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

const (
	PointSize    = 60
	WindowWidth  = 1280
	WindowHeight = 860
	MoveTime     = 15

	CountPointX = WindowWidth / PointSize
	CountPointY = WindowHeight / PointSize
)

var GameFont font.Face
var CoinImage *ebiten.Image
