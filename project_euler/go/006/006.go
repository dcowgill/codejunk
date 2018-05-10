/*

The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten
natural numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first one
hundred natural numbers and the square of the sum.

*/
package p006

var N = 100

func solve() int {
	// Sum of the first N naturals is n(n+1)/2
	a := N * (N + 1) / 2

	// Sum of the squares of the first N naturals is n(n+1)(2n+1)/6
	b := N * (N + 1) * (2*N + 1) / 6

	return a*a - b
}
