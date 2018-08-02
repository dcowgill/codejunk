/*


Good morning! Here's your coding interview problem for today.

This problem was asked by Amazon.

Given a matrix of 1s and 0s, return the number of "islands" in the matrix. A 1
represents land and 0 represents water, so an island is a group of 1s that are
neighboring and their perimiter is surrounded by water.

For example, this matrix has 4 islands.

1 0 0 0 0
0 0 1 1 0
0 1 1 0 0
0 0 0 0 0
1 1 0 0 1
1 1 0 0 1

*/
package dcp084

func countIslands(mat [][]int) int {
	type pos struct {
		row, col int
	}

	// isLand reports whether p (1) is in bounds and (2) represents a land tile.
	isLand := func(p pos) bool {
		return p.row >= 0 && p.row < len(mat) && p.col >= 0 && p.col < len(mat[0]) && mat[p.row][p.col] != 0
	}

	// neighbors returns the land tiles that abut p.
	neighbors := func(p pos) []pos {
		answer := make([]pos, 0, 4)
		for _, dir := range []pos{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
			q := pos{row: p.row + dir.row, col: p.col + dir.col}
			if isLand(q) {
				answer = append(answer, q)
			}
		}
		return answer
	}

	// visit marks p, its neighbors, its neighbors' neighbors, etc. as visited.
	visited := make(map[pos]bool)
	var visit func(p pos)
	visit = func(p pos) {
		visited[p] = true
		for _, q := range neighbors(p) {
			if !visited[q] {
				visit(q)
			}
		}
	}

	// Count the islands.
	n := 0
	for row := range mat {
		for col := range mat[row] {
			p := pos{row, col}
			if isLand(p) && !visited[p] {
				visit(p)
				n++
			}
		}
	}
	return n
}
