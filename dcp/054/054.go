/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Dropbox.

Sudoku is a puzzle where you're given a partially-filled 9 by 9 grid with
digits. The objective is to fill the grid with the constraint that every row,
column, and box (3 by 3 subgrid) must contain all of the digits from 1 to 9.

Implement an efficient sudoku solver.

*/
package dcp054

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type (
	Row uint8 // 1 thru 9
	Col uint8 // 1 thru 9

	// A square in a sudoku board.
	Square struct {
		row Row
		col Col
	}
)

// Generates squares as the cross product of rows and cols.
func cross(rows []Row, cols []Col) []Square {
	squares := make([]Square, 0, len(rows)*len(cols))
	for _, r := range rows {
		for _, c := range cols {
			squares = append(squares, Square{r, c})
		}
	}
	return squares
}

// Unit is a set of nine squares whose values must be 1-9.
type Unit []Square

// Reports whether unit u contains square s.
func (u Unit) contains(s Square) bool {
	for _, t := range u {
		if s == t {
			return true
		}
	}
	return false
}

var (
	rows    = []Row{1, 2, 3, 4, 5, 6, 7, 8, 9}
	cols    = []Col{1, 2, 3, 4, 5, 6, 7, 8, 9}
	squares = cross(rows, cols)
	units   map[Square][]Unit
	peers   map[Square]Unit
)

func init() {
	// Create set of all units.
	var unitlist []Unit
	for _, col := range cols {
		unitlist = append(unitlist, cross(rows, []Col{col}))
	}
	for _, row := range rows {
		unitlist = append(unitlist, cross([]Row{row}, cols))
	}
	for _, rows := range [][]Row{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} {
		for _, cols := range [][]Col{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} {
			unitlist = append(unitlist, cross(rows, cols))
		}
	}

	// Map each square to the set of units to which it belongs.
	units = make(map[Square][]Unit, len(squares))
	for _, s := range squares {
		var us []Unit
		for _, u := range unitlist {
			if u.contains(s) {
				us = append(us, u)
			}
		}
		units[s] = us
	}

	// Map each square to the set of squares in all of its units.
	peers = make(map[Square]Unit, len(squares))
	for _, s := range squares {
		var v Unit
		for _, u := range units[s] {
			for _, s2 := range u {
				if s2 != s && !v.contains(s2) {
					v = append(v, s2)
				}
			}
		}
		peers[s] = v
	}
}

// Represents a single digit (1-9).
type Digit int8

// Stores the set of digits permitted in a square.
type DigitSet []Digit

// Returns a set that contains the values in s, including d.
// If s already contains d, returns s.
func (s DigitSet) plus(d Digit) DigitSet {
	if s.contains(d) {
		return s
	}
	return append(s, d)
}

// Returns a set that contains the values in s, excluding d.
// If s does not contain d, returns s.
func (s DigitSet) minus(d Digit) DigitSet {
	i := s.indexOf(d)
	if i < 0 {
		return s
	}
	s2 := make(DigitSet, len(s)-1)
	copy(s2, s[:i])
	copy(s2[i:], s[i+1:])
	return s2
}

// Returns the index of d in s. Returns -1 if s does not contain d.
func (s DigitSet) indexOf(d Digit) int {
	for i, d2 := range s {
		if d2 == d {
			return i
		}
	}
	return -1
}

// Reports whether s contains d.
func (s DigitSet) contains(d Digit) bool { return s.indexOf(d) >= 0 }

// String implements the Stringer interface.
func (s DigitSet) String() string {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	var b strings.Builder
	for _, d := range s {
		_, _ = b.WriteString(strconv.Itoa(int(d)))
	}
	return b.String()
}

// Stores the possible digits for each square.
type Constraints map[Square]DigitSet

// Reports whether cons represents a solved puzzle.
func (cons Constraints) solved() bool {
	for _, digits := range cons {
		if len(digits) != 1 {
			return false
		}
	}
	return true
}

// Returns a deep copy of the constraints.
func (cons Constraints) copy() Constraints {
	copy := make(Constraints, len(cons))
	for s, digits := range cons {
		copy[s] = digits
	}
	return copy
}

