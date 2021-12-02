package d04

import (
	"adventofcode2021/lib"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const N = 5

func Run() {
	lib.Run(4, part1, part2)
}

func part1() int64 {
	sequence, grids := parseInput(realInput)
	for _, x := range sequence {
		for _, g := range grids {
			if won, score := g.mark(x); won {
				return int64(score)
			}
		}
	}
	panic("no grid won")
}

func part2() int64 {
	var (
		sequence, grids = parseInput(realInput)
		lastTurn        int
		lastScore       int
	)
nextGrid:
	for _, g := range grids {
		for turn, x := range sequence {
			if won, score := g.mark(x); won {
				if turn > lastTurn {
					lastTurn, lastScore = turn, score
				}
				continue nextGrid
			}
		}
	}
	return int64(lastScore)
}

type Grid struct {
	values [N][N]int8
	marked [N][N]bool
}

func (g *Grid) mark(value int8) (bool, int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if g.values[i][j] == value && !g.marked[i][j] {
				g.marked[i][j] = true
				if g.isRowMarked(i) || g.isColMarked(j) {
					return true, g.sumUnmarked() * int(value)
				}
			}
		}
	}
	return false, 0
}

func (g *Grid) isRowMarked(r int) bool {
	for c := 0; c < N; c++ {
		if !g.marked[r][c] {
			return false
		}
	}
	return true
}

func (g *Grid) isColMarked(c int) bool {
	for r := 0; r < N; r++ {
		if !g.marked[r][c] {
			return false
		}
	}
	return true
}

func (g *Grid) sumUnmarked() int {
	var sum int
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if !g.marked[r][c] {
				sum += int(g.values[r][c])
			}
		}
	}
	return sum
}

func parseInput(lines []string) ([]int8, []*Grid) {
	sequence := parseCommaSeparatedInt8s(lines[0])
	lines = lines[1:]
	var grids []*Grid
	for len(lines) >= N+1 {
		grids = append(grids, parseGrid(lines[1:]))
		lines = lines[N+1:]
	}
	return sequence, grids
}

var whitespace = regexp.MustCompile(`\s+`)

func parseGrid(lines []string) *Grid {
	var g Grid
	for i := 0; i < N; i++ {
		row := whitespace.Split(strings.TrimSpace(lines[i]), -1)
		if len(row) != N {
			panic(fmt.Sprintf("parseGrid: line %d (%q) has %d values (%q), expected %d", i, lines[i], len(row), row, N))
		}
		for j, s := range row {
			g.values[i][j] = parseInt8(s)
		}
	}
	return &g
}

func parseInt8(s string) int8 {
	x, err := strconv.ParseInt(s, 10, 8)
	if err != nil || x < math.MinInt8 || x > math.MaxInt8 {
		panic(fmt.Sprintf("couldn't convert %q to int8: %v", s, x))
	}
	return int8(x)
}

func parseCommaSeparatedInt8s(s string) []int8 {
	parts := strings.Split(s, ",")
	result := make([]int8, len(parts))
	for i, t := range parts {
		result[i] = parseInt8(t)
	}
	return result
}
