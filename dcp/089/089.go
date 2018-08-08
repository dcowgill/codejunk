/*

Good morning! Here's your coding interview problem for today.

This problem was asked by LinkedIn.

Determine whether a tree is a valid binary search tree.

A binary search tree is a tree with two children, left and right, and satisfies
the constraint that the key in the left child must be less than or equal to the
root and the key in the right child must be greater than or equal to the root.

*/
package dcp089

import "math"

type Key int64

type Node struct {
	left  *Node
	right *Node
	key   Key
}

func isValid(tree *Node) bool {
	return isValidSubtree(tree, math.MinInt64, math.MaxInt64)
}

func isValidSubtree(tree *Node, minKey, maxKey Key) bool {
	switch {
	case tree == nil:
		return true
	case tree.key < minKey || tree.key > maxKey:
		return false
	case !isValidSubtree(tree.left, minKey, tree.key):
		return false
	case !isValidSubtree(tree.right, tree.key, maxKey):
		return false
	}
	return true
}
