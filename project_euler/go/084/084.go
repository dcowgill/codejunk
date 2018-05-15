// https://projecteuler.net/problem=84
package p084

import (
	"fmt"
	"math/rand"
	"sort"
)

// All the spaces on the monopoly board, in clockwise order.
var spaces = []string{
	"GO", "A1", "CC1", "A2", "T1", "R1", "B1", "CH1", "B2", "B3", "JAIL",
	"C1", "U1", "C2", "C3", "R2", "D1", "CC2", "D2", "D3", "FP",
	"E1", "CH2", "E2", "E3", "R3", "F1", "F2", "U2", "F3", "G2J",
	"G1", "G2", "CC3", "G3", "R4", "CH3", "H1", "T2", "H2",
}

// Reports the index of the named space.
func find(name string) int {
	for i, s := range spaces {
		if name == s {
			return i
		}
	}
	panic(fmt.Sprintf("not a valid space: %q", name))
}

// Notable spaces.
var (
	GO   = find("GO")
	JAIL = find("JAIL")
	G2J  = find("G2J")
	CH1  = find("CH1")
	CH2  = find("CH2")
	CH3  = find("CH3")
	CC1  = find("CC1")
	CC2  = find("CC2")
	CC3  = find("CC3")
	C1   = find("C1")
	E3   = find("E3")
	H2   = find("H2")
	R1   = find("R1")
)

// Reports whether s is a Chance space.
func chance(s int) bool { return s == CH1 || s == CH2 || s == CH3 }

// Reports whether s is a Community Chest space.
func communityChest(s int) bool { return s == CC1 || s == CC2 || s == CC3 }

// Reports whether s is a railroad.
func railroad(s int) bool { return spaces[s][0] == 'R' }

// Reports whether s is a utility.
func utility(s int) bool { return spaces[s][0] == 'U' }

// Locates the next space satisfying the predicate.
func next(pred func(space int) bool, from int) int {
	s := from
	for {
		if pred(s) {
			return s
		}
		s = (s + 1) % len(spaces)
	}
}

// Returns a random permutation of the first n integers.
func permutation(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	for i := len(a) - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return a
}

// A card is a command to move from one space to another.
type Card func(from int) (to int)

// A deck of cards.
type Deck struct {
	cards []Card
	pos   int
}

// Draws the top card, follows its instructions--reporting the space on
// which the player should be sent--then puts the card on the bottom.
func (d *Deck) draw(from int) int {
	top := d.cards[d.pos]
	d.pos = (d.pos + 1) % len(d.cards)
	if top != nil {
		return top(from)
	}
	return from // "blank" card
}

// Creates a new deck using the given cards, padded to 16 blank cards.
func newDeck(cards ...Card) *Deck {
	a := make([]Card, 16)
	b := make([]Card, len(a))
	copy(a, cards)
	for i, j := range permutation(len(a)) {
		b[i] = a[j]
	}
	return &Deck{cards: b}
}

// Parameters to the simulation.
type simulationParams struct {
	numTurns int
	numDice  int
	dieSize  int
}

// NumTurns reports the number of simulation turns. Uses a suitable default if
// p.numTurns is not a positive value.
func (p simulationParams) NumTurns() int {
	if p.numTurns <= 0 {
		return 1000000
	}
	return p.numTurns
}

// Rolls an n-sided die and reports the result.
func die(n int) int { return rand.Intn(n) + 1 }

// Rolls the dice and returns their sum. The second value will be true
// if and only if all the dice came up the same.
func (p simulationParams) roll() (int, bool) {
	first := die(p.dieSize)
	sum := first
	same := true
	for i := 1; i < p.numDice; i++ {
		r := die(p.dieSize)
		same = same && r == first
		sum += r
	}
	return sum, same
}

// Pairs a space and a frequency count.
type SpaceCount struct{ Space, Count int }

// Sorts by Count, low to high.
type ByCount []SpaceCount

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count < a[j].Count }

// simulate runs the Monopoly game simulation and returns the frequency
// counts the spaces on which each turn ended, ordered by space.
func simulate(params simulationParams) []SpaceCount {
	var visit func(int) int

	// Create the Chance deck.
	cc := newDeck(
		func(from int) int { return GO },
		func(from int) int { return JAIL },
	)

	// Create the Community Chest deck.
	ch := newDeck(
		func(from int) int { return GO },
		func(from int) int { return JAIL },
		func(from int) int { return C1 },
		func(from int) int { return E3 },
		func(from int) int { return H2 },
		func(from int) int { return R1 },
		func(from int) int { return next(railroad, from) },
		func(from int) int { return next(railroad, from) },
		func(from int) int { return next(utility, from) },
		func(from int) int { return visit((from - 3 + len(spaces)) % len(spaces)) },
	)

	// Create the visitation function. Only a few spaces are special.
	var curr int
	visit = func(space int) int {
		switch {
		case space == G2J:
			return JAIL
		case communityChest(space):
			return cc.draw(space)
		case chance(space):
			return ch.draw(space)
		}
		return space
	}

	// Initialize frequency counters.
	freq := make([]SpaceCount, len(spaces))
	for i := range freq {
		freq[i].Space = i
	}

	// Run the simulation.
	doubles := 0
	for i := 0; i < params.NumTurns(); i++ {
		sum, same := params.roll()
		if same {
			doubles++
		} else {
			doubles = 0
		}
		if doubles >= 3 {
			doubles = 0
			curr = JAIL
		} else {
			curr = visit((curr + sum) % len(spaces))
		}
		freq[curr].Count++
	}
	return freq
}

// Normalizes counts to the range [0,1].
func normalize(freq []SpaceCount) []float64 {
	sum := 0
	for _, sc := range freq {
		sum += sc.Count
	}
	a := make([]float64, len(freq))
	for i, sc := range freq {
		a[i] = float64(sc.Count) / float64(sum)
	}
	return a
}

func solve(params simulationParams, topN int) string {
	// Run the simulation.
	freq := simulate(params)

	// Sort by frequency.
	sort.Sort(sort.Reverse(ByCount(freq)))

	// Format the answer key.
	answer := ""
	for i := 0; i < topN; i++ {
		answer += fmt.Sprintf("%02d", freq[i].Space)
	}
	return answer
}
