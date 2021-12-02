package d05

import (
	"adventofcode2021/lib"
	"fmt"
)

func Run()         { lib.Run(5, part1, part2) }
func part1() int64 { return solve(parseInput(realInput), false) }
func part2() int64 { return solve(parseInput(realInput), true) }

func solve(lines []Line, countDiagonals bool) int64 {
	grid := make(map[Point]int)
	for _, l := range lines {
		if l.p1.y == l.p2.y { // horizontal
			p1, p2 := l.p1, l.p2
			if p1.x > p2.x {
				p1, p2 = p2, p1
			}
			for x := p1.x; x <= p2.x; x++ {
				grid[Point{x, l.p1.y}]++
			}
		} else if l.p1.x == l.p2.x { // vertical
			p1, p2 := l.p1, l.p2
			if p1.y > p2.y {
				p1, p2 = p2, p1
			}
			for y := p1.y; y <= p2.y; y++ {
				grid[Point{l.p1.x, y}]++
			}
		} else if countDiagonals {
			xd := direction(l.p1.x, l.p2.x)
			yd := direction(l.p1.y, l.p2.y)
			for x, y := l.p1.x, l.p1.y; x != l.p2.x+xd && y != l.p2.y+yd; x, y = x+xd, y+yd {
				grid[Point{x, y}]++
			}
		}
	}
	answer := 0
	for _, v := range grid {
		if v >= 2 {
			answer++
		}
	}
	return int64(answer)
}

func direction(a, b int64) int64 {
	if a < b {
		return 1
	}
	return -1
}

type Point struct{ x, y int64 }
type Line struct{ p1, p2 Point }

func parseInput(input []string) []Line {
	lines := make([]Line, len(input))
	for i, s := range input {
		l := &lines[i]
		if _, err := fmt.Sscanf(s, "%d,%d -> %d,%d", &l.p1.x, &l.p1.y, &l.p2.x, &l.p2.y); err != nil {
			panic(fmt.Sprintf("failed to parse line %d: %q", i, s))
		}
	}
	return lines
}
