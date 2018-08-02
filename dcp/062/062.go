/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

There is an N by M matrix of zeroes. Given N and M, write a function to count
the number of ways of starting at the top-left corner and getting to the
bottom-right corner. You can only move right or down.

For example, given a 2 by 2 matrix, you should return 2, since there are two
ways to get to the bottom-right:

Right, then down
Down, then right

Given a 5 by 5 matrix, there are 70 ways to get to the bottom-right.

*/
package dcp062

func dfs(n, m, i, j int) int {
	if i == n-1 && j == m-1 {
		return 1
	}
	x := 0
	if i < n-1 {
		x += dfs(n, m, i+1, j) // right
	}
	if j < m-1 {
		x += dfs(n, m, i, j+1) // down
	}
	return x
}
