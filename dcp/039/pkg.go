/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Dropbox.

Conway's Game of Life takes place on an infinite two-dimensional board of square
cells. Each cell is either dead or alive, and at each tick, the following rules
apply:

Any live cell with less than two live neighbours dies.
Any live cell with two or three live neighbours remains living.
Any live cell with more than three live neighbours dies.
Any dead cell with exactly three live neighbours becomes a live cell.
A cell neighbours another cell if it is horizontally, vertically, or diagonally adjacent.

Implement Conway's Game of Life. It should be able to be initialized with a
starting list of live cell coordinates and the number of steps it should run
for. Once initialized, it should print out the board state at each step. Since
it's an infinite board, print out only the relevant coordinates, i.e. from the
top-leftmost live cell to bottom-rightmost live cell.

You can represent a live cell with an asterisk (*) and a dead cell with a dot (.).

*/
package dcp039

import (
	"fmt"
	"io"
)

type Cell struct{ x, y int }

func (c Cell) plus(d Cell) Cell {
	return Cell{
		x: c.x + d.x,
		y: c.y + d.y,
	}
}

type State struct {
	cells    map[Cell]bool // specifies which cells are alive
	min, max Cell          // extent of field
}

func NewState(init []Cell) *State {
	st := State{cells: make(map[Cell]bool)}
	for _, c := range init {
		st.Set(c)
	}
	return &st
}

func (st *State) Extent() (min, max Cell) {
	return st.min, st.max
}

func (st *State) Set(c Cell) {
	st.cells[c] = true
	// Update extent.
	if len(st.cells) == 1 {
		st.min = c
		st.max = c
	} else {
		if st.min.x > c.x {
			st.min.x = c.x
		}
		if st.min.y > c.y {
			st.min.y = c.y
		}
		if st.max.x < c.x {
			st.max.x = c.x
		}
		if st.max.y < c.y {
			st.max.y = c.y
		}
	}
}

func (st *State) IsAlive(c Cell) bool {
	return st.cells[c]
}

func (st *State) NumLiveNeighbors(c Cell) int {
	var neighbors = []Cell{
		{-1, -1}, {0, -1}, {+1, -1}, {+1, 0},
		{+1, +1}, {0, +1}, {-1, +1}, {-1, 0},
	}
	n := 0
	for _, d := range neighbors {
		if st.cells[c.plus(d)] {
			n++
		}
	}
	return n
}

func (st *State) Next() *State {
	st2 := NewState(nil)
	for x := st.min.x - 1; x <= st.max.x+1; x++ {
		for y := st.min.y - 1; y <= st.max.y+1; y++ {
			c := Cell{x, y}
			n := st.NumLiveNeighbors(c)
			if n == 3 || (st.IsAlive(c) && n == 2) {
				st2.Set(c)
			}
		}
	}
	return st2
}

func printState(w io.Writer, st *State, buffer int) {
	min, max := st.Extent()
	for y := min.y - buffer; y <= max.y+buffer; y++ {
		for x := min.x - buffer; x <= max.x+buffer; x++ {
			fmt.Fprintf(w, "%c", cellRune(st.IsAlive(Cell{x, y})))
		}
		fmt.Fprint(w, "\n")
	}
}

func cellRune(alive bool) rune {
	if alive {
		return '#'
	}
	return '.'
}
