/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Microsoft.

Let X be a set of n intervals on the real line. We say that a set of points P
"stabs" X if every interval in X contains at least one point in P. Compute the
smallest set of points that stabs X.

For example, given the intervals [(1, 4), (4, 5), (7, 9), (9, 12)], you should
return [4, 9].

*/
package dcp200

import "sort"

type interval struct {
	l, h int
}

func minStabSet(a []interval) []int {
	if len(a) == 0 {
		return nil
	}
	sort.Slice(a, func(i, j int) bool { return a[i].h < a[j].h })
	var stab []int
	endpoint := a[0].h
	for i := 1; i < len(a); i++ {
		if a[i].l > endpoint {
			stab = append(stab, endpoint)
			endpoint = a[i].h
		}
	}
	stab = append(stab, endpoint)
	return stab
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
