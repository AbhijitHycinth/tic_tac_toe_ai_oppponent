package game

import (
	"fmt"
	"main/grid"
	"main/ui"
)

// Players
type Player struct {
	Name string
	mark grid.Marking
}

// Winning combinations
var winningCombinations = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

// Initialize game
type Game struct {
	grid.Grid
	Player1       Player
	Player2       Player
	isOver        bool
	currentPlayer Player
	*ui.UI
}

// New game
func NewGame(p1, p2 string) *Game {
	return &Game{
		Grid:    grid.NewGrid(),
		Player1: Player{Name: p1, mark: grid.X},
		Player2: Player{Name: p2, mark: grid.O},
		isOver:  false,
		UI:      ui.NewUI(),
	}
}

// Use a circular linked list to alternate player turns
func (g *Game) NextTurn() {
	if g.currentPlayer == g.Player1 {
		g.currentPlayer = g.Player2
	} else {
		g.currentPlayer = g.Player1
	}
	p := g.currentPlayer
	row, col := g.UI.PromptTurn(p.Name)
	g.Grid.Mark(col, row, p.mark)
}

// Start game
func (g *Game) Play() {
	var p Player
	for !g.isOver {
		g.NextTurn()
		g.isOver = g.CheckWin()
		g.Grid.Draw()
	}
	fmt.Print("\nThe winner is:", p.Name)
}

// Check if a player has won
func (g *Game) CheckWin() bool {
	for _, combination := range winningCombinations {
		if g.Grid[combination[0]] == g.Grid[combination[1]] && g.Grid[combination[1]] == g.Grid[combination[2]] && g.Grid[combination[2]] != grid.EMPTY {
			return true
		}
	}
	return false
}
