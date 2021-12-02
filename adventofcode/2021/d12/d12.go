package d12

import (
	"adventofcode2021/lib"
)

func Run()         { lib.Run(12, part1, part2) }
func part1() int64 { return explore(buildGraph(realInput), simpleVisitor{}) }
func part2() int64 { return explore(buildGraph(realInput), repeatVisitor{}) }

func explore(g graph, v visitor) int64 {
	npaths := 0
	var visit func(cave, visitor)
	visit = func(c cave, v visitor) {
		if c == "end" {
			npaths++
			return
		}
		if v2, ok := v.visit(c); ok {
			for _, dst := range g.edges[c] {
				visit(dst, v2)
			}
		}
	}
	visit("start", v)
	return int64(npaths)
}

type visitor interface {
	visit(cave) (visitor, bool)
}

type simpleVisitor []cave

func (v simpleVisitor) visit(c cave) (visitor, bool) {
	if c.large() || !contains(v, c) {
		return append(v, c), true
	}
	return v, false
}

type repeatVisitor struct {
	path      []cave
	revisited bool
}

func (v repeatVisitor) visit(c cave) (visitor, bool) {
	if c.large() || !contains(v.path, c) {
		return v.add(c), true
	}
	if c == "start" || v.revisited {
		return v, false
	}
	v2 := v.add(c)
	v2.revisited = true
	return v2, true
}

func (v repeatVisitor) add(c cave) repeatVisitor {
	v2 := v
	v2.path = append(v2.path, c)
	return v2
}

func contains(a []cave, v cave) bool {
	for _, u := range a {
		if u == v {
			return true
		}
	}
	return false
}

type cave string

func (c cave) large() bool {
	for i := 0; i < len(c); i++ {
		if c[i] < 'A' || c[i] > 'Z' {
			return false
		}
	}
	return true
}

type graph struct {
	edges map[cave][]cave
}

func buildGraph(input [][]string) graph {
	edges := make(map[cave][]cave)
	for _, edge := range input {
		a, b := cave(edge[0]), cave(edge[1])
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}
	return graph{edges}
}
