/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Amazon.

Given a node in a binary tree, return the next bigger element, also known as the inorder successor.

For example, the inorder successor of 22 is 30.

   10
  /  \
 5    30
     /  \
   22    35

You can assume each node has a parent pointer.

*/
package dcp133

type node struct {
	left  *node
	right *node
	value int
}

// Returns the inorder successor of the specified value in the tree. If the
// value has no successor, returns zero.
func findInorderSuccessor(root *node, value int) int {
	var answer int
	traverseInorder(root, func(n *node) bool {
		if n.value > value {
			answer = n.value
			return false
		}
		return true
	})
	return answer
}

// Performs an inorder traversal of the tree, calling fn once for each node. If
// fn returns false, the traversal is immediately halted.
func traverseInorder(tree *node, fn func(*node) bool) bool {
	if tree == nil {
		return true
	}
	return traverseInorder(tree.left, fn) && fn(tree) && traverseInorder(tree.right, fn)
}
