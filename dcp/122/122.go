/*

Good morning! Here's your coding interview problem for today.

This question was asked by Zillow.

You are given a 2-d matrix where each cell represents number of coins in that
cell. Assuming we start at matrix[0][0], and can only move right or down, find
the maximum number of coins you can collect by the bottom right corner.

For example, in this matrix

0 3 1 1
2 0 0 4
1 5 3 1

The most we can collect is 0 + 2 + 1 + 5 + 3 + 1 = 12 coins.

*/
package dcp122

type point struct {
	row, col int
}

// Reports the maximum number of coins that can be collected from the grid,
// starting at the top left, or point (0, 0).
func maxCoins(grid [][]int) int {
	cache := make(map[point]int)
	var fn func(pos point) int
	fn = func(pos point) int {
		// Don't solve the same problem more than once.
		if amount, ok := cache[pos]; ok {
			return amount
		}
		// If we find ourselves off the grid, stop exploring.
		if pos.row < 0 || pos.row >= len(grid) {
			return 0
		}
		if pos.col < 0 || pos.col >= len(grid[pos.row]) {
			return 0
		}
		// Try each choice, down and right, and add the greater of the two to the
		// number of coins in the current grid square. Memoize the result.
		x := fn(point{pos.row + 1, pos.col})
		y := fn(point{pos.row, pos.col + 1})
		amount := grid[pos.row][pos.col] + max(x, y)
		cache[pos] = amount
		return amount
	}
	return fn(point{0, 0})
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
