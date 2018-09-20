package dcp133

import (
	"math/rand"
	"testing"
)

// Test procedure: generate a tree at random, get all values in tree in sorted
// order, then verify we get the correct successor of each value. Repeat.
func TestFindInorderSuccessor(t *testing.T) {
	const (
		numTrials = 1000
		numNodes  = 100
	)
	for i := 0; i < numTrials; i++ {
		tree := randTree(numNodes)
		values := treeValues(tree)
		for i := 0; i < len(values)-1; i++ {
			succ := findInorderSuccessor(tree, values[i])
			if succ != values[i+1] {
				t.Fatalf("successor(%d) returned %d, want %d", values[i], succ, values[i+1])
			}
		}
		final := values[len(values)-1]
		if succ := findInorderSuccessor(tree, final); succ != 0 {
			t.Fatalf("successor(%d) returned %d, want 0", final, succ)
		}
	}
}

// Inserts the value into the tree.
func insert(root *node, value int) *node {
	if root == nil {
		return &node{value: value}
	}
	switch {
	case value < root.value:
		root.left = insert(root.left, value)
	case value > root.value:
		root.right = insert(root.right, value)
	}
	return root
}

// Generates a random binary tree with "size" nodes.
func randTree(size int) *node {
	var tree *node
	for i := 0; i < size; i++ {
		tree = insert(tree, rand.Intn(10*size))
	}
	return tree
}

// Returns the values in the tree in ascending order.
func treeValues(tree *node) []int {
	var values []int
	traverseInorder(tree, func(n *node) bool {
		values = append(values, n.value)
		return true
	})
	return values
}
