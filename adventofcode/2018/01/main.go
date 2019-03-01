package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
	// Simply sum up the frequencies.
	freq := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("invalid input: %q", line)
		}
		freq += n
	}
	fmt.Println(freq)
}

func part2() {
	// Read all frequency shifts into memory.
	var xs []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		x, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("invalid input: %q", line)
		}
		xs = append(xs, x)
	}

	// Cycle through the values forever until a frequency is seen twice.
	seen := make(map[int]bool)
	freq := 0
	for {
		for _, x := range xs {
			freq += x
			if seen[freq] {
				fmt.Println(freq)
				return
			}
			seen[freq] = true
		}
	}
}
