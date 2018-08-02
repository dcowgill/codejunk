/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Given a stream of elements too large to store in memory, pick a random element
from the stream with uniform probability.

*/
package dcp015

import "math/rand"

// choose selects a value from the channel, with uniform probability.
// Technique borrowed from The Practice of Programming, Chapter 3.
func choose(ch <-chan int) int {
	var v, n int
	for s := range ch {
		n++
		if rand.Int()%n == 0 {
			v = s
		}
	}
	return v
}
