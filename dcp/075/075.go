/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Microsoft.

Given an array of numbers, find the length of the longest increasing subsequence
in the array. The subsequence does not necessarily have to be contiguous.

For example, given the array [0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15],
the longest increasing subsequence has length 6: it is 0, 2, 6, 9, 11, 15.

*/
package dcp075

// Returns a longest increasing subsequence of a.
func longestIncreasingSubseq(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	piles := sortIntoPiles(a)
	var seq []int
	for p := piles[len(piles)-1].bottom(); p != nil; p = p.next {
		seq = append(seq, p.value)
	}
	for i, j := 0, len(seq)-1; i < j; i, j = i+1, j-1 {
		seq[i], seq[j] = seq[j], seq[i]
	}
	return seq
}

type Card struct {
	value int
	next  *Card
}

type Pile struct {
	cards []*Card
}

func newPile(c *Card) *Pile   { return &Pile{cards: []*Card{c}} }
func (p *Pile) add(c *Card)   { p.cards = append(p.cards, c) }
func (p *Pile) top() *Card    { return p.cards[len(p.cards)-1] }
func (p *Pile) bottom() *Card { return p.cards[0] }

// Patience sort with back links.
func sortIntoPiles(cards []int) []*Pile {
	piles := []*Pile{newPile(&Card{value: cards[0]})}
nextCard:
	for _, card := range cards[1:] {
		for i, pile := range piles {
			if pile.top().value >= card {
				var next *Card
				if i > 0 {
					next = piles[i-1].top()
				}
				pile.add(&Card{card, next})
				continue nextCard
			}
		}
		var next *Card
		if len(piles) != 0 {
			next = piles[len(piles)-1].top()
		}
		piles = append(piles, newPile(&Card{card, next}))
	}
	return piles
}
