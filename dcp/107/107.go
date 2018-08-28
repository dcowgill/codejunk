/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Microsoft.

Print the nodes in a binary tree level-wise. For example, the following should
print 1, 2, 3, 4, 5.

  1
 / \
2   3
   / \
  4   5

*/
package dcp107

import (
	"container/list"
	"fmt"
	"io"
)

type node struct {
	left  *node
	right *node
	value int
}

func traverseBinaryTreeBreadthFirst(root *node, fn func(*node)) {
	frontier := list.New() // deque
	frontier.PushBack(root)
	for frontier.Len() > 0 {
		elem := frontier.Front()
		frontier.Remove(elem)
		node := elem.Value.(*node)
		if node != nil {
			fn(node)
			frontier.PushBack(node.left)
			frontier.PushBack(node.right)
		}
	}
}

func printBinaryTreeBreadthFirst(w io.Writer, root *node) {
	sep := ""
	traverseBinaryTreeBreadthFirst(root, func(n *node) {
		fmt.Fprintf(w, "%s%d", sep, n.value)
		sep = ", "
	})
}
