/*

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/
package p003

func solve() int {
	var n int64 = 600851475143
	factors := PrimeFactors(n, NewPrimeSieve())
	return factors[len(factors)-1].Base
}

type primeFactor struct {
	Base     int
	Exponent int
}

func PrimeFactors(n int64, s *PrimeSieve) []primeFactor {
	factors := make([]primeFactor, 0)
	s.Reset()
	for n > 1 {
		p := int64(s.Next())
		if n%p == 0 {
			exp := 0
			for n%p == 0 {
				n /= p
				exp++
			}
			factors = append(factors, primeFactor{int(p), exp})
		}
	}
	return factors
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
