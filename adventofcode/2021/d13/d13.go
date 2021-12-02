package d13

import (
	"adventofcode2021/lib"
	"fmt"
)

func Run() {
	lib.Run(13, part1, part2)
}

func part1() int64 {
	prob := realInput
	return int64(len(executeFold(prob.dots, prob.folds[0])))
}

/*

This is the day with an answer that is not an integer. Printing the final grid
shows the following diagram:

.##....##.####..##..#....#..#.###....##
#..#....#....#.#..#.#....#..#.#..#....#
#.......#...#..#....#....#..#.#..#....#
#.##....#..#...#.##.#....#..#.###.....#
#..#.#..#.#....#..#.#....#..#.#....#..#
.###..##..####..###.####..##..#.....##.

I.e. "GJZGLUPJ".

Need to fix the "run" framework to accept answers of any type, not just int64.

*/

func part2() int64 {
	prob := realInput
	dots := prob.dots
	for _, f := range prob.folds {
		dots = executeFold(dots, f)
	}
	return 0 // FIXME
}

func executeFold(dots []point, f fold) []point {
	dotmap := make(map[point]bool)
	if f.axis == "x" {
		xmax := f.n * 2
		for _, d := range dots {
			if d.x == f.n {
				panic("found dot on x-axis of fold")
			} else if d.x > f.n {
				d.x = xmax - d.x
			}
			dotmap[d] = true
		}
	} else if f.axis == "y" {
		ymax := f.n * 2
		for _, d := range dots {
			if d.y == f.n {
				panic("found dot on y-axis of fold")
			} else if d.y > f.n {
				d.y = ymax - d.y
			}
			dotmap[d] = true
		}
	} else {
		panic("illegal axis")
	}
	newdots := make([]point, 0, len(dotmap))
	for p := range dotmap {
		newdots = append(newdots, p)
	}
	return newdots
}

func printDots(dots []point) {
	m := make(map[point]bool, len(dots))
	for _, d := range dots {
		m[d] = true
	}
	max := maxPoint(dots)
	for y := 0; y <= max.y; y++ {
		for x := 0; x <= max.x; x++ {
			ch := "."
			if m[point{x, y}] {
				ch = "#"
			}
			fmt.Print(ch)
		}
		fmt.Println("")
	}
}

func maxPoint(ps []point) point {
	var max point
	for _, p := range ps {
		if p.x > max.x {
			max.x = p.x
		}
		if p.y > max.y {
			max.y = p.y
		}
	}
	return max
}
