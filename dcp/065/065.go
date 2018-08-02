/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Amazon.

Given a N by M matrix of numbers, print out the matrix in a clockwise spiral.

For example, given the following matrix:

[[1,  2,  3,  4,  5],
 [6,  7,  8,  9,  10],
 [11, 12, 13, 14, 15],
 [16, 17, 18, 19, 20]]

You should print out the following:

1
2
3
4
5
10
15
20
19
18
17
16
11
6
7
8
9
14
13
12

*/
package dcp065

func spiral(mat [][]int) []int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return nil
	}
	var acc []int
	var rowfn, colfn func(row, col, nrows, ncols int, dir int)
	rowfn = func(row, col, nrows, ncols int, dir int) {
		if ncols == 0 {
			return
		}
		end := col + dir*ncols
		for i := col; i != end; i += dir {
			acc = append(acc, mat[row][i])
		}
		colfn(row+dir, end-dir, nrows-1, ncols, dir)
	}
	colfn = func(row, col, nrows, ncols int, dir int) {
		if nrows == 0 {
			return
		}
		end := row + dir*nrows
		for i := row; i != end; i += dir {
			acc = append(acc, mat[i][col])
		}
		rowfn(end-dir, col-dir, nrows, ncols-1, -dir)
	}
	rowfn(0, 0, len(mat), len(mat[0]), 1)
	return acc
}
