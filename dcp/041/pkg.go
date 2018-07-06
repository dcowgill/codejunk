/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Given an unordered list of flights taken by someone, each represented as
(origin, destination) pairs, and a starting airport, compute the person's
itinerary. If no such itinerary exists, return null. If there are multiple
possible itineraries, return the lexicographically smallest one. All flights
must be used in the itinerary.

For example, given the list of flights [('SFO', 'HKO'), ('YYZ', 'SFO'), ('YUL',
'YYZ'), ('HKO', 'ORD')] and starting airport 'YUL', you should return the list
['YUL', 'YYZ', 'SFO', 'HKO', 'ORD'].

Given the list of flights [('SFO', 'COM'), ('COM', 'YYZ')] and starting airport
'COM', you should return null.

Given the list of flights [('A', 'B'), ('A', 'C'), ('B', 'C'), ('C', 'A')] and
starting airport 'A', you should return the list ['A', 'B', 'C', 'A', 'C'] even
though ['A', 'C', 'A', 'B', 'C'] is also a valid itinerary. However, the first
one is lexicographically smaller.

*/
package dcp041

import (
	"container/heap"
)

type airport string

type flight [2]airport

func (f flight) from() airport { return f[0] }
func (f flight) to() airport   { return f[1] }

type airportSet []airport

func (vs airportSet) add(v airport) airportSet {
	for _, u := range vs {
		if u == v {
			return vs
		}
	}
	return append(vs, v)
}

type flightGraph struct{ edges map[airport]airportSet }

func newFlightGraph(flights []flight) *flightGraph {
	g := flightGraph{make(map[airport]airportSet)}
	for _, f := range flights {
		g.edges[f.from()] = g.edges[f.from()].add(f.to())
	}
	return &g
}
func (g *flightGraph) destinations(from airport) []airport {
	return g.edges[from]
}
func (g *flightGraph) numFlights() int {
	n := 0
	for _, s := range g.edges {
		n += len(s)
	}
	return n
}

type itinerary []airport

func newItinerary(f flight) itinerary     { return itinerary{f.from(), f.to()} }
func (it itinerary) currAirport() airport { return it[len(it)-1] }
func (it itinerary) numFlights() int      { return len(it) - 1 }
func (it itinerary) hasFlight(f flight) bool {
	for i := 1; i < len(it); i++ {
		if it[i-1] == f.from() && it[i] == f.to() {
			return true
		}
	}
	return false
}
func (it itinerary) addFlight(f flight) itinerary {
	it2 := make(itinerary, len(it)+1)
	copy(it2, it)
	it2[len(it)] = f.to()
	return it2
}
func (it itinerary) less(it2 itinerary) bool {
	for i, x := range it {
		if i >= len(it2) {
			return false
		}
		y := it2[i]
		switch {
		case x < y:
			return true
		case x > y:
			return false
		}
	}
	return len(it) < len(it2)
}

type frontier []itinerary

func (h frontier) Len() int            { return len(h) }
func (h frontier) Less(i, j int) bool  { return h[i].less(h[j]) }
func (h frontier) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *frontier) Push(x interface{}) { *h = append(*h, x.(itinerary)) }
func (h *frontier) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findItinerary(g *flightGraph, start airport) itinerary {
	var frontier frontier
	for _, dst := range g.destinations(start) {
		heap.Push(&frontier, newItinerary(flight{start, dst}))
	}
	for frontier.Len() != 0 {
		it := heap.Pop(&frontier).(itinerary)
		src := it.currAirport()
		for _, dst := range g.destinations(src) {
			f := flight{src, dst}
			if !it.hasFlight(f) {
				newIt := it.addFlight(f)
				if newIt.numFlights() == g.numFlights() {
					return newIt
				}
				heap.Push(&frontier, newIt)
			}
		}
	}
	return nil
}
