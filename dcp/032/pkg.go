/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Jane Street.

Suppose you are given a table of currency exchange rates, represented as a 2D
array. Determine whether there is a possible arbitrage: that is, whether there
is some sequence of trades you can make, starting with some amount A of any
currency, so that you can end up with some amount greater than A of that
currency.

There are no transaction costs and you can trade fractional quantities.

*/
package dcp032

import (
	"math"
)

/*

Notes:

Property of logarithms: ln(xy) = ln(x) + ln(y)
Corollary: convert rates to graph whose weights are -1 * ln(w).

For each source currency X, apply Bellman-Ford, which works w/ negative
weights, to find shortest distance to every other currency. If path to source
currency is negative, there is exists an arbitrage opportunity from X.

Explanation: if -ln(x)+-ln(y)+...-ln(x)<0 then ln(xy...x)>0 which means the
product of the weights, xy...x, is greater than one.

*/

// Reports whether there is an arbitrage opportunity given exchange rates.
// rates[x][y] is the exchange rate from currency x to currency y.
func arbitrage(rates [][]float64) bool {
	graph := negLnMat(rates)
	for i := range graph {
		dist, _ := bellmanFord(graph, i)
		if dist[i] < 0 {
			return true
		}
	}
	return false
}

// Creates a new matrix with corresponding values -ln(x).
func negLnMat(a [][]float64) [][]float64 {
	b := make([][]float64, len(a))
	for i := range b {
		b[i] = make([]float64, len(a[i]))
		for j, v := range a[i] {
			b[i][j] = -math.Log(v)
		}
	}
	return b
}

// Straight Outta Wikipedia.
func bellmanFord(g [][]float64, src int) ([]float64, []int) {
	numVertices := len(g)
	dist := make([]float64, numVertices)
	pred := make([]int, numVertices)
	for i := range g {
		dist[i] = math.Inf(1) // to start, all vertices have infinite weight
		pred[i] = -1          // and a null predecessor.
	}
	dist[src] = 0 // weight is zero at source vertex

	for i := 1; i < numVertices; i++ {
		for u := range g {
			for v := range g {
				if u != v {
					w := dist[u] + g[u][v]
					if w < dist[v] {
						dist[v] = w
						pred[v] = u
					}
				}
			}
		}
	}
	return dist, pred
}
