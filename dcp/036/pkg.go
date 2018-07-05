/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Dropbox.

Given the root to a binary search tree, find the second largest node in the tree.

*/
package dcp036

import (
	"math"
)

// A node in a binary search tree.
type Node struct {
	left  *Node // subtree whose values are less than this node
	right *Node // subtree whose values are greater than this node
	value int64
}

// Reports the second-largest value in the tree.
// If the tree contains fewer than two values, returns math.MinInt64.
func secondLargest(tree *Node) int64 {
	if tree == nil {
		return math.MinInt64
	}
	_, x := topTwo(tree)
	return x
}

// Reports the largest and second-largest values in the tree.
func topTwo(tree *Node) (int64, int64) {
	if tree.right != nil { // don't need to explore left subtree
		a, b := topTwo(tree.right)
		return max(a, b), max(min(a, b), tree.value)
	}
	if tree.left != nil {
		a, b := topTwo(tree.left)
		return tree.value, max(a, b)
	}
	return tree.value, math.MinInt64
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
