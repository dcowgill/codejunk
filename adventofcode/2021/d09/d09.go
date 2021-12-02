package d09

import (
	"adventofcode2021/lib"
	"sort"
)

const MAX_DEPTH = 9

func Run() {
	lib.Run(9, part1, part2)
}

func part1() int64 {
	risk := 0
	g := parseInput(realInput)
	for r := 0; r < g.nrows(); r++ {
		for c := 0; c < g.ncols(); c++ {
			if deeperThanNeighbors(g, r, c) {
				risk += int(g.depthAt(r, c) + 1)
			}
		}
	}
	return int64(risk)
}

func deeperThanNeighbors(g *grid, r, c int) bool {
	d := g.depthAt(r, c)
	for _, cell := range g.neighbors(r, c) {
		if d >= g.depthAt(cell.row, cell.col) {
			return false
		}
	}
	return true
}

func part2() int64 {
	g := parseInput(realInput)

	basins := lib.Make2DArray[int](g.nrows(), g.ncols())
	currBasin := 1

	// Uses depth-first search to simulate flow both up and down.
	var flow func(r, c int) bool
	flow = func(r, c int) bool {
		if basins[r][c] != 0 || g.depthAt(r, c) >= MAX_DEPTH {
			return false
		}
		basins[r][c] = currBasin
		for _, cell := range g.neighbors(r, c) {
			if g.depthAt(cell.row, cell.col) < MAX_DEPTH {
				flow(cell.row, cell.col)
			}
		}
		return true
	}

	// Recursively color-in the basins.
	for r := 0; r < g.nrows(); r++ {
		for c := 0; c < g.ncols(); c++ {
			if flow(r, c) {
				currBasin++
			}
		}
	}

	// Map basin id to number of cells in each basin.
	basinSize := make(map[int]int)
	for r := 0; r < g.nrows(); r++ {
		for c := 0; c < g.ncols(); c++ {
			if b := basins[r][c]; b != 0 {
				basinSize[b]++
			}
		}
	}

	// Compute the product of the sizes of the three largest basins.
	sizes := lib.MapValues(basinSize)
	sort.Ints(sizes)
	return int64(lib.Product(sizes[len(sizes)-3:]))
}

var DIRS = [4]cell{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}

type cell struct{ row, col int }
type grid struct{ depths [][]int8 }

func (g *grid) nrows() int { return len(g.depths) }
func (g *grid) ncols() int { return len(g.depths[0]) }

func (g *grid) inBounds(r, c int) bool {
	return r >= 0 && r < g.nrows() && c >= 0 && c < g.ncols()
}

func (g *grid) neighbors(r, c int) []cell {
	result := make([]cell, 0, len(DIRS))
	for _, dir := range DIRS {
		r2, c2 := r+dir.row, c+dir.col
		if g.inBounds(r2, c2) {
			result = append(result, cell{r2, c2})
		}
	}
	return result
}

func (g *grid) depthAt(r, c int) int8 {
	if g.inBounds(r, c) {
		return g.depths[r][c]
	}
	return MAX_DEPTH
}

func parseInput(lines []string) *grid {
	nrows := len(lines)
	ncols := len(lines[0])
	depths := lib.Make2DArray[int8](nrows, ncols)
	for r, line := range lines {
		for c, b := range line {
			depths[r][c] = int8(b - '0')
		}
	}
	return &grid{depths: depths}
}
