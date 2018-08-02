/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

A knight's tour is a sequence of moves by a knight on a chessboard such that all
squares are visited once.

Given N, write a function to return the number of knight's tours on an N by N
chessboard.

*/
package main

import "fmt"

// Note: I haven't solved this one. A single tour can be found in linear time
// (see Warnsdorf's rule) but I don't know how to count the number of tours
// without brute force, i.e. search with backtracking.

type square struct{ row, col int }

func (s square) add(t square) square {
	return square{s.row + t.row, s.col + t.col}
}

func valid(s square, n int) bool {
	return s.row >= 0 && s.row < n && s.col >= 0 && s.col < n
}

var moves = []square{{-2, -1}, {-2, +1}, {-1, +2}, {+1, +2}, {+2, +1}, {+2, -1}, {+1, -2}, {-1, -2}}

// Try all possible tours. Impractical for N > 5.
func bruteForce(N int) int {
	visited := newBoolMatrix(N, N)
	numTours := 0
	var dfs func(src square, numVisited int) int
	dfs = func(src square, numVisited int) int {
		if numVisited == N*N-1 {
			return 1
		}
		visited[src.row][src.col] = true
		n := 0
		for _, m := range moves {
			dst := src.add(m)
			if valid(dst, N) && !visited[dst.row][dst.col] {
				n += dfs(dst, numVisited+1)
			}
		}
		visited[src.row][src.col] = false
		return n
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			numTours += dfs(square{i, j}, 1)
		}
	}
	return numTours
}

// Returns an r x c matrix of bools.
func newBoolMatrix(r, c int) [][]bool {
	mem := make([]bool, r*c)
	mat := make([][]bool, r)
	for i := 0; i < r; i++ {
		mat[i], mem = mem[:c], mem[c:]
	}
	return mat
}

func main() {
	fmt.Println(bruteForce(5))
}
