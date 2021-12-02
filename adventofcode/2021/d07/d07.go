package d07

import (
	"adventofcode2021/lib"
	"math"
)

func Run()         { lib.Run(7, part1, part2) }
func part1() int64 { return solve(realInput, cost1) }
func part2() int64 { return solve(realInput, cost2) }

func solve(crabs []int, costFn func(x, y int) int) int64 {
	var (
		maxX    = lib.Greatest(crabs)
		minCost = math.MaxInt
	)
	for x := 0; x <= maxX; x++ {
		cost := 0
		for _, c := range crabs {
			cost += costFn(x, c)
		}
		if cost < minCost {
			minCost = cost
		}
	}
	return int64(minCost)
}

func cost1(x, y int) int {
	return lib.Abs(x - y)
}

func cost2(x, y int) int {
	n := lib.Abs(x - y)
	return n * (n + 1) / 2
}
