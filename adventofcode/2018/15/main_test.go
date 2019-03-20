package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	var tests = []struct {
		input     string
		final     string
		numRounds int
		totalHP   int
	}{
		{
			input: `
#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`,
			final: `
#######
#G....#
#.G...#
#.#.#G#
#...#.#
#....G#
#######`,
			numRounds: 47,
			totalHP:   590,
		},
		{
			input: `
#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`,
			final: `
#######
#...#E#
#E#...#
#.E##.#
#E..#E#
#.....#
#######`,
			numRounds: 37,
			totalHP:   982,
		},
		{
			input: `
#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`,
			final: `
#######
#.E.E.#
#.#E..#
#E.##.#
#.E.#.#
#...#.#
#######`,
			numRounds: 46,
			totalHP:   859,
		},
		{
			input: `
#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`,
			final: `
#######
#G.G#.#
#.#G..#
#..#..#
#...#G#
#...G.#
#######`,
			numRounds: 35,
			totalHP:   793,
		},
		{
			input: `
#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`,
			final: `
#######
#.....#
#.#G..#
#.###.#
#.#.#.#
#G.G#G#
#######`,
			numRounds: 54,
			totalHP:   536,
		},
		{
			input: `
#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`,
			final: `
#########
#.G.....#
#G.G#...#
#.G##...#
#...##..#
#.G.#...#
#.......#
#.......#
#########`,
			numRounds: 20,
			totalHP:   937,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			world := loadWorldFromString(tt.input)
			numRounds := runUntilDone(world)
			actualMap := strings.TrimSpace(world.String())
			expectedMap := strings.TrimSpace(tt.final)
			passed := actualMap == expectedMap && numRounds == tt.numRounds &&
				world.totalHP() == tt.totalHP
			if passed {
				return
			}
			var b strings.Builder
			fmt.Fprintf(&b, "end state map:\n%s\n\n", actualMap)
			fmt.Fprintf(&b, "expected:\n%s\n\n", expectedMap)
			fmt.Fprintf(&b, "took %d rounds; expected %d\n", numRounds, tt.numRounds)
			fmt.Fprintf(&b, "total HP left was %d; expected %d\n\n", world.totalHP(), tt.totalHP)
			t.Fatal(b.String())
		})
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		input     string
		final     string
		numRounds int
		totalHP   int
		bonus     int
	}{
		{
			input: `
#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`,
			final: `
#######
#..E..#
#...E.#
#.#.#.#
#...#.#
#.....#
#######`,
			numRounds: 29,
			totalHP:   172,
			bonus:     12,
		},
		{
			input: `
#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`,
			final: `
#######
#.E.E.#
#.#E..#
#E.##E#
#.E.#.#
#...#.#
#######`,
			numRounds: 33,
			totalHP:   948,
			bonus:     1,
		},
		{
			input: `
#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`,
			final: `
#######
#.E.#.#
#.#E..#
#..#..#
#...#.#
#.....#
#######`,
			numRounds: 37,
			totalHP:   94,
			bonus:     12,
		},
		{
			input: `
#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`,
			final: `
#######
#...E.#
#.#..E#
#.###.#
#.#.#.#
#...#.#
#######`,
			numRounds: 39,
			totalHP:   166,
			bonus:     9,
		},
		{
			input: `
#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`,
			final: `
#########
#.......#
#.E.#...#
#..##...#
#...##..#
#...#...#
#.......#
#.......#
#########`,
			numRounds: 30,
			totalHP:   38,
			bonus:     31,
		},
	}
	const maxBonus = 1000 // a reasonable upper bound
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			world := loadWorldFromString(tt.input)
			numElvesAtStart := world.numUnitsByKind(elfKind)
			for bonus := 0; bonus <= maxBonus; bonus++ {
				world := world.clone()
				world.elfBonus = bonus
				numRounds := runUntilDone(world)
				if world.numUnitsByKind(elfKind) != numElvesAtStart {
					continue // one or more elves were killed; need more bonus damage
				}
				actualMap := strings.TrimSpace(world.String())
				expectedMap := strings.TrimSpace(tt.final)
				passed := actualMap == expectedMap && numRounds == tt.numRounds &&
					world.totalHP() == tt.totalHP && bonus == tt.bonus
				if passed {
					return
				}
				var b strings.Builder
				fmt.Fprintf(&b, "end state map:\n%s\n\n", actualMap)
				fmt.Fprintf(&b, "expected:\n%s\n\n", expectedMap)
				fmt.Fprintf(&b, "took %d rounds; expected %d\n", numRounds, tt.numRounds)
				fmt.Fprintf(&b, "elf bonus was %d; expected %d\n", bonus, tt.bonus)
				fmt.Fprintf(&b, "total HP left was %d; expected %d\n\n", world.totalHP(), tt.totalHP)
				t.Fatal(b.String())
			}
		})
	}
}

func loadWorldFromString(s string) *World {
	return readInput(strings.NewReader(strings.TrimSpace(s)))
}
