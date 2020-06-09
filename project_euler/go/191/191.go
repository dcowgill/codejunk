package main

import "fmt"

// Only six kinds of sequence need to be tracked.
// L = number of 'L' chars in string
// S = number of consecutive 'A' chars in suffix
type state struct {
	NS0 int // L = 0, S = 0
	NS1 int // L = 0, S = 1
	NS2 int // L = 0, S = 2
	LS0 int // L = 1, S = 0
	LS1 int // L = 1, S = 1
	LS2 int // L = 1, S = 2
}

// Computes the state for the next day.
func (st state) next() state {
	var x state

	// Append an 'O'
	x.NS0 += st.NS0 + st.NS1 + st.NS2
	x.LS0 += st.LS0 + st.LS1 + st.LS2

	// Append an 'A'
	x.NS1 += st.NS0
	x.NS2 += st.NS1
	x.LS1 += st.LS0
	x.LS2 += st.LS1

	// Append an 'L'
	x.LS0 += st.NS0 + st.NS1 + st.NS2

	return x
}

// Counts valid sequences.
func (st state) total() int {
	return st.NS0 + st.NS1 + st.NS2 + st.LS0 + st.LS1 + st.LS2
}

func main() {
	st := state{NS0: 1, NS1: 1, LS0: 1} // state after day 1
	for i := 1; i <= 30; i++ {
		fmt.Printf("day %2d: %+v (%d)\n", i, st, st.total())
		st = st.next()
	}
}
