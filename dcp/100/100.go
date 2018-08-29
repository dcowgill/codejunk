/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

You are in an infinite 2D grid where you can move in any of the 8 directions:

  (x,y) to
    (x+1, y),
    (x-1, y),
    (x, y+1),
    (x, y-1),
    (x-1, y-1),
    (x+1, y+1),
    (x-1, y+1),
    (x+1, y-1)

You are given a sequence of points and the order in which you need to cover the
points. Give the minimum number of steps in which you can achieve it. You start
from the first point.

Example:

Input: [(0, 0), (1, 1), (1, 2)]
Output: 2

It takes 1 step to move from (0, 0) to (1, 1). It takes one more step to move
from (1, 1) to (1, 2).

*/
package dcp100

// A point in the 2D grid.
type pt struct {
	x, y int
}

// Computes the minimum number of steps required to complete the path.
func pathSteps(path []pt) int {
	var d int
	for i := 1; i < len(path); i++ {
		d += distance(path[i-1], path[i])
	}
	return d
}

// The distance between two points in the grid is the sum of (1) the lesser of
// their horizontal and vertical deltas, representing the diagonal movement, and
// (2) the difference of those deltas, representing the non-diagonal movement.
func distance(src, dst pt) int {
	xd := abs(dst.x - src.x)
	yd := abs(dst.y - src.y)
	return min(xd, yd) + abs(xd-yd)
}

// Helpers:

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
