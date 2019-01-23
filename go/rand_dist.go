package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var mean float64
	var ntrials int
	flag.Float64Var(&mean, "mean", 1.0, "mean")
	flag.IntVar(&ntrials, "n", 1, "number of trials")
	flag.Parse()

	fmt.Println("====================")
	fmt.Println("Poisson")
	fmt.Println("====================")
	sim(mean, ntrials, poissonGen, poissonPDF)

	fmt.Println("====================")
	fmt.Println("Exponential")
	fmt.Println("====================")
	sim(1/mean, ntrials, expGen, expPDF)
}

func sim(λ float64, ntrials int, gen func(float64) int, pdf func(float64, int) float64) {
	hits := make(map[int]int)
	for i := 0; i < ntrials; i++ {
		hits[gen(λ)]++
	}
	min, max := 9999, 0
	for k := range hits {
		if k < min {
			min = k
		}
		if k > max {
			max = k
		}
	}
	total := 0
	for i := min; i <= max; i++ {
		v := hits[i]
		if v == 0 {
			continue
		}
		total += i * v
		pct := 100 * float64(v) / float64(ntrials)
		fmt.Printf("%3d = %5d (%5.2f%%) (%5.2f%%)\n", i, v, pct, 100*pdf(λ, i))
	}
	fmt.Printf("mean = %.2f\n", float64(total)/float64(ntrials))
}

//====================
// exponential
//====================

func expPDF(λ float64, x int) float64 {
	if x < 0 {
		return 0
	}
	return λ * math.Exp(-λ*float64(x))
}

func expGen(λ float64) int {
	u := rand.Float64()
	return int(math.Round(math.Log(1-u) / (-λ)))
}

//====================
// poisson
//====================

func poissonPDF(λ float64, x int) float64 {
	return math.Exp(-λ) * math.Pow(λ, float64(x)) / factorial(x)
}

func factorial(x int) float64 {
	n := 1.0
	for i := 2; i <= x; i++ {
		n *= float64(i)
	}
	return n
}

// Knuth.
func poissonGen(λ float64) int {
	l := math.Exp(-λ)
	k := 0
	p := 1.0
	for p > l {
		k += 1
		u := rand.Float64()
		p *= u
	}
	return k - 1
}
