package main

// https://www.interviewstreet.com/challenges/dashboard/#problem/4f40dfda620c4

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	MODULO = 1000 * 1000 * 1000
	START  = 1
)

func pop(s []int) ([]int, int) {
	n := len(s)
	return s[:n-1], s[n-1]
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func readInput(r io.Reader) (data [][]int) {
	re := regexp.MustCompile(`^\s*(\d+)\s+(\d+)\s*$`)
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			return
		} else if err != nil {
			log.Fatal(err)
		}
		matches := re.FindStringSubmatch(line)
		p := [...]int{atoi(matches[1]), atoi(matches[2])}
		data = append(data, p[:])
	}
	return
}

type kingdom struct {
	N            int
	Edges        [][]int
	InverseEdges [][]int
	Reachable    []bool
}

// Taken verbatim from http://en.wikipedia.org/wiki/Topological_sort#Algorithms
func topologicalSort(k *kingdom) []int {
	// Compute per-vertex inbound edge counts.
	in := make([]int, k.N+1)
	for x := 1; x <= k.N; x++ {
		if k.Reachable[x] {
			for _, y := range k.Edges[x] {
				in[y]++
			}
		}
	}

	// Ensure there isn't a cycle including the start vertex.
	if in[START] != 0 {
		return nil
	}

	var (
		L = make([]int, 0)
		S = []int{START}
		x int
	)
	for len(S) != 0 {
		S, x = pop(S)
		L = append(L, x)
		if x == k.N {
			return L // success
		}
		for _, y := range k.Edges[x] {
			if in[y]--; in[y] == 0 {
				S = append(S, y)
			}
		}
	}
	return nil // cycle
}

// Dynamic programming solution, given topological ordering
func countPaths(k *kingdom, vertexOrder []int) int {
	paths := make([]int, k.N+1)
	paths[k.N] = 1
	for i := len(vertexOrder) - 1; i >= 0; i-- {
		x := vertexOrder[i]
		for _, y := range k.InverseEdges[x] {
			paths[y] = (paths[y] + paths[x]) % MODULO
		}
	}
	return paths[START]
}

func main() {
	var (
		data = readInput(os.Stdin)
		N    = data[0][0]
	)

	k := kingdom{
		N:            N,
		Edges:        make([][]int, N+1),
		InverseEdges: make([][]int, N+1),
		Reachable:    make([]bool, N+1),
	}

	// Populate edge data structures.
	for _, edge := range data[1:] {
		x, y := edge[0], edge[1]
		k.Edges[x] = append(k.Edges[x], y)
		k.InverseEdges[y] = append(k.InverseEdges[y], x)
	}

	// Find the set of vertexes reachable from START.
	k.Reachable[START] = true
	frontier := []int{START}
	for len(frontier) != 0 {
		var v int
		frontier, v = pop(frontier)
		for _, u := range k.Edges[v] {
			if !k.Reachable[u] {
				k.Reachable[u] = true
				frontier = append(frontier, u)
			}
		}
	}

	// Ensure vertex N is reachable.
	if !k.Reachable[k.N] {
		fmt.Println(0)
		return
	}

	// Get vertexes in topologically sorted order.
	order := topologicalSort(&k)

	// Count paths backwards from destination vertex.
	if order != nil {
		fmt.Println(countPaths(&k, order))
	} else {
		fmt.Println("INFINITE PATHS")
	}
}
