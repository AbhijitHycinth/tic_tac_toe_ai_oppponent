package grid

import "testing"

// Test the Mark() method
func TestMark(t *testing.T) {
	// given
	g := NewGrid()
	// when valid coordinates are used
	{
		err := g.Mark(0, 1, X) // should succeed
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if g[3] != X {
			t.Errorf("expected cell at (0, 1) to be marked with X, but it is not")
		}
	}
	// When coordinates are out of bounds
	{
		err := g.Mark(3, 2, O) // should fail
		if err == nil {
			t.Errorf("expected an error for invalid coordinates, got none")
		}
	}
}
