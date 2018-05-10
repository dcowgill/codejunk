package main

import (
	"fmt"
	"math/big"
)

// Computes n choose k.
func binomialCoefficient(n, k int) *big.Int {
	p := big.NewRat(1, 1)
	for i := 1; i <= k; i++ {
		p.Mul(p, big.NewRat(int64(n+1-i), int64(i)))
	}
	c := p.Num()
	return c.Div(c, p.Denom())
}

func main() {
	bells := []*big.Int{big.NewInt(1)}
	for n := 0; ; n++ {
		fmt.Printf("%s\n", bells[len(bells)-1])
		b := big.NewInt(0)
		for k := 0; k <= n; k++ {
			x := binomialCoefficient(n, k)
			b.Add(b, x.Mul(x, bells[k]))
		}
		bells = append(bells, b)
	}
}
