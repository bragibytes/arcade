package pong

import (
	"image/color"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

type ball struct {
	pixel.Circle
	velocity pixel.Vec
	color    color.Color
}

func newBall() *ball {
	ball := &ball{
		pixel.C(ballStartingPos, 10.0),
		pixel.V(8, 0),
		colornames.White,
	}

	return ball
}

func (b *ball) draw(win *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Push(b.Center)
	imd.Circle(b.Radius, 0)
	imd.Draw(win)
}

func (b *ball) update(r, l *paddle) {

	// Check collision with right paddle
	if b.Center.X+b.Radius >= r.X-r.w/2 &&
		(b.Center.Y < r.Y+r.h/2 && b.Center.Y > r.Y-r.h/2) {
		b.Center.X -= r.w/2 + b.Radius
		b.velocity.X = -b.velocity.X
		b.velocity.Y = r.areaHit(b)
	}
	// Check collision with left paddle
	if b.Center.X-b.Radius <= l.X+l.w/2 &&
		(b.Center.Y < l.Y+l.h/2 && b.Center.Y > l.Y-l.h/2) {
		b.Center.X += l.w/2 + b.Radius
		b.velocity.X = -b.velocity.X
		b.velocity.Y = l.areaHit(b)
	}

	// Check collision with left wall
	if b.Center.X <= 0 {
		b.Center = ballStartingPos
		b.velocity.X = -b.velocity.X
		r.score++
	}
	// Check collision with right wall
	if b.Center.X >= winWidth {
		b.Center = ballStartingPos
		b.velocity.X = -b.velocity.X
		l.score++
	}

	// Check collision with top and bottom
	if b.Center.Y-b.Radius < 0 || b.Center.Y+b.Radius > winHeight {
		b.velocity.Y = -b.velocity.Y
	}

	b.Center = b.Center.Add(b.velocity)
}
