package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Edges map[int][]int

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func readEdges(r io.Reader) Edges {
	var (
		edges = make(Edges)
		re    = regexp.MustCompile(`^(\d+)\s+(\d+)`)
		lr    = bufio.NewReader(r)
	)
	for {
		s, err := lr.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		matches := re.FindStringSubmatch(s)
		if matches != nil {
			v, u := atoi(matches[1]), atoi(matches[2])
			edges[v] = append(edges[v], u)
			edges[u] = append(edges[u], v)
		}
	}
	return edges
}

func findcc(edges Edges) []int {
	var (
		frontier = list.New()
		cc       []int
	)
	// No edges, no subgraph.
	if len(edges) == 0 {
		return nil
	}
	// Push the "first" vertex onto the frontier.
	for v, _ := range edges {
		frontier.PushBack(v)
		break
	}
	// Breadth-first search until we've exhausted the frontier.
	for e := frontier.Front(); e != nil; e = frontier.Front() {
		v := frontier.Remove(e).(int)
		if neighbors, ok := edges[v]; ok {
			cc = append(cc, v)
			delete(edges, v)
			// Add v's neighbors to the frontier.
			for _, u := range neighbors {
				frontier.PushBack(u)
			}
		}
	}
	return cc
}

func main() {
	edges := readEdges(os.Stdin)
	for len(edges) != 0 {
		cc := findcc(edges)
		fmt.Println(cc)
	}
}
