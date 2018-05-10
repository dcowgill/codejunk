/*

215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 21000?

*/

package p016

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

const N = 1000

func solve() int {
	var (
		b = big.NewInt(2)
		e = big.NewInt(1000)
		n = new(big.Int)
	)
	return sumDigits(n.Exp(b, e, nil))
}
