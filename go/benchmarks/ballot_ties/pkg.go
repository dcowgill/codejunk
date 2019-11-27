package ballot_ties

import "github.com/dcowgill/i64"

type (
	empty  struct{}
	Rank   uint8
	Ballot []Rank
)

func NewBallot(n int) Ballot { return make([]Rank, n) }

func (b Ballot) HasTiesMap() bool {
	table := make(map[Rank]empty, 256)
	for _, r := range b {
		if r != 0 {
			if _, ok := table[r]; ok {
				return true
			}
			table[r] = empty{}
		}
	}
	return false
}

func (b Ballot) HasTiesBits() bool {
	var table [4]i64.Bits
	for _, r := range b {
		if r != 0 {
			i := r / 64
			if table[i].Test(int(r)) {
				return true
			}
			table[i].Set(int(r))
		}
	}
	return false
}

func (b Ballot) HasTiesBruteForce() bool {
	n := len(b)
	for i := 0; i < n-1; i++ {
		if r := b[i]; r != 0 {
			for j := i + 1; j < n; j++ {
				if r == b[j] {
					return true
				}
			}
		}
	}
	return false
}

func (b Ballot) HasTiesTable() bool {
	table := make([]bool, 256) // doesn't escape: gets allocated on stack
	for _, r := range b {
		if r != 0 {
			if table[r] {
				return true
			}
			table[r] = true
		}
	}
	return false
}
