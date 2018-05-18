/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

A unival tree (which stands for "universal value") is a tree where all nodes
under it have the same value.

Given the root to a binary tree, count the number of unival subtrees.

For example, the following tree has 5 unival subtrees:

   0
  / \
 1   0
    / \
   1   0
  / \
 1   1

*/
package dcp008

type node struct {
	left  *node
	right *node
	value int
}

// univals returns two values: the number of unival subtrees of n (including n
// itself, if it is one), and whether n is itself a unival tree.
func univals(n *node) (int, bool) {
	if n == nil {
		return 0, true
	}

	ln, lu := univals(n.left)
	rn, ru := univals(n.right)
	c := ln + rn // number of unival trees beneath n
	u := false   // true if n is a unival tree

	// n is a unival tree if both of its subtrees
	// (1) are univals and
	// (2) have the same value as n, or are nil.
	if (lu && (n.left == nil || n.left.value == n.value)) && (ru && (n.right == nil || n.right.value == n.value)) {
		c = c + 1
		u = true
	}

	return c, u
}
