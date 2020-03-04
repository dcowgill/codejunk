package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type histogram map[int]int

func chiSquared(expected, observed histogram) float64 {
	var sum float64
	for x, exp := range expected {
		obs := observed[x]
		e, o := float64(exp), float64(obs)
		sum += math.Pow(o-e, 2) / e
	}
	return sum
}

func simulate(nrolls int) histogram {
	h := make(histogram)
	for i := 0; i < nrolls; i++ {
		h[d6()]++
	}
	return h
}

func d6() int {
	return rand.Intn(6) + 1
}

func makeExpected(nrolls int) histogram {
	h := make(histogram)
	for i := 1; i <= 6; i++ {
		h[i] = nrolls / 6
	}
	return h
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	const nrolls = 60
	expected := makeExpected(nrolls)

	// sample := histogram{1: 8, 2: 9, 3: 19, 4: 6, 5: 8, 6: 10}
	sample := histogram{1: 8, 2: 12, 3: 8, 4: 12, 5: 8, 6: 12}
	x0 := chiSquared(expected, sample)

	const ntrials = 10000
	n := 0
	var max float64
	var maxObs histogram
	for i := 0; i < ntrials; i++ {
		hist := simulate(nrolls)
		x1 := chiSquared(expected, hist)
		if x1 >= x0 {
			n++
		}
		if x1 > max {
			max, maxObs = x1, hist
		}
	}

	fmt.Printf("x0 = %.3f\n", x0)
	fmt.Printf("%.2f%% (%d/%d) of trials had X^2 >= x0\n", 100*float64(n)/ntrials, n, ntrials)
	fmt.Printf("max X^2 = %.3f\n", max)
	fmt.Printf("max_hist = %+v\n", maxObs)
}
