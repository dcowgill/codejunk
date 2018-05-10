/*

A Pythagorean triplet is a set of three natural numbers, a b c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000. Find
the product abc.

*/

package p009

const N = 1000

func solve() int {
	for a := 1; a < N/3; a++ {
		for b := a + 1; b < N/2; b++ {
			for c := b + 1; c < N; c++ {
				if a*a+b*b == c*c {
					// If (a,b,c) is a pythagorean triple, then (ka,kb,kc) is
					// also a triple for all k>2. Therefore, if ka+kb+kc=N for
					// some integer k, the answer is ka*kb*kc.
					s := a + b + c
					if N%s == 0 {
						k := N / s
						return k * k * k * a * b * c
					}
				}
			}
		}
	}
	panic("not found")
}
