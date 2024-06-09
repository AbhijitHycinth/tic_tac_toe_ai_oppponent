package ui

import (
	"fmt"
)

type UI struct {
}

// NewUI returns a new instance of the UI
func NewUI() *UI {
	return &UI{}
}

// PromptTurn prompts the player for their next move
func (ui *UI) PromptTurn(playerName string) (x, y int8) {
	fmt.Printf("\n%s it is your turn:", playerName)
	fmt.Printf("\n enter cell coordinates:")
	fmt.Scanln(&x, &y)

	return x, y
}
