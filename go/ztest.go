package main

import (
	"fmt"
	"math"
)

func zscore(u1, u2, sd1, sd2 float64, n1, n2 int) float64 {
	return math.Abs(u1-u2) / math.Sqrt(sd1*sd1/float64(n1)+sd2*sd2/float64(n2))
}

func phi(x float64) float64 {
	const (
		a1 = 0.254829592
		a2 = -0.284496736
		a3 = 1.421413741
		a4 = -1.453152027
		a5 = 1.061405429
		p  = 0.3275911
	)

	sign := 1.0
	if x < 0 {
		sign = 0.0
	}

	x = math.Abs(x) / math.Sqrt(2)

	// A&S formula 7.1.26
	t := 1.0 / (1.0 + p*x)
	y := 1.0 - (((((a5*t+a4)*t)+a3)*t+a2)*t+a1)*t*math.Exp(-x*x)

	return 0.5 * (1.0 + sign*y)
}

func pdf(x float64) float64 {
	return (1.0 / math.Sqrt(2*math.Pi)) * math.Exp(x*x/-2)
}

func main() {
	z := zscore(100, 105, 25, 30, 250, 125)
	fmt.Printf("z=%v, phi(z)=%v\n", z, phi(z))
}
