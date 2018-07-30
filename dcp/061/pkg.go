/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Implement integer exponentiation. That is, implement the pow(x, y) function,
where x and y are integers and returns x^y.

Do this faster than the naive method of repeated multiplication.

For example, pow(2, 10) should return 1024.

*/
package dcp061

// Iterative version.
func fastpow(x, y int) int {
	if y == 0 {
		return 1
	}
	z := 1
	for y > 1 {
		if y%2 == 0 {
			x *= x
			y /= 2
		} else {
			z *= x
			x *= x
			y = (y - 1) / 2
		}
	}
	return x * z
}

// Non tail-recursive version.
func fastpowRec(x, y int) int {
	switch {
	case y == 0:
		return 1
	case y == 1:
		return x
	case y%2 == 0:
		return fastpow(x*x, y/2)
	default:
		return x * fastpow(x*x, (y-1)/2)
	}
}
