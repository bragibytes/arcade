package main

import (
	"deadgopher/arcade/breakout"
	"deadgopher/arcade/pong"
	"deadgopher/arcade/tetris"
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) == 1 {
		switch args[0] {
		case "pong":
			pong.Run()
		case "breakout":
			breakout.Run()
		case "tetris":
			tetris.Run()
		default:
			fmt.Println("Could not find that game")
		}
	} else {
		fmt.Println("Please specify one game.")
		fmt.Println("e.g. arcade pong\ne.g. arcade breakout \ne.g. arcade tetris")
	}

}
