package main

import (
	"fmt"
	"sort"
)

var rankToInt = map[byte]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10, '9': 9,
	'8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
}

type Reverse struct{ sort.Interface }

func (r Reverse) Less(i, j int) bool { return r.Interface.Less(j, i) }

type Card struct {
	Rank int
	Suit int
}

func (c Card) String() string {
	return fmt.Sprintf("Card(r=%d,s=%d)", c.Rank, c.Suit)
}

type Hand struct {
	cards [5]Card
	ranks [5]int
	suits [5]int
	freqs map[int]int
}

func NewHand(cards ...Card) *Hand {
	if len(cards) != 5 {
		return nil
	}
	h := Hand{freqs: make(map[int]int, 14)}
	// Copy stuff into hand struct.
	copy(h.cards[:], cards)
	for i, c := range cards {
		h.ranks[i] = c.Rank
		h.suits[i] = c.Suit
	}
	// Sort ranks and correct for ace-low straight.
	sort.Sort(Reverse{sort.IntSlice(h.ranks[:])})
	if h.ranks == [...]int{14, 5, 4, 3, 2} {
		h.ranks = [...]int{5, 4, 3, 2, 1}
	}
	// Compute rank frequencies.
	for _, r := range h.ranks {
		h.freqs[r]++
	}
	return &h
}

func (h *Hand) Straight() bool {
	for i := 1; i < 5; i++ {
		if h.ranks[i] != h.ranks[i-1]-1 {
			return false
		}
	}
	return true
}

func (h *Hand) Flush() bool {
	s := h.suits[0]
	for i := 1; i < 5; i++ {
		if h.suits[i] != s {
			return false
		}
	}
	return true
}

func (h *Hand) Kind(n int) {

}

func main() {
	sf := NewHand(Card{4, 0}, Card{2, 0}, Card{5, 0}, Card{14, 0}, Card{3, 0})
	fh := NewHand(Card{4, 0}, Card{4, 1}, Card{4, 2}, Card{11, 0}, Card{11, 1})
	fmt.Printf("%v straight=%v flush=%v\n", sf, sf.Straight(), sf.Flush())
	fmt.Printf("%v straight=%v flush=%v\n", fh, fh.Straight(), fh.Flush())
}
