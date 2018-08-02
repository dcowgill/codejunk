/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

You are given an array of non-negative integers that represents a
two-dimensional elevation map where each element is unit-width wall and the
integer is the height. Suppose it will rain and all spots between two walls get
filled up.

Compute how many units of water remain trapped on the map in O(N) time and O(1) space.

For example, given the input [2, 1, 2], we can hold 1 unit of water in the middle.

Given the input [3, 0, 1, 3, 0, 5], we can hold 3 units in the first index, 2 in
the second, and 3 in the fourth index (we cannot hold 5 since it would run off
to the left), so we can trap 8 units of water.

*/
package dcp030

func rainfall(wall []int) int {
	if len(wall) == 0 {
		return 0
	}
	var (
		i = 0             // current position from the left
		j = len(wall) - 1 // current position from the right
		l = wall[i]       // highest wall to the left of i
		r = wall[j]       // highest wall to the right of j
		n = 0             // rainfall accumulator
	)
	inc := func(x int) {
		if h := min(l, r); h > wall[x] {
			n += h - wall[x]
		}
	}
	for i < j {
		if wall[i] < wall[j] {
			i++
			l = max(l, wall[i])
			inc(i)
		} else {
			j--
			r = max(r, wall[j])
			inc(j)
		}
	}
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
