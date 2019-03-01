package main

import (
	"flag"
	"fmt"
)

func main() {
	part := flag.Int("part", 1, "which part")
	flag.Parse()
	switch *part {
	case 1:
		part1()
	case 2:
		part2()
	}
}

func part1() {
	fmt.Println(highScore(458, 71307))
}

func part2() {
	fmt.Println(highScore(458, 71307*100))
}

func highScore(numPlayers, numMarbles int) int {
	var (
		scores  = make([]int, numPlayers)
		marbles = newCircle(0)
		player  = 0
	)
	for i := 1; i <= numMarbles; i++ {
		var score int
		marbles, score = addMarble(marbles, i)
		scores[player] += score
		player = (player + 1) % numPlayers
	}
	max := 0
	for _, v := range scores {
		if v > max {
			max = v
		}
	}
	return max
}

type node struct {
	prev  *node // counter-clockwise
	next  *node // clockwise
	value int
}

func newCircle(value int) *node {
	n := &node{value: value}
	n.prev = n
	n.next = n
	return n
}

// Adds a marble with the specified value to the circle. Returns the new current
// marble and how much to add to the current player's score.
func addMarble(curr *node, value int) (head *node, score int) {
	if value%23 == 0 {
		p := curr
		for i := 0; i < 7; i++ {
			p = p.prev
		}
		p.prev.next, p.next.prev = p.next, p.prev
		return p.next, value + p.value
	}
	p := curr.next
	q := p.next
	n := &node{prev: p, next: q, value: value}
	p.next = n
	q.prev = n
	return n, 0
}
