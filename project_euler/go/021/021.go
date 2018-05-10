/*

Let d(n) be defined as the sum of proper divisors of n (numbers less than n
which divide evenly into n).

If d(a) = b and d(b) = a, where a != b, then a and b are an amicable pair and
each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55
and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71
and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.

*/

package p021

import "math"

var memo = make(map[int64]int64)

func sumOfProperDivisors(n int64) int64 {
	if sum, ok := memo[n]; ok {
		return sum
	}
	var sum int64
	for _, d := range ProperDivisors(n) {
		sum += d
	}
	memo[n] = sum
	return sum
}

func solve() int64 {
	var sum, a, b int64
	for a = 1; a < 10000; a++ {
		for b = a + 1; b < 10000; b++ {
			if sumOfProperDivisors(a) == b && sumOfProperDivisors(b) == a {
				sum += a + b
			}
		}
	}
	return sum
}

func ProperDivisors(n int64) []int64 {
	var (
		sqrtn    = int64(math.Sqrt(float64(n)))
		divisors = []int64{1}
		m        int64
	)
	for m = 2; m <= sqrtn; m++ {
		if n%m == 0 {
			if m != n/m {
				divisors = append(divisors, m, n/m)
			} else {
				divisors = append(divisors, m)
			}
		}
	}
	return divisors
}
