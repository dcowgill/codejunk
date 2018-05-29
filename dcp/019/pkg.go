/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

A builder is looking to build a row of N houses that can be of K different
colors. He has a goal of minimizing cost while ensuring that no two neighboring
houses are of the same color.

Given an N by K matrix where the nth row and kth column represents the cost to
build the nth house with kth color, return the minimum cost which achieves this
goal.

*/
package dcp019

// minCost reports the minimum cost of painting a row of N houses, such that
// consecutive houses do not share a color, where N is len(costs) and
// costs[n][k] is the cost of painting the nth house in color k. Assumes the
// cost matrix has size N x K, where K is the number of available colors.
func mincost(costs [][]int) int {
	// Infer the number of available colors from the cost matrix.
	numColors := len(costs[0])

	// The problem has shared substructures, so memoize calls to f.
	type key struct{ house, k int }
	cache := make(map[key]int)

	// f reports the minimum cost of painting the remaining houses, given that h
	// houses have already been painted, and the previous house used color pk.
	var f func(h, pk int) int
	f = func(h, pk int) int {
		if h == len(costs) {
			return 0 // all houses have been painted
		}
		if min, ok := cache[key{h, pk}]; ok {
			return min // already computed this state
		}
		min := 1<<63 - 1
		for k := 0; k < numColors; k++ {
			if k != pk {
				x := costs[h][k] + f(h+1, k)
				if x < min {
					min = x
				}
			}
		}
		cache[key{h, pk}] = min
		return min
	}

	// Begin with zero houses painted and no previous color.
	return f(0, -1)
}
