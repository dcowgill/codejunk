/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Invert a binary tree.

For example, given the following tree:

    a
   / \
  b   c
 / \  /
d   e f
should become:

  a
 / \
 c  b
 \  / \
  f e  d

*/
package dcp083

type node struct {
	value string
	left  *node
	right *node
}

func invertTree(tree *node) *node {
	if tree != nil {
		tree.left, tree.right = invertTree(tree.right), invertTree(tree.left)
	}
	return tree
}
