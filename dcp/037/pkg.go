/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

The power set of a set is the set of all its subsets. Write a function that,
given a set, generates its power set.

For example, given the set {1, 2, 3}, it should return {{}, {1}, {2}, {3}, {1,
2}, {1, 3}, {2, 3}, {1, 2, 3}}.

You may also use a list or array to represent a set.

*/
package dcp037

func powerset(xs []int) [][]int {
	var acc [][]int
	var rec func(cur []int, n int)
	rec = func(cur []int, n int) {
		if n == len(xs) {
			acc = append(acc, cur)
			return
		}
		rec(append(cur, xs[n]), n+1)
		rec(cur, n+1)
	}
	rec([]int{}, 0)
	return acc
}
