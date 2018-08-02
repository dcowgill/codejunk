/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Microsoft.

You have an N by N board. Write a function that, given N, returns the number of
possible arrangements of the board where N queens can be placed on the board
without threatening each other, i.e. no two queens share the same row, column,
or diagonal.

*/
package dcp038

// Solves the N-queens problem for n.
func queens(n int) int {
	return search(newBoard(n), 0)
}

func search(b [][]bool, col int) int {
	if col == len(b) {
		return 1
	}
	n := 0
	for row := 0; row < len(b); row++ {
		if legal(b, row, col) {
			b[row][col] = true
			n += search(b, col+1)
			b[row][col] = false
		}
	}
	return n
}

// Reports whether a queen can legally be placed at (row, col). Assumes there is
// one queen in every column to the left of col and zero queens to the right.
func legal(b [][]bool, row, col int) bool {
	// Check the current row.
	for c := 0; c <= col; c++ {
		if b[row][c] {
			return false
		}
	}
	// Check the diagonal along increasing rows.
	for r, c := row+1, col-1; r < len(b) && c >= 0; r, c = r+1, c-1 {
		if b[r][c] {
			return false
		}
	}
	// Check the diagonal along decreasing rows.
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if b[r][c] {
			return false
		}
	}
	return true // safe!
}

// Allocates an NxN matrix of bools.
func newBoard(n int) [][]bool {
	mem := make([]bool, n*n)
	mat := make([][]bool, n)
	for i := 0; i < n; i++ {
		mat[i], mem = mem[:n], mem[n:]
	}
	return mat
}
