package main

import (
	"fmt"
	"math/rand"
	"time"
)

type normalDist struct {
	mean, stdDev float64
}

func (d normalDist) sample() float64 {
	return rand.NormFloat64()*d.stdDev + d.mean
}

func sim(d1, d2 normalDist, pop int, p1, p2 float64, cutoff float64) (int, int) {
	pop1 := int(float64(pop) * p1)
	pop2 := int(float64(pop) * p2)
	var n1, n2 int
	for i := 0; i < pop1; i++ {
		if s := d1.sample(); s >= cutoff {
			n1++
		}
	}
	for i := 0; i < pop2; i++ {
		if s := d2.sample(); s >= cutoff {
			n2++
		}
	}
	return n1, n2
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	const (
		pop    = 1000 * 1000
		p1     = 0.76
		p2     = 0.13
		cutoff = 1170
	)
	d1 := normalDist{1170, 179}
	d2 := normalDist{995, 167}
	n1, n2 := sim(d1, d2, pop, p1, p2, cutoff)
	fmt.Printf("n1 = %d (%d%%), n2 = %d (%d%%)\n", n1, 100*n1/pop, n2, 100*n2/pop)
}
