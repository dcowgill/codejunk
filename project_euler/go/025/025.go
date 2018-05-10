/*

The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn1 + Fn2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144

The 12th term, F12, is the first term to contain three digits.

What is the first term in the Fibonacci sequence to contain 1000 digits?

*/

package p025

import "math/big"

// Calls f for successive values in the Fibonacci sequence until f
// returns true, then returns the index of the value in the sequence.
func fibs(f func(n *big.Int) bool) int {
	var (
		a = big.NewInt(1)
		b = big.NewInt(1)
		t = new(big.Int)
		i = 1
	)
	for {
		if f(a) {
			return i
		}
		t.Set(b)
		b.Add(a, b)
		a.Set(t)
		i++
	}
}

const NDIGITS = 1000

func solve() int {
	minval := new(big.Int)
	minval.Exp(big.NewInt(10), big.NewInt(NDIGITS-1), nil)
	return fibs(func(n *big.Int) bool { return n.Cmp(minval) >= 0 })
}
