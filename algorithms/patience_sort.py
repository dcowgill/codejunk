#!/usr/bin/env python3

import sys

class Card:
    def __init__(self, value, back_link):
        self.value = value
        self.back_link = back_link

    def __repr__(self):
        return str(self.value)

def top_card(pile):
    return pile[-1]

def sort_into_piles(cards):
    piles = [[Card(cards[0], None)]]
    cards = cards[1:]
    for card in cards:
        for i, pile in enumerate(piles):
            if top_card(pile).value >= card:
                back_link = top_card(piles[i-1]) if i > 0 else None
                pile.append(Card(card, back_link))
                break
        else:
            back_link = top_card(piles[-1]) if piles else None
            piles.append([Card(card, back_link)])
    return piles

def merge(piles):
    cards = []
    while piles:
        n = 0
        mincard = piles[n][-1].value
        for i, pile in enumerate(piles):
            if top_card(pile).value < mincard:
                n, mincard = i, top_card(pile).value
        cards.append(mincard)
        piles[n] = piles[n][:-1]
        if not piles[n]:
            piles[-1], piles[n] = piles[n], piles[-1]
            piles = piles[:-1]
    return cards

def longest_increasing_subsequence(piles):
    p = piles[-1][0] # start with bottom card of rightmost pile
    sequence = []
    while p:
        sequence.append(p)
        p = p.back_link
    return [card.value for card in reversed(sequence)]

def main():
    if len(sys.argv) == 1:
        print("usage: patience_sort.py <number> [<number>...]")
        sys.exit(1)
    cards = [int(s) for s in sys.argv[1:]]
    piles = sort_into_piles(cards)
    li_seq = longest_increasing_subsequence(piles)
    sorted_cards = merge(piles)
    print("initial cards =", cards)
    print("piles =", piles)
    print("longest increasing subsequence =", li_seq)
    print("sorted cards =", sorted_cards)

if __name__ == '__main__':
    main()
