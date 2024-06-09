package grid

import (
	"fmt"
)

type Marking int8

const (
	EMPTY Marking = iota
	X
	O
)

var markingValue = map[Marking]string{
	X:     "X",
	O:     "O",
	EMPTY: "  ",
}

func (m Marking) Value() string {
	return markingValue[m]
}

type Grid []Marking

// NewGrid returns a new grid with all cells initialized to EMPTY.
func NewGrid() Grid {
	g := make(Grid, 9)
	return g
}

// Draw prints the grid to standard output.
func (g Grid) Draw() {
	delimiter := "|"
	j := 0
	for _, val := range g {
		if j == 0 {
			fmt.Println()
		}
		fmt.Printf("\t%s", val.Value())
		if j < 2 {
			fmt.Print(delimiter)
		}
		j = (j + 1) % 3
	}
}

// Mark marks the given coordinates on the grid with the specified marking.
func (g Grid) Mark(col, row int8, m Marking) error {
	if col < 0 || col > 2 || row < 0 || row > 2 {
		return fmt.Errorf("invalid coordinate: %d,%d", col, row)
	}

	cellNumber := (row * 3) + col
	(g)[cellNumber] = m
	return nil
}
