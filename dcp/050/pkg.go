/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Microsoft.

Suppose an arithmetic expression is given as a binary tree. Each leaf is an
integer and each internal node is one of '+', '−', '∗', or '/'.

Given the root to such a tree, write a function to evaluate it.

For example, given the following tree:

    *
   / \
  +    +
 / \  / \
3  2  4  5

You should return 45, as it is (3 + 2) * (4 + 5).

*/
package dcp050

import "fmt"

type expr interface {
	eval() int
}

type op int

const (
	opAdd op = iota + 1
	opSub
	opMul
	opDiv
)

type binaryExpr struct {
	lhs expr
	rhs expr
	op  op
}

func (e binaryExpr) eval() int {
	lhs := e.lhs.eval()
	rhs := e.rhs.eval()
	switch e.op {
	case opAdd:
		return lhs + rhs
	case opSub:
		return lhs - rhs
	case opMul:
		return lhs * rhs
	case opDiv:
		return lhs / rhs
	}
	panic(fmt.Sprintf("invalid binary operator: %d", e.op))
}

type constExpr int

func (e constExpr) eval() int {
	return int(e)
}
