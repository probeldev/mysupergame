package config

import "golang.org/x/image/font"

const (
	PointSize   = 30
	WindowWidth = 640
	WindowHeigh = 480
	MoveTime    = 15

	CountPointX = WindowWidth / PointSize
	CountPointY = WindowHeigh / PointSize
)

var GameFont font.Face
