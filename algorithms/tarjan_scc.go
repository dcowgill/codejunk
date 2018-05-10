// https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm
package main

import (
	"encoding/json"
	"fmt"
)

// graph is a directed graph.
type graph struct {
	edges [][]int
}

// add adds an edge from v to u.
func (g *graph) add(v, u int) {
	if g.has(v, u) {
		return
	}
	for len(g.edges) <= v {
		g.edges = append(g.edges, nil)
	}
	g.edges[v] = append(g.edges[v], u)
}

// Reports whether g contains an edge from v to u.
func (g *graph) has(v, u int) bool {
	for _, x := range g.successors(v) {
		if x == u {
			return true
		}
	}
	return false
}

// Reports the number of vertices in the graph.
func (g *graph) numVertices() int {
	return len(g.edges)
}

// Gets the vertices to which v has an edge.
func (g *graph) successors(v int) []int {
	if len(g.edges) <= v {
		return nil
	}
	return g.edges[v]
}

type vertex struct {
	id      int
	index   int
	lowLink int
	onStack bool
}

func (v *vertex) isIndexUndefined() bool {
	return v.index == 0
}

// n is the number of vertices. vertices are numbered [0, n).
func tarjan(g *graph) [][]*vertex {
	n := g.numVertices()
	index := 0
	var s []*vertex
	// s := make([]bool, n) // stack
	vertices := make([]*vertex, n)
	for i := range vertices {
		vertices[i] = &vertex{id: i}
	}

	var sccs [][]*vertex
	var strongConnect func(v *vertex)
	strongConnect = func(v *vertex) {
		v.index = index
		v.lowLink = index
		index++
		s = append(s, v)
		v.onStack = true

		// Consider successors of v
		for _, x := range g.successors(v.id) {
			w := vertices[x]
			if w.isIndexUndefined() {
				// Successor w has not yet been visited; recurse on it
				strongConnect(w)
				v.lowLink = min(v.lowLink, w.lowLink)
			} else if w.onStack {
				// Successor w is in stack S and hence in the current SCC
				v.lowLink = min(v.lowLink, w.index)
			}
		}

		// If v is a root node, pop the stack and generate a
		// strongly connected component.
		if v.lowLink == v.index {
			var w *vertex
			var scc []*vertex
			for v != w {
				w, s = s[len(s)-1], s[:len(s)-1]
				w.onStack = false
				scc = append(scc, w)
			}
			sccs = append(sccs, scc)
		}
	}

	for _, v := range vertices {
		if v.isIndexUndefined() {
			strongConnect(v)
		}
	}

	return sccs
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	g := new(graph)

	g.add(0, 1)
	g.add(1, 2)
	g.add(1, 3)
	g.add(3, 4)
	g.add(4, 5)
	g.add(4, 6)
	g.add(5, 4)
	g.add(6, 7)
	g.add(7, 3)

	fmt.Println(dumps(tarjan(g)))
}

// Quick-and-dirty value-to-json string.
func dumps(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "    ")
	return string(data)
}
