/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

You are given an N by M 2D matrix of lowercase letters. Determine the minimum
number of columns that can be removed to ensure that each row is ordered from
top to bottom lexicographically. That is, the letter at each column is
lexicographically later as you go down each row. It does not matter whether each
row itself is ordered lexicographically.

For example, given the following table:

cba
daf
ghi

This is not ordered because of the a in the center. We can remove the second
column to make it ordered:

ca
df
gi

So your function should return 1, since we only needed to remove 1 column.

As another example, given the following table:

abcdef

Your function should return 0, since the rows are already ordered (there's only
one row).

As another example, given the following table:

zyx
wvu
tsr

Your function should return 3, since we would need to remove all the columns to
order it.

*/
package dcp076

// TODO: is there clever approach here? This just does the obvious thing.
func minColsToRemove(mat [][]rune) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}
	numRows := len(mat)
	numCols := len(mat[0])
	isSorted := func(col int) bool {
		for row := 1; row < numRows; row++ {
			if mat[row][col] < mat[row-1][col] {
				return false
			}
		}
		return true
	}
	n := 0 // number of columns to remove
	for col := 0; col < numCols; col++ {
		if !isSorted(col) {
			n++
		}
	}
	return n
}
