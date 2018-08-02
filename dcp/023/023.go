/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

You are given an M by N matrix consisting of booleans that represents a board.
Each True boolean represents a wall. Each False boolean represents a tile you
can walk on.

Given this matrix, a start coordinate, and an end coordinate, return the minimum
number of steps required to reach the end coordinate from the start. If there is
no possible path, then return null. You can move up, left, down, and right. You
cannot move through walls. You cannot wrap around the edges of the board.

For example, given the following board:

[[f, f, f, f],
[t, t, f, t],
[f, f, f, f],
[f, f, f, f]]

and start = (3, 0) (bottom left) and end = (0, 0) (top left), the minimum number
of steps required to reach the end is 7, since we would need to go through (1,
2) because there is a wall everywhere else on the second row.

*/
package dcp023

import (
	"container/heap"
)

type point struct {
	x, y int
}

type step struct {
	point point
	steps int
}

// frontier is a priority queue of steps, with the top element being the least
// distant from the starting point, in terms of number of steps.
type frontier []step

func (h frontier) Len() int            { return len(h) }
func (h frontier) Less(i, j int) bool  { return h[i].steps < h[j].steps }
func (h frontier) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *frontier) Push(x interface{}) { *h = append(*h, x.(step)) }
func (h *frontier) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// bfs does a breadth-first search the space described by the adjacency matrix.
// Returns the minimum number of steps from begin to end.
func bfs(adj [][]bool, begin, end point) int {
	fr := new(frontier)
	visited := make(map[point]bool)
	heap.Push(fr, step{begin, 0})
	for fr.Len() != 0 {
		next := heap.Pop(fr).(step)
		if next.point == end {
			return next.steps
		}
		for _, p := range neighbors(adj, next.point) {
			if !visited[p] {
				heap.Push(fr, step{p, next.steps + 1})
				visited[p] = true
			}
		}
	}
	return -1
}

// neighbors returns the points reachable in one step from p.
func neighbors(adj [][]bool, p point) []point {
	qs := []point{
		{p.x - 1, p.y}, // left
		{p.x + 1, p.y}, // right
		{p.x, p.y - 1}, // up
		{p.x, p.y + 1}, // down
	}
	result := make([]point, 0, 4)
	for _, q := range qs {
		if q.x >= 0 && q.x < len(adj) && q.y >= 0 && q.y < len(adj[0]) {
			if !adj[q.x][q.y] { // not a wall
				result = append(result, q)
			}
		}
	}
	return result
}
