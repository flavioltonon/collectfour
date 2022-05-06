package main

import (
	"collectfour/internal/game"
)

func main() {
	g := game.New(
		game.NewPlayer("Player 1", "1"),
		game.NewPlayer("Player 2", "2"),
	)

	if err := g.Run(); err != nil {
		panic(err)
	}
}
