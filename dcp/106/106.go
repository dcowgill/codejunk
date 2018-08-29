/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Pinterest.

Given an integer list where each number represents the number of hops you can
make, determine whether you can reach to the last index starting at index 0.

For example, [2, 0, 1, 0] returns true while [1, 1, 0, 1] returns false.

*/
package dcp106

func canReachEnd(hops []int) bool {
	// Once a position has been tried, we do not need to try again. Otherwise,
	// the previous call would have returned true, terminating the search.
	memo := make(map[int]struct{})
	var try func(pos int) bool
	try = func(pos int) bool {
		// Check the cache; have we tried pos before?
		if _, ok := memo[pos]; ok {
			return false
		}
		memo[pos] = struct{}{} // remember

		// If we jumped past the end, signal failure.
		// If we have reached the final position, signal success.
		switch len(hops) - pos {
		case 0:
			return false
		case 1:
			return true
		}

		// Try all possible number of hops from this position.
		for i := hops[pos]; i >= 1; i-- {
			if pos+i < len(hops) && try(pos+i) {
				return true
			}
		}

		// It is impossible to reach the end from this position.
		return false
	}

	// Begin the search.
	return try(0)
}
