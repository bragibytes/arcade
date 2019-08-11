package pong

import "github.com/faiface/pixel"

const (
	winWidth, winHeight     = 800, 600
	winnngScore         int = 5
)

var (
	ballStartingPos = pixel.V(winWidth/2, winHeight/2)
)

type gameState int

const (
	paused gameState = iota
	play
	over
)
