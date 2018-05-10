/*

The following iterative sequence is defined for the set of positive integers:

n -> n/2 (n is even)
n -> 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13  40  20  10  5  16  8  4  2  1

It can be seen that this sequence (starting at 13 and finishing at 1) contains
10 terms. Although it has not been proved yet (Collatz Problem), it is thought
that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.

*/

package p014

func next(n int64) int64 {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func solve() int {
	var best, bestLen int
	for n := 500001; n < 1000000; n += 2 {
		l := 1
		var i int64
		for i = int64(n); i != 1; i = next(i) {
			l++
		}
		if l > bestLen {
			bestLen = l
			best = n
		}
	}
	return best
}
