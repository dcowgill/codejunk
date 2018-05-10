/*

n! means n  (n  1)  ...  3  2  1

For example, 10! = 10  9  ...  3  2  1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!

*/

package p020

import (
	"math/big"
	"strconv"
)

func sumDigits(n *big.Int) int {
	s := n.String()
	t := 0
	for i, _ := range s {
		x, _ := strconv.Atoi(s[i : i+1])
		t += x
	}
	return t
}

const N = 100

func solve() int {
	return sumDigits(Factorial(N))
}

var ONE = big.NewInt(1)

func Factorial(n int) *big.Int {
	var (
		f = big.NewInt(1)
		i = big.NewInt(int64(n))
	)
	for ; i.Cmp(ONE) > 0; i.Sub(i, ONE) {
		f.Mul(f, i)
	}
	return f
}
