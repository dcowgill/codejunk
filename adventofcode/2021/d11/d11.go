package d11

import (
	"adventofcode2021/lib"
)

func Run() {
	lib.Run(11, part1, part2)
}

func part1() int64 {
	g := makeGrid(realInput)
	nflashes := 0
	for i := 0; i < 100; i++ {
		nflashes += step(g)
	}
	return int64(nflashes)
}

func part2() int64 {
	g := makeGrid(realInput)
	nsteps := 0
	for {
		nsteps++
		if step(g) == g.nrows()*g.ncols() {
			return int64(nsteps)
		}
	}
	panic("unreachable")
}

const MAX = 9

func step(g *grid) int {
	var flash, bump func(pos)
	bump = func(p pos) {
		n := g.levels[p.row][p.col]
		if n == MAX {
			g.levels[p.row][p.col] = MAX + 1
			flash(p)
		} else if n < MAX {
			g.levels[p.row][p.col]++
		}
	}
	flash = func(p pos) {
		for _, dir := range DIRS {
			neighbor := p.add(dir)
			if g.inbounds(neighbor) {
				bump(neighbor)
			}
		}
	}
	for r := 0; r < g.nrows(); r++ {
		for c := 0; c < g.ncols(); c++ {
			bump(pos{r, c})
		}
	}
	nflashes := 0
	for r := 0; r < g.nrows(); r++ {
		for c := 0; c < g.ncols(); c++ {
			if g.levels[r][c] > MAX {
				g.levels[r][c] = 0
				nflashes++
			}
		}
	}
	return nflashes
}

var DIRS = []pos{{-1, -1}, {-1, 0}, {-1, +1}, {0, -1}, {0, +1}, {+1, -1}, {+1, 0}, {+1, +1}}

type pos struct{ row, col int }

func (p pos) add(q pos) pos { return pos{p.row + q.row, p.col + q.col} }

type grid struct{ levels [][]int }

func (g *grid) nrows() int { return len(g.levels) }
func (g *grid) ncols() int { return len(g.levels[0]) }
func (g *grid) inbounds(p pos) bool {
	return p.row >= 0 && p.row < g.nrows() && p.col >= 0 && p.col < g.ncols()
}

func makeGrid(lines []string) *grid {
	levels := lib.Make2DArray[int](len(lines), len(lines[0]))
	for r, line := range lines {
		for c := 0; c < len(line); c++ {
			levels[r][c] = int(line[c] - '0')
		}
	}
	return &grid{levels: levels}
}
