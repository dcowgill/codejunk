/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Two Sigma.

Using a function rand7() that returns an integer from 1 to 7 (inclusive) with
uniform probability, implement a function rand5() that returns an integer from 1
to 5 (inclusive).

*/
package dcp071

import (
	"math/rand"
)

// Provided.
func rand7() int { return rand.Intn(7) + 1 }

// Returns [1,5] and rerolls [6,7].
func rand5_reroll() int {
	for {
		if x := rand7(); x <= 5 {
			return x
		}
	}
}

// Generates a random base-7 int in [0,6666] and returns it modulo 5 + 1.
func rand5_modulo() int {
	r := func() int { return rand7() - 1 } // generates ints in [0, 7)
	n := 2401*r() + 343*r() + 49*r() + 7*r() + r()
	return n%5 + 1
}
