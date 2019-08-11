package pong

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel/text"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type paddle struct {
	pixel.Vec
	color color.Color
	speed float64
	w, h  float64
	score int
	side  string
}

func newPaddle(side string) *paddle {

	var pos pixel.Vec
	if side == "left" {
		pos = pixel.V(winWidth/2, winHeight/2).Sub(pixel.V(winWidth/2-20, 0))
	} else if side == "right" {
		pos = pixel.V(winWidth/2, winHeight/2).Add(pixel.V(winWidth/2-20, 0))
	}

	paddle := &paddle{
		pos,
		colornames.White,
		5.0,
		10,
		70,
		0,
		side,
	}

	return paddle
}

func (p *paddle) rect() pixel.Rect {
	return pixel.R(p.X-p.w/2, p.Y-p.h/2, p.X+p.w/2, p.Y+p.h/2)
}

func (p *paddle) draw(win *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Color = p.color
	imd.Push(p.rect().Min)
	imd.Push(p.rect().Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func (p *paddle) drawScore(win *pixelgl.Window, atlas *text.Atlas) {
	var pos pixel.Vec
	if p.side == "left" {
		pos = pixel.V(30, winHeight-30)
	} else if p.side == "right" {
		pos = pixel.V(winWidth-30, winHeight-30)
	}
	t := text.New(pos, atlas)
	fmt.Fprintf(t, "%d", p.score)
	t.Draw(win, pixel.IM.Scaled(t.Orig, 2))
}

func (p *paddle) areaHit(ball *ball) float64 {
	var a, f pixel.Vec
	if p.side == "right" {
		f, a = p.rect().Vertices()[0], p.rect().Vertices()[1]
	} else if p.side == "left" {
		a, f = p.rect().Vertices()[2], p.rect().Vertices()[3]
	}
	b := pixel.Lerp(a, f, .2)
	c := pixel.Lerp(a, f, .4)
	d := pixel.Lerp(a, f, .6)
	e := pixel.Lerp(a, f, .8)

	switch y := ball.Center.Y; {
	case y < a.Y && y > b.Y:
		return 10
	case y < b.Y && y > c.Y:
		return 5
	case y < c.Y && y > d.Y:
		return ball.velocity.Y
	case y < d.Y && y > e.Y:
		return -5
	case y < e.Y && y > f.Y:
		return -10
	default:
		return ball.velocity.Y
	}
}

func (p *paddle) update(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyUp) {
		p.Y += p.speed
	}
	if win.Pressed(pixelgl.KeyDown) {
		p.Y -= p.speed
	}
}

func (p *paddle) aiUpdate(ball *ball) {
	if ball.Center.Y > p.Y {
		p.Y += p.speed
	} else if ball.Center.Y < p.Y {
		p.Y -= p.speed
	}
}
