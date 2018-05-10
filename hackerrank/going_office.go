package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func readInput(r io.Reader) <-chan []int {
	var (
		re = regexp.MustCompile(`\d+`)
		ch = make(chan []int)
		br = bufio.NewReader(r)
	)
	go func() {
		defer close(ch)
		for {
			s, err := br.ReadString('\n')
			if err == io.EOF {
				return
			} else if err != nil {
				panic(err)
			}
			var a []int
			for _, m := range re.FindAllString(s, -1) {
				i, err := strconv.Atoi(m)
				if err != nil {
					panic(err)
				}
				a = append(a, i)
			}
			ch <- a
		}
	}()
	return ch
}

type graph struct {
	numVertexes int
	numEdges    int
	edges       map[int]map[int]int
	src, dst    int
	queries     [][]int
}

func newGraph() *graph {
	g := graph{}
	g.edges = make(map[int]map[int]int, 0)
	return &g
}

func (g *graph) addDirectedEdge(u, v, w int) {
	if m, ok := g.edges[u]; ok {
		m[v] = w
	} else {
		g.edges[u] = map[int]int{v: w}
	}
}

func (g *graph) addEdge(u, v, w int) {
	g.addDirectedEdge(u, v, w)
	g.addDirectedEdge(v, u, w)
}

func (g *graph) setEdgeWeight(u, v, w int) int {
	oldW := g.edges[u][v]
	g.edges[u][v] = w
	g.edges[v][u] = w
	return oldW
}

func parseInput(c <-chan []int) *graph {
	g := newGraph()
	a := <-c
	g.numVertexes, g.numEdges = a[0], a[1]
	for i := 0; i < g.numEdges; i++ {
		a = <-c
		g.addEdge(a[0], a[1], a[2])
	}
	a = <-c
	g.src, g.dst = a[0], a[1]
	a = <-c
	numQueries := a[0]
	for i := 0; i < numQueries; i++ {
		g.queries = append(g.queries, <-c)
	}
	return g
}

const MAXDIST = 2147483647

type edge struct{ u, v int }
type edgeSet map[edge]bool

type item struct {
	v       int  // vertex ID
	dist    int  // distance to here
	index   int  // index of item in the heap
	visited bool // has node been visited?
}

type priorityQueue []*item

func (q priorityQueue) Len() int { return len(q) }

func (q priorityQueue) Less(i, j int) bool {
	return q[i].dist < q[j].dist
}

func (q priorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *priorityQueue) Push(x interface{}) {
	a := *q
	n := len(a)
	a = a[0 : n+1]
	item := x.(*item)
	item.index = n
	a[n] = item
	*q = a
}

func (q *priorityQueue) Pop() interface{} {
	a := *q
	n := len(a)
	item := a[n-1]
	item.index = -1 // for safety
	*q = a[0 : n-1]
	return item
}

func (q *priorityQueue) setDist(i *item, dist int) {
	heap.Remove(q, i.index)
	i.dist = dist
	heap.Push(q, i)
}

type dijkstra struct {
	g     *graph
	items []item
	q     priorityQueue
	prev  map[int][]int
}

func newDijkstra(g *graph) *dijkstra {
	d := dijkstra{
		g:     g,
		items: make([]item, g.numVertexes),
		q:     make(priorityQueue, g.numVertexes),
		prev:  make(map[int][]int, 0),
	}
	for i := 0; i < g.numVertexes; i++ {
		d.items[i].v = i
		d.q[i] = &d.items[i]
		d.prev[i] = make([]int, 0)
	}
	return &d
}

func (d *dijkstra) shortestPathDistance() int {
	// Reset the items array.
	for i := 0; i < len(d.items); i++ {
		d.items[i].dist = MAXDIST
		d.items[i].visited = false
	}
	d.items[d.g.src].dist = 0

	// Reset the prev array.
	for k, v := range d.prev {
		d.prev[k] = v[:0]
	}

	// Rebuild the priority queue.
	d.q = d.q[:0]
	for i := 0; i < len(d.items); i++ {
		heap.Push(&d.q, &d.items[i])
	}

	// Dijkstra's algorithm.
	for len(d.q) != 0 {
		curr := heap.Pop(&d.q).(*item)
		for u, w := range d.g.edges[curr.v] {
			if !d.items[u].visited {
				newDist := curr.dist + w
				if newDist < d.items[u].dist {
					d.q.setDist(&d.items[u], newDist)
					d.prev[u] = append(d.prev[u][:0], curr.v)
				} else if newDist == d.items[u].dist {
					d.prev[u] = append(d.prev[u], curr.v)
				}
			}
		}
		curr.visited = true
		if curr.v == d.g.dst {
			return curr.dist
		}
	}
	return -1
}

func prevToPaths(d *dijkstra, y, z int) (paths []edgeSet) {
	if len(d.prev[y]) == 0 {
		paths = []edgeSet{edgeSet{}}
		return
	}
	for _, x := range d.prev[y] {
		for _, p := range prevToPaths(d, x, y) {
			p[edge{x, y}] = true
			paths = append(paths, p)
		}
	}
	return
}

func (d *dijkstra) bestPathEdges() edgeSet {
	paths := prevToPaths(d, d.g.dst, -1)
	// intersection of paths
}

func main() {
	g := parseInput(readInput(os.Stdin))
	d := newDijkstra(g)

	// Find edges which are on all best paths.
	bestDist := d.shortestPathDistance()
	paths := prevToPaths(d, g.dst, -1)
	fmt.Println(bestDist)
	fmt.Println(paths)
	//bestPathEdges := d.bestPathEdges()

	// Special case: no path from src->dst.
	if bestDist < 0 {
		for _ = range g.queries {
			fmt.Println("Infinity")
			return
		}
	}

	return

	for _, q := range g.queries {
		u, v := q[0], q[1]
		w := g.setEdgeWeight(u, v, MAXDIST)
		if n := d.shortestPathDistance(); n >= 0 {
			fmt.Println(n)
		} else {
			fmt.Println("Infinity")
		}
		fmt.Println(d)
		g.setEdgeWeight(u, v, w)
	}
}
