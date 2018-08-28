/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Given a binary tree, return all paths from the root to leaves.

For example, given the tree

   1
  / \
 2   3
    / \
   4   5

it should return [[1, 2], [1, 3, 4], [1, 3, 5]].

*/
package dcp110

type node struct {
	left  *node
	right *node
	value int
}

func allPathsToLeaves(tree *node) [][]int {
	// impl recursively finds all paths to the leaves of n, but in reverse
	// (because it's inefficient to prepend elements to a slice).
	var impl func(n *node) [][]int
	impl = func(n *node) [][]int {
		if n == nil {
			return nil
		}
		if n.left == nil && n.right == nil {
			return [][]int{{n.value}} // leaf
		}
		var result [][]int
		for _, path := range impl(n.left) {
			result = append(result, append(path, n.value))
		}
		for _, path := range impl(n.right) {
			result = append(result, append(path, n.value))
		}
		return result
	}
	// Reverse the paths returned by impl before returning.
	answer := impl(tree)
	for _, path := range answer {
		reverse(path)
	}
	return answer
}

// Reverses a slice in-place.
func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
