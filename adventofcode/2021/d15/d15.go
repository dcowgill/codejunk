package d15

import (
	"adventofcode2021/lib"
	"container/heap"
)

func Run()         { lib.Run(15, part1, part2) }
func part1() int64 { return solve(realInput, 1) }
func part2() int64 { return solve(realInput, 5) }

func solve(input []string, expand int) int64 {
	g := expandGrid(parseInput(realInput), expand)
	src := point{0, 0}
	dst := point{g.nrows - 1, g.ncols - 1}
	return int64(bfs(g, src, dst))
}

func bfs(g *grid, src, dst point) int {
	var (
		visited   = make(map[point]bool)
		pq        = &priorityQueue{}
		neighbors = make([]point, len(DIRS))
	)
	heap.Push(pq, pqItem{src, 0})
	for pq.Len() > 0 {
		it := heap.Pop(pq).(pqItem)
		if it.p == dst {
			return it.cost
		}
		neighbors = g.neighbors(it.p, neighbors)
		for _, p := range neighbors {
			if !visited[p] {
				heap.Push(pq, pqItem{p, it.cost + g.cost(p)})
				visited[p] = true
			}
		}
	}
	panic("no path found")
}

type point struct {
	row, col int
}

type grid struct {
	costs [][]int
	nrows int
	ncols int
}

func (g *grid) cost(p point) int {
	return g.costs[p.row][p.col]
}

func (g *grid) inBounds(p point) bool {
	return p.row >= 0 && p.row < g.nrows && p.col >= 0 && p.col < g.ncols
}

var DIRS = [4]point{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}

func (g *grid) neighbors(p point, storage []point) []point {
	result := storage[:0]
	for _, dir := range DIRS {
		p2 := point{p.row + dir.row, p.col + dir.col}
		if g.inBounds(p2) {
			result = append(result, p2)
		}
	}
	return result
}

func parseInput(lines []string) *grid {
	var (
		nrows = len(lines)
		ncols = len(lines[0])
		costs = lib.Make2DArray[int](nrows, ncols)
	)
	for r, row := range lines {
		for c := 0; c < len(row); c++ {
			costs[r][c] = int(row[c] - '0')
		}
	}
	return &grid{costs: costs, nrows: nrows, ncols: ncols}
}

type pqItem struct {
	p    point
	cost int
}

type priorityQueue struct {
	heap []pqItem
}

func (pq priorityQueue) Len() int            { return len(pq.heap) }
func (pq priorityQueue) Less(i, j int) bool  { return pq.heap[i].cost < pq.heap[j].cost }
func (pq priorityQueue) Swap(i, j int)       { pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i] }
func (pq *priorityQueue) Push(x interface{}) { pq.heap = append(pq.heap, x.(pqItem)) }
func (pq *priorityQueue) Pop() interface{} {
	old := pq.heap
	n := len(old)
	x := old[n-1]
	pq.heap = old[0 : n-1]
	return x
}

func expandGrid(g *grid, n int) *grid {
	var (
		nrows2 = g.nrows * n
		ncols2 = g.ncols * n
		g2     = &grid{
			costs: lib.Make2DArray[int](nrows2, ncols2),
			nrows: nrows2,
			ncols: ncols2,
		}
	)
	for r := 0; r < g.nrows; r++ {
		for c := 0; c < g.ncols; c++ {
			g2.costs[r][c] = g.costs[r][c]
		}
	}
	dupe := func(r1, c1, r2, c2 int) {
		for sr, dr := r1, r2; sr < r1+g.nrows; sr, dr = sr+1, dr+1 {
			for sc, dc := c1, c2; sc < c1+g.ncols; sc, dc = sc+1, dc+1 {
				g2.costs[dr][dc] = (g2.costs[sr][sc] % 9) + 1
			}
		}
	}
	for i := 0; i < n-1; i++ {
		dupe(0, g.ncols*i, 0, g.ncols*(i+1))
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n; j++ {
			dupe(g.nrows*i, g.ncols*j, g.nrows*(i+1), g.ncols*j)
		}
	}
	return g2
}

// func printGrid(g *grid) {
// 	for r := 0; r < g.nrows; r++ {
// 		for c := 0; c < g.ncols; c++ {
// 			fmt.Printf("%d", g.costs[r][c])
// 		}
// 		fmt.Println("")
// 	}
// }
