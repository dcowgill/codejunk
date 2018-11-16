package main

import (
	"fmt"
	"sort"
)

type (
	vertex    string
	vertexSet map[vertex]bool
	edgeSet   map[vertex]vertexSet
	graph     struct{ edges, incoming edgeSet }
)

func (s vertexSet) add(v vertex)           { s[v] = true }
func (s vertexSet) remove(v vertex)        { delete(s, v) }
func (s vertexSet) contains(v vertex) bool { return s[v] }
func (s vertexSet) size() int              { return len(s) }

func (es edgeSet) add(v, w vertex) {
	vs := es[v]
	if vs == nil {
		vs = make(vertexSet)
		es[v] = vs
	}
	vs.add(w)
}

func (es edgeSet) remove(v, w vertex) { delete(es[v], w) }

func newGraph() *graph {
	return &graph{
		edges:    make(map[vertex]vertexSet),
		incoming: make(map[vertex]vertexSet),
	}
}

func (g *graph) addEdge(src, dst vertex) {
	g.edges.add(src, dst)
	g.incoming.add(dst, src)
}

func (g *graph) removeEdge(src, dst vertex) {
	g.edges.remove(src, dst)
	g.incoming.remove(dst, src)
}

func (g *graph) hasIncomingEdge(v vertex) bool { return g.incoming[v].size() != 0 }

func (g *graph) numEdges() int {
	n := 0
	for _, vs := range g.edges {
		n += vs.size()
	}
	return n
}

func kahnsAlgo(g *graph) []vertex {
	var S []vertex
	for src := range g.edges {
		if !g.hasIncomingEdge(src) {
			S = append(S, src)
		}
	}
	var L []vertex
	for len(S) != 0 {
		sort.Slice(S, func(i, j int) bool { return S[i] < S[j] }) // ensure consistent output
		node := S[0]
		S = S[1:]
		L = append(L, node)
		for m := range g.edges[node] {
			g.removeEdge(node, m)
			if !g.hasIncomingEdge(m) {
				S = append(S, m)
			}
		}
	}
	if g.numEdges() != 0 {
		panic("graph contains a cycle")
	}
	return L
}

func parse(seqs [][]string) *graph {
	g := newGraph()
	for _, seq := range seqs {
		for i := 0; i < len(seq)-1; i++ {
			src, dst := vertex(seq[i]), vertex(seq[i+1])
			g.addEdge(src, dst)
		}
	}
	return g
}

func main() {
	graph := parse([][]string{
		{"B", "A", "E", "D"}, // B->A, A->E, etc. are edges
		{"B", "C", "E", "D", "F"},
		{"B", "A", "C", "D"},
		{"E", "F"},
		{"X", "E", "F"},
	})
	fmt.Printf("graph = %+v\n", graph)
	L := kahnsAlgo(graph)
	fmt.Printf("L = %q\n", L)
}
