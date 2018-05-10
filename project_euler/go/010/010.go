/*

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

*/

package p010

const N = 2000000

func solve() int {
	sum := 0
	s := NewPrimeSieve()
	for p := s.Next(); p < N; p = s.Next() {
		sum += p
	}
	return sum
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
