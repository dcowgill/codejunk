/*
Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

The area of a circle is defined as πr^2. Estimate π to 3 decimal places using a
Monte Carlo method.

Hint: The basic equation of a circle is x^2 + y^2 = r^2.

*/
package dcp014

import (
	"math"
	"math/rand"
)

// Area of unit square = 1
// Area of unit circle = π*r^2 = π/4
// area(circle) : area(square) = hits : total
// π*r^2/1 = hits/total
// π = hits/total/r^2 = 4*hits/total
func estimatePi(trials int) float64 {
	in := 0
	for i := 0; i < trials; i++ {
		x := rand.Float64()
		y := rand.Float64()
		r := math.Sqrt(x*x + y*y)
		if r <= 1 {
			in++
		}
	}
	return 4 * float64(in) / float64(trials)
}
