package dcp054

import (
	"reflect"
	"testing"
)

func TestInitValues(t *testing.T) {
	if len(squares) != 81 {
		t.Errorf("len(squares) is %d, want 81", len(squares))
	}
	for _, s := range squares {
		if len(units[s]) != 3 {
			t.Errorf("len(units[%v]) is %d, want 3", s, len(units[s]))
		}
	}
	for _, s := range squares {
		if len(peers[s]) != 20 {
			t.Errorf("len(peers[%v]) is %d, want 20", s, len(peers[s]))
		}
	}
	if s, expected := sq("C2"), []Unit{
		sqs("A2", "B2", "C2", "D2", "E2", "F2", "G2", "H2", "I2"),
		sqs("C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9"),
		sqs("A1", "A2", "A3", "B1", "B2", "B3", "C1", "C2", "C3"),
	}; !reflect.DeepEqual(units[s], expected) {
		t.Fatalf("units[%v] is %+v, want %+v", s, units[s], expected)
	}
	if s, expected := sq("C2"), Unit(sqs(
		"A2", "B2", "D2", "E2", "F2", "G2", "H2", "I2",
		"C1", "C3", "C4", "C5", "C6", "C7", "C8", "C9",
		"A1", "A3", "B1", "B3",
	)); !reflect.DeepEqual(peers[s], expected) {
		t.Fatalf("peers[%v] is %+v, want %+v", s, peers[s], expected)
	}
}

func TestSolve(t *testing.T) {
	var tests = []struct {
		grid     string
		solution string
	}{
		{
			"..3.2.6..9..3.5..1..18.64....81.29..7.......8..67.82....26.95..8..2.3..9..5.1.3..",
			"483921657967345821251876493548132976729564138136798245372689514814253769695417382",
		},
		{
			"4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......",
			"417369825632158947958724316825437169791586432346912758289643571573291684164875293",
		},
	}
	for _, tt := range tests {
		result := search(parseConstraints(tt.grid))
		solution := toGrid(result)
		if solution != tt.solution {
			t.Fatalf("solve(%q) returned %q, want %q", tt.grid, solution, tt.solution)
		}
	}
}

// Shorthand for creating a square from a string, e.g. "B7".
func sq(s string) Square {
	return Square{row: Row(s[0] - 'A' + 1), col: Col(s[1] - '1' + 1)}
}

// Applies sq to a sequence of strings.
func sqs(a ...string) []Square {
	b := make([]Square, len(a))
	for i, s := range a {
		b[i] = sq(s)
	}
	return b
}