// Eliminates all the other values (except d) from cons[s] and propagates.
// Returns cons on success, or nil if a contradiction is found.
func assign(cons Constraints, s Square, d Digit) Constraints {
	for _, d2 := range cons[s].minus(d) {
		if eliminate(cons, s, d2) == nil {
			return nil
		}
	}
	return cons
}

// Eliminates d from cons[s]; propagates when possible.
// Returns cons on success, or nil if a contradiction is found.
func eliminate(cons Constraints, s Square, d Digit) Constraints {
	if !cons[s].contains(d) {
		return cons // already eliminated
	}
	cons[s] = cons[s].minus(d)
	// If square s is reduced to one value d2, then eliminate d2 from the peers of s.
	switch len(cons[s]) {
	case 0: // contradiction: removed last value
		return nil
	case 1:
		d2 := cons[s][0]
		for _, s2 := range peers[s] {
			if eliminate(cons, s2, d2) == nil {
				return nil
			}
		}
	}
	// If any unit u is reduced to only one place for d, then put it there.
	for _, u := range units[s] {
		switch dplaces := places(cons, u, d); len(dplaces) {
		case 0: // contradiction: no place for d
			return nil
		case 1: // d can only be in one place in unit; assign it there
			if assign(cons, dplaces[0], d) == nil {
				return nil
			}
		}
	}
	return cons
}

// Returns the set of squares in the unit which may legally hold digit d.
func places(cons Constraints, unit Unit, d Digit) []Square {
	var squares []Square
	for _, s := range unit {
		if cons[s].contains(d) {
			squares = append(squares, s)
		}
	}
	return squares
}

// Depth-first search for solutions, trying all possible digits in an arbitrary
// square and backtracking if a contradiction is found.
func search(cons Constraints) Constraints {
	if cons == nil {
		return nil // failed earlier
	}
	if cons.solved() {
		return cons // already solved
	}
	// Heuristic: try the square with the fewest allowed digits.
	var s Square
	min := 10
	for s2, digits := range cons {
		if n := len(digits); n > 1 && n < min {
			min, s = n, s2
		}
	}
	// Try each possible digit. Backtrack on failure.
	for _, d := range cons[s] {
		if cons := search(assign(cons.copy(), s, d)); cons != nil {
			return cons // solved
		}
	}
	return nil // no solution
}

// Parses a starting grid, which maps a subset of the squares to a single digit.
// Treats the '.' character as a square with no initial constraint.
func parseGrid(grid string) map[Square]Digit {
	chars := make([]rune, 0, len(squares))
	for _, c := range grid {
		if (c >= '0' && c <= '9') || c == '.' {
			chars = append(chars, c)
		}
	}
	if len(chars) != len(squares) {
		return nil // invalid input
	}
	m := make(map[Square]Digit, len(squares))
	for i, s := range squares {
		if chars[i] != '.' {
			m[s] = Digit(chars[i] - '1' + 1)
		}
	}
	return m
}

// Parses a starting grid and applies its constraints.
func parseConstraints(grid string) Constraints {
	cons := make(Constraints, len(squares))
	for _, s := range squares {
		cons[s] = DigitSet{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}
	for s, d := range parseGrid(grid) {
		if assign(cons, s, d) == nil {
			return nil
		}
	}
	return cons
}

// Pretty prints the constraints to standard out.
func printConstraints(cons Constraints) {
	width := 0
	for _, ds := range cons {
		if len(ds) > width {
			width = len(ds)
		}
	}
	hsect := strings.Repeat("-", 3*width)
	hrule := hsect + "+" + hsect + "+" + hsect
	for _, row := range rows {
		for _, col := range cols {
			fmt.Print(center(cons[Square{row, col}].String(), width))
			if col == 3 || col == 6 {
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
		if row == 3 || row == 6 {
			fmt.Println(hrule)
		}
	}
}

// Center-aligns s with spaces to either side.
func center(s string, n int) string {
	if len(s) >= n {
		return s
	}
	m := n - len(s)
	l := m / 2
	r := m - l
	return strings.Repeat(" ", l) + s + strings.Repeat(" ", r)
}

// Returns the grid representation of cons. Assumes cons is solved.
func toGrid(cons Constraints) string {
	var b strings.Builder
	for _, s := range squares {
		b.WriteString(strconv.Itoa(int(cons[s][0])))
	}
	return b.String()
}
