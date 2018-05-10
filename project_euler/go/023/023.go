/*

A perfect number is a number for which the sum of its proper divisors is
exactly equal to the number. For example, the sum of the proper divisors of 28
would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than
n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest
number that can be written as the sum of two abundant numbers is 24. By
mathematical analysis, it can be shown that all integers greater than 28123
can be written as the sum of two abundant numbers. However, this upper limit
cannot be reduced any further by analysis even though it is known that the
greatest number that cannot be expressed as the sum of two abundant numbers is
less than this limit.

Find the sum of all the positive integers which cannot be written as the sum
of two abundant numbers.

*/

package p023

import "math"

func isAbundant(n int) bool {
	var sum int64
	for _, d := range ProperDivisors(int64(n)) {
		sum += d
	}
	return sum > int64(n)
}

const LIMIT = 28123

func solve() int {
	var abundants []int
	for i := 1; i <= LIMIT; i++ {
		if isAbundant(i) {
			abundants = append(abundants, i)
		}
	}

	sumOfTwoAbundants := make([]bool, LIMIT+1)
	for _, a := range abundants {
		for _, b := range abundants {
			if a+b <= LIMIT {
				sumOfTwoAbundants[a+b] = true
			}
		}
	}

	sum := 0
	for i := 1; i < len(sumOfTwoAbundants); i++ {
		if !sumOfTwoAbundants[i] {
			sum += i
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
