/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Given a function that generates perfectly random numbers between 1 and k
(inclusive), where k is an input, write a function that shuffles a deck of cards
represented as an array using only swaps.

It should run in O(N) time.

Hint: Make sure each one of the 52! permutations of the deck is equally likely.

*/
package dcp051

import (
	"math/rand"
)

// Shuffles cards in place.
// Algorithm P, AOCP Volume 1.
func shuffle(cards []int) {
	for i := len(cards) - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
}
