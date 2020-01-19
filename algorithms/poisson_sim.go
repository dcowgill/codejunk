package main

import (
	"container/heap"
	"container/list"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Compute the probability of Poisson event within time x given rate l.
// l is the rate parameter.
// x is time.
func poissonF(l, x float64) float64 {
	return 1 - math.Exp(-l*x)
}

// Compute a random time to next Poisson event given rate l.
// l is the rate parameter.
// t is time.
func poissonR(l float64) float64 {
	return poissonU(l, rand.Float64())
}

// Value of Poisson variable with rate l at value y.
func poissonU(l, y float64) float64 {
	return -math.Log(y) / l
}

// All events know what time they occurred.
type event interface {
	when() float64
}

// Represents the arrival of a new customer.
type customerEvent struct {
	arrival float64
}

func (e customerEvent) when() float64 { return e.arrival }

// Represents the completion of work for a customer.
type workEvent struct {
	arrival float64
	begin   float64
	end     float64
}

func (e workEvent) when() float64 { return e.end }

type simulationParams struct {
	arrivalRate float64 // number of customers who arrive per minute
	serviceRate float64 // number of customers 1 teller can service per minute
	numTellers  int     // total number of tellers
	maxQueueLen int     // max # of customers waiting before new arrivals are discouraged (0 = no max)
}

type simulation struct {
	params simulationParams // initial parameters

	// simulation state
	events       eventHeap     // priority queue; next thing to occur
	queue        customerQueue // customers waiting in line
	curTime      float64       // current time index (minutes)
	availTellers int           // number of available tellers

	// stats
	numArrivals    int     // # of customers who have arrived
	numDiscouraged int     // # of customers who left because queue was too long
	numServiced    int     // # of customers who were serviced by a teller
	totalWaitTime  float64 // total time spent waiting in line by all customers
}

func newSimulation(params simulationParams) *simulation {
	return &simulation{
		params:       params,
		queue:        newCustomerQueue(),
		availTellers: params.numTellers,
	}
}

func (sim *simulation) run(until float64) {
	// Helper: schedule the next customer arrival.
	addNextCust := func() {
		heap.Push(&sim.events, customerEvent{
			arrival: sim.curTime + poissonR(sim.params.arrivalRate),
		})
	}
	addNextCust()
	for sim.curTime < until {
		// Process the next event.
		evt := heap.Pop(&sim.events).(event)
		sim.curTime = evt.when() // advance clock
		switch evt := evt.(type) {
		case customerEvent:
			sim.numArrivals++
			addNextCust()
			if sim.params.maxQueueLen == 0 || sim.queue.len() < sim.params.maxQueueLen {
				sim.queue.push(evt.when())
			} else {
				sim.numDiscouraged++
			}
		case workEvent:
			sim.availTellers++
			sim.numServiced++
			sim.totalWaitTime += evt.begin - evt.arrival
		}
		// Assign any waiting customers to available tellers.
		for sim.availTellers > 0 && sim.queue.len() > 0 {
			heap.Push(&sim.events, workEvent{
				arrival: sim.queue.pop(),
				begin:   sim.curTime,
				end:     sim.curTime + poissonR(sim.params.serviceRate),
			})
			sim.availTellers--
		}
	}
}

func main() {
	// Get some params from the command line.
	var (
		numTellers      int
		numMinutes      int
		maxWaiting      int
		arrivalsPerHour float64
		servicedPerHour float64
	)
	flag.IntVar(&numTellers, "tellers", 1, "number of tellers")
	flag.IntVar(&numMinutes, "minutes", 8*60, "number of minutes to simulate")
	flag.IntVar(&maxWaiting, "maxqueue", 0, "maximum number of customers in queue")
	flag.Float64Var(&arrivalsPerHour, "arrival-rate", 6, "customers arriving per hour")
	flag.Float64Var(&servicedPerHour, "service-rate", 6, "customers serviced per teller-hour")
	flag.Parse()

	// Make sure parameters aren't broken.
	numTellers = maxInt(numTellers, 1)
	numMinutes = maxInt(numMinutes, 1)
	maxWaiting = maxInt(maxWaiting, 1)
	arrivalsPerHour = math.Max(arrivalsPerHour, 0)
	servicedPerHour = math.Max(servicedPerHour, 0)

	// Run the simulation.
	rand.Seed(time.Now().UnixNano())
	sim := newSimulation(simulationParams{
		arrivalRate: arrivalsPerHour / 60.0,
		serviceRate: servicedPerHour / 60.0,
		numTellers:  numTellers,
		maxQueueLen: maxWaiting,
	})
	sim.run(float64(numMinutes))

	// Print a summary.
	pct := func(n, d int) float64 { return 100 * float64(n) / float64(d) }
	fmt.Printf("minutes simulated: %d\n", numMinutes)
	fmt.Printf("tellers simulated: %d\n", numTellers)
	fmt.Printf("total customers: %d\n", sim.numArrivals)
	fmt.Printf("number discouraged: %d (%.1f%%)\n", sim.numDiscouraged, pct(sim.numDiscouraged, sim.numArrivals))
	fmt.Printf("number serviced: %d (%.1f%%)\n", sim.numServiced, pct(sim.numServiced, sim.numArrivals))
	fmt.Printf("mean wait time (m): %.1f \n", sim.totalWaitTime/float64(sim.numServiced))
}

type eventHeap []event

func (h eventHeap) Len() int           { return len(h) }
func (h eventHeap) Less(i, j int) bool { return h[i].when() < h[j].when() }
func (h eventHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *eventHeap) Push(x interface{}) { *h = append(*h, x.(event)) }

func (h *eventHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type customerQueue struct {
	ll *list.List
}

func newCustomerQueue() customerQueue        { return customerQueue{list.New()} }
func (q customerQueue) len() int             { return q.ll.Len() }
func (q customerQueue) push(arrival float64) { q.ll.PushBack(arrival) }
func (q customerQueue) pop() float64 {
	front := q.ll.Front()
	x := front.Value.(float64)
	q.ll.Remove(front)
	return x
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
