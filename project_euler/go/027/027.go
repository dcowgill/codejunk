/*

Euler published the remarkable quadratic formula:

n² + n + 41

It turns out that the formula will produce 40 primes for the consecutive
values n = 0 to 39. However, when n = 40, 40² + 40 + 41 = 40(40 + 1) +
41 is divisible by 41, and certainly when n = 41, 41² + 41 + 41 is
clearly divisible by 41.

Using computers, the incredible formula n² - 79n + 1601 was discovered,
which produces 80 primes for the consecutive values n = 0 to 79. The
product of the coefficients, -79 and 1601, is -126479.

Considering quadratics of the form:

n² + an + b, where |a| < 1000 and |b| < 1000

where |n| is the modulus/absolute value of n
e.g. |11| = 11 and |4| = 4

Find the product of the coefficients, a and b, for the quadratic
expression that produces the maximum number of primes for consecutive
values of n, starting with n = 0.

*/

package p027

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func countPrimes(s *PrimeSieve, a, b int) int {
	c := 0
	for n := 0; s.IsPrime(abs(n*n + a*n + b)); n++ {
		c++
	}
	return c
}

func solve() int {
	var most, bestA, bestB int
	s := NewPrimeSieve()
	for b := -1000; b <= 1000; b++ {
		if !s.IsPrime(abs(b)) {
			continue // n=0 gives b; therefore |b| must be prime
		}
		for a := -1000; a <= 1000; a++ {
			n := countPrimes(s, a, b)
			if n >= most {
				most, bestA, bestB = n, a, b
			}
		}
	}
	return bestA * bestB
}

type PrimeSieve struct {
	x          int
	composites map[int][]int
	primes     []int
	n          int
}

func NewPrimeSieve() *PrimeSieve {
	s := new(PrimeSieve)
	s.x = 1
	s.composites = make(map[int][]int)
	s.primes = make([]int, 0)
	s.n = 0
	return s
}

func (s *PrimeSieve) Next() int {
	if s.n >= len(s.primes) {
		for {
			s.x++
			if primes, ok := s.composites[s.x]; ok {
				for _, prime := range primes {
					k := s.x + prime
					if xs, ok := s.composites[k]; ok {
						s.composites[k] = append(xs, prime)
					} else {
						s.composites[k] = []int{prime}
					}
				}
				delete(s.composites, s.x)
			} else {
				s.composites[s.x*s.x] = []int{s.x}
				break
			}
		}
		s.primes = append(s.primes, s.x)
	}
	s.n++
	return s.primes[s.n-1]
}

func (s *PrimeSieve) Reset() {
	s.n = 0
}

func (s *PrimeSieve) IsPrime(x int) bool {
	s.Reset()
	p := s.Next()
	for p < x {
		p = s.Next()
	}
	return p == x
}
