/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Microsoft.

Given a 2D matrix of characters and a target word, write a function that returns
whether the word can be found in the matrix by going left-to-right, or
up-to-down.

For example, given the following matrix:

[['F', 'A', 'C', 'I'],
 ['O', 'B', 'Q', 'P'],
 ['A', 'N', 'O', 'B'],
 ['M', 'A', 'S', 'S']]

and the target word 'FOAM', you should return true, since it's the leftmost
column. Similarly, given the target word 'MASS', you should return true, since
it's the last row.

*/
package dcp063

// Rabin-Karp string search (using a very bad rolling hash function). See:
// https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm
//
// Knuth–Morris–Pratt or Boyer-Moore are more efficient, but Rabin-Karp is
// straightforward to implement, especially with this simple hash function, and
// requires little pre-processing.
func indexRabinKarp(s, substr string) int {
	n := len(substr)
	if len(s) < n {
		return -1
	}
	var hashss, h uint32
	for i := 0; i < n; i++ {
		hashss += uint32(substr[i])
		h += uint32(s[i])
	}
	for i := n; i < len(s); i++ {
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
		h += uint32(s[i])
		h -= uint32(s[i-n])
	}
	return -1
}

// Unrolls the rune matrix into a single string.
// Uses newlines to separate words.
// Consumes twice as much memory as the matrix.
func mat2str(mat [][]rune) string {
	var b []rune // TODO: size is known; could alloc once
	numRows := len(mat)
	numCols := len(mat[0])
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			b = append(b, mat[i][j])
		}
		b = append(b, '\n')
	}
	for j := 0; j < numCols; j++ {
		for i := 0; i < numRows; i++ {
			b = append(b, mat[i][j])
		}
		b = append(b, '\n')
	}
	return string(b)
}
