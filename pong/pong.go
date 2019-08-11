package pong

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pong",
		Bounds: pixel.R(0, 0, winWidth, winHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	rightPaddle := newPaddle("right")
	leftPaddle := newPaddle("left")
	ball := newBall()

	gameState := play

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	// Game Loop
	for !win.Closed() {
		win.Clear(colornames.Black)

		switch gameState {
		case play:
			if win.JustPressed(pixelgl.KeySpace) {
				gameState = paused
			}
			if rightPaddle.score == winnngScore || leftPaddle.score == winnngScore {
				gameState = over
			}
			// Updates
			rightPaddle.update(win)
			leftPaddle.aiUpdate(ball)
			ball.update(rightPaddle, leftPaddle)
		case paused:
			if win.JustPressed(pixelgl.KeySpace) {
				gameState = play
			}
		case over:
			t := text.New(win.Bounds().Center().Sub(pixel.V(250, 0)), basicAtlas)
			if win.JustPressed(pixelgl.KeySpace) {
				rightPaddle.score, leftPaddle.score = 0, 0
				t.Clear()
				gameState = play
			}

			var text string
			var color color.Color
			if rightPaddle.score > leftPaddle.score {
				text = "Player 1 Wins!"
				color = colornames.Green
			} else {
				text = "Player 2 Wins!"
				color = colornames.Red
			}

			t.Clear()
			t.Color = color
			fmt.Fprintln(t, text)
			t.Color = colornames.Blue
			fmt.Fprintln(t, "Press space to play again...")

			t.Draw(win, pixel.IM.Scaled(t.Orig, 3))

		}

		// Renders
		rightPaddle.draw(win)
		rightPaddle.drawScore(win, basicAtlas)

		leftPaddle.draw(win)
		leftPaddle.drawScore(win, basicAtlas)

		ball.draw(win)

		// Window update
		win.Update()
	}
}

// Run Pong
func Run() {
	pixelgl.Run(run)
}
