/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Apple.

Given a tree, find the largest tree/subtree that is a BST.

Given a tree, return the size of the largest tree/subtree that is a BST.

*/
package dcp093

import (
	"math"
)

// Assumption: "largest" and "size" refer to the number of nodes in a
// tree/subtree, not the sum of its values--which could be of any type.

type Node struct {
	left  *Node
	right *Node
	value int64
}

func largestBST(root *Node) int {
	n, _ := largestBSTRec(root, math.MinInt64, math.MaxInt64)
	return n
}

func largestBSTRec(node *Node, minVal, maxVal int64) (int, bool) {
	if node == nil {
		return 0, true
	}
	validHere := true // can this node be part of parent BST?
	if node.value < minVal || node.value > maxVal {
		minVal, maxVal = math.MinInt64, math.MaxInt64
		validHere = false
	}
	nL, okL := largestBSTRec(node.left, minVal, node.value)
	nR, okR := largestBSTRec(node.right, node.value, maxVal)
	if okL && okR {
		return nL + nR + 1, validHere
	}
	return max(nL, nR), false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
