/*

Starting in the top left corner of a 22 grid, there are 6 routes (without
backtracking) to the bottom right corner.


How many routes are there through a 2020 grid?

*/

package p015

import (
	"math/big"
)

const N = 20

func solve() string {
	return new(big.Int).Binomial(2*N, N).String()
}
