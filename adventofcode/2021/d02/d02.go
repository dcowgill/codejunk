package d02

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"adventofcode2021/lib"
)

func Run() {
	lib.Run(2, part1, part2)
}

func part1() int64 {
	var depth, pos int64
	foreachMove(realInput, func(m move) {
		switch m.dir {
		case "forward":
			pos += m.n
		case "up":
			depth -= m.n
		case "down":
			depth += m.n
		}
	})
	return pos * depth
}

func part2() int64 {
	var depth, pos, aim int64
	foreachMove(realInput, func(m move) {
		switch m.dir {
		case "forward":
			pos += m.n
			depth += aim * m.n
		case "up":
			aim -= m.n
		case "down":
			aim += m.n
		}
	})
	return pos * depth
}

func foreachMove(input []string, f func(m move)) {
	for _, s := range input {
		f(parseMove(s))
	}
}

type move struct {
	dir string
	n   int64
}

func parseMove(s string) move {
	parts := strings.SplitN(s, " ", 2)
	if len(parts) != 2 {
		log.Fatalf("error parsing move %q: expected 2 parts, got %d", s, len(parts))
	}
	return move{
		dir: parts[0],
		n:   parseInt64(parts[1]),
	}
}

func parseInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("couldn't convert %q to int64: %v", s, err))
	}
	return n
}
