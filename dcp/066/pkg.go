/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Square.

Assume you have access to a function toss_biased() which returns 0 or 1 with a
probability that's not 50-50 (but also not 0-100 or 100-0). You do not know the
bias of the coin.

Write a function to simulate an unbiased coin toss.

*/
package dcp066

import (
	"math/rand"
)

// Returns a function that returns true with probability p.
func makeBiasedCoin(p float64) func() bool {
	return func() bool {
		return rand.Float64() < p
	}
}

// Converts a biased coin to a fair coin.
func makeFairCoin(coin func() bool) func() bool {
	return func() bool {
		for {
			// Accept only HT or TH, which occur with equal probability.
			switch a, b := coin(), coin(); {
			case a && !b:
				return true
			case b && !a:
				return false
			}
		}
	}
}
