package model

import "test/config"

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
	if p.X == config.CountPointX-1 {
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
	if p.Y == config.CountPointY-1 {

		return
	}
	p.Y++
}
