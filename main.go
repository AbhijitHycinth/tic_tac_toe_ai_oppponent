package main

import (
	"fmt"
	"main/game"
)

func main() {
	// take user input for two player names
	var p1 string
	fmt.Print("Enter player 1 name: ")
	fmt.Scanln(&p1)

	var p2 string
	fmt.Print("Enter player 2 name: ")
	fmt.Scanln(&p2)

	g := game.NewGame(p1, p2)
	g.Play()
}
