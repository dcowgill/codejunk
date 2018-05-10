/*

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
that the 6th prime is 13.

What is the 10001st prime number?

*/
package p007

const N = 10001

func solve() int {
	s := NewPrimeSieve()
	for i := 0; i < N-1; i++ {
		s.Next()
	}
	return s.Next()
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
