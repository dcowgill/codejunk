package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	part := flag.Int("part", 1, "which part")
	flag.Parse()
	switch *part {
	case 1:
		part1()
	case 2:
		part2()
	}
}

const (
	serial = 7400
	size   = 300
)

func part1() {
	sums := summedAreaTable(initPowerLevels(serial))
	bestI, bestJ, best := maxSquareSum(sums, 3)
	fmt.Printf("%d,%d = %d\n", bestI+1, bestJ+1, best)
}

func part2() {
	sums := summedAreaTable(initPowerLevels(serial))
	var bestI, bestJ, bestSize int
	best := math.MinInt64
	for n := 1; n <= size; n++ {
		if i, j, sum := maxSquareSum(sums, n); sum > best {
			bestI, bestJ, bestSize, best = i, j, n, sum
		}
	}
	fmt.Printf("%d,%d,%d = %d\n", bestI+1, bestJ+1, bestSize, best)
}

func initPowerLevels(serialNum int) [][]int {
	levels := make([][]int, size)
	for i := range levels {
		levels[i] = make([]int, size)
		for j := range levels[i] {
			levels[i][j] = powerLevel(i+1, j+1, serialNum)
		}
	}
	return levels
}

func powerLevel(x, y, serialNum int) int {
	rackID := x + 10
	return (((rackID*y+serialNum)*rackID)/100)%10 - 5
}

// Returns the summed-area table of the power levels.
func summedAreaTable(levels [][]int) [][]int {
	sums := make([][]int, len(levels))
	for i := range levels {
		sums[i] = make([]int, len(levels[i]))
		copy(sums[i], levels[i])
	}
	for i := 1; i < size; i++ {
		sums[i][0] += sums[i-1][0]
	}
	for j := 1; j < size; j++ {
		sums[0][j] += sums[0][j-1]
	}
	for i := 1; i < size; i++ {
		for j := 1; j < size; j++ {
			sums[i][j] += sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1]
		}
	}
	return sums
}

// Reports the sum of the power levels in the size n square at (i, j).
func squareSum(sums [][]int, i, j, n int) int {
	if i == 0 || j == 0 {
		return sums[i+n-1][j+n-1]
	}
	x, y := i+n-1, j+n-1 // lower right corner of square
	return sums[x][y] - sums[x-n][y] - sums[x][y-n] + sums[x-n][y-n]
}

// Finds the size n square with the greatest sum of power levels.
func maxSquareSum(sums [][]int, n int) (bestI, bestJ, best int) {
	best = math.MinInt64
	for i := 0; i <= size-n; i++ {
		for j := 0; j <= size-n; j++ {
			if sum := squareSum(sums, i, j, n); sum > best {
				bestI, bestJ, best = i, j, sum
			}
		}
	}
	return bestI, bestJ, best
}
