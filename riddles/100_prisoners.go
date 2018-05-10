// https://en.wikipedia.org/wiki/100_prisoners_problem
package main

import (
	"flag"
	"fmt"
	"math/rand"
)

var (
	numPrisoners int
	numTries     int
)

func main() {
	flag.IntVar(&numPrisoners, "prisoners", 100, "number of prisoners")
	flag.IntVar(&numTries, "tries", 50, "number of boxes each prisoner may open")
	numTrials := flag.Int("trials", 1, "number of simulations")
	seed := flag.Int64("seed", 0, "PRNG seed")
	flag.Parse()

	// Seed the PRNG (optional).
	if *seed != 0 {
		rand.Seed(*seed)
	}

	// Run the simulations.
	numSuccesses := 0
	for i := 0; i < *numTrials; i++ {
		if simulate() {
			numSuccesses++
		}
	}

	// Display the results.
	fmt.Printf("trials = %d, successes = %d, success rate = %0.2f%%\n",
		*numTrials, numSuccesses, 100*float64(numSuccesses)/float64(*numTrials))

}

// simulate runs a simulation of the 100 prisoners problem.
// Reports whether the prisoners were successful.
func simulate() bool {
	boxes := rand.Perm(numPrisoners)
	for i := 0; i < numPrisoners; i++ {
		if !search(boxes, i) {
			return false
		}
	}
	return true
}

// search simulates one prisoner's attempt to find his own number in the boxes,
// using the strategy described here:
// https://en.wikipedia.org/wiki/100_prisoners_problem#Strategy
func search(boxes []int, prisoner int) bool {
	box := prisoner
	for i := 0; i < numTries; i++ {
		box = boxes[box]
		if box == prisoner {
			return true
		}
	}
	return false
}
