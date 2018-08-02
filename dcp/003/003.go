/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Given the root to a binary tree, implement serialize(root), which serializes the
tree into a string, and deserialize(s), which deserializes the string back into
the tree.

For example, given the following Node class

class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
The following test should pass:

node = Node('root', Node('left', Node('left.left')), Node('right'))
assert deserialize(serialize(node)).left.left.val == 'left.left'

*/
package dcp003

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Node struct {
	Left  *Node
	Right *Node

	// Because Python is dynamically typed, it's unclear
	// from the problem description whether the node values
	// are always strings. Since the problem is much harder
	// otherwise, let's assume they have type "string".
	Val string
}

func serialize(root *Node) string {
	var sb strings.Builder
	var rec func(*Node)
	rec = func(n *Node) {
		if n == nil {
			sb.WriteByte('_')
			return
		}
		sb.WriteString(fmt.Sprintf("%q", n.Val))
		rec(n.Left)
		rec(n.Right)
	}
	rec(root)
	return sb.String()
}

func deserialize(s string) *Node {
	buf := bytes.NewBufferString(s)
	var rec func() *Node
	rec = func() *Node {
		var n *Node
		switch b, err := buf.ReadByte(); {
		case err != nil:
			panic(err)
		case b == '_':
			return nil
		}
		buf.UnreadByte()
		n = &Node{}
		_, err := fmt.Fscanf(buf, "%q", &n.Val)
		if err != nil {
			panic(err)
		}
		n.Left = rec()
		n.Right = rec()
		return n
	}
	return rec()
}

// JSON-based reference implementation, for testing:

func serializeJSON(n *Node) string {
	data, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func deserializeJSON(input string) *Node {
	var n *Node
	if err := json.Unmarshal([]byte(input), &n); err != nil {
		panic(err)
	}
	return n
}
