/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Given the root of a binary tree, return a deepest node. For example, in the
following tree, return d.

    a
   / \
  b   c
 /
d

*/
package dcp080

// A node in a binary tree.
type node struct {
	left, right *node
	value       string
}

// Returns the value of the deepest node in the tree. If multiple nodes have
// equal depth, the function arbitrarily chooses one of them to return.
func findDeepestNode(root *node) string {
	var dfs func(cur *node, depth int) (*node, int)
	dfs = func(cur *node, depth int) (*node, int) {
		if cur == nil {
			return nil, 0
		}
		deepest, maxDepth := cur, depth
		if ln, ld := dfs(cur.left, depth+1); ld > maxDepth {
			deepest, maxDepth = ln, ld
		}
		if rn, rd := dfs(cur.right, depth+1); rd > maxDepth {
			deepest, maxDepth = rn, rd
		}
		return deepest, maxDepth
	}
	if deepest, _ := dfs(root, 1); deepest != nil {
		return deepest.value
	}
	return ""
}
