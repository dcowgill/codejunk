/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Amazon.

You are given a list of data entries that represent entries and exits of groups
of people into a building. An entry looks like this:

{"timestamp": 1526579928, count: 3, "type": "enter"}

This means 3 people entered the building. An exit looks like this:

{"timestamp": 1526580382, count: 2, "type": "exit"}

This means that 2 people exited the building. timestamp is in Unix time.

Find the busiest period in the building, that is, the time with the most people
in the building. Return it as a pair of (start, end) timestamps. You can assume
the building always starts off and ends up empty, i.e. with 0 people inside.

*/
package dcp171

import (
	"sort"
	"time"
)

// Event defines an entry into or exit from the building.
type Event struct {
	when    time.Time // When this event occurred.
	isEntry bool      // Are the people entering or exiting the building?
	count   int       // Number of people entering/exiting.
}

// A syntactic convenience.
type Timespan struct {
	begin, end time.Time
}

// Returns the timespan during which the building was at maximum occupancy.
// If there are multiple such spans of maximum occupancy, returns the first.
// All events must have count >= 1.
// Assumes the building starts and ends with zero people.
func findPeriodOfMaxOccupancy(events []Event) Timespan {
	// The problem does not state that the input events are in any particular
	// order, so first ensure they are chronologically sorted.
	// Also, copy the slice to avoid touching the input.
	sortedEvents := make([]Event, len(events))
	copy(sortedEvents, events)
	sort.Slice(sortedEvents, func(i, j int) bool {
		return sortedEvents[i].when.Before(sortedEvents[j].when)
	})

	// Simulate the events.
	var (
		cur  int      // current building population
		max  int      // maximum population seen
		span Timespan // earliest timespan during which population was equal to max
	)
	for i, event := range sortedEvents {
		if !event.isEntry {
			cur -= event.count
			continue // An exit event cannot produce a new maximum.
		}
		cur += event.count
		if cur > max {
			// The expression "events[i+1]" is safe here because the building
			// always ends up empty, i.e. we can't reach maximum occupancy on
			// the last event. And anyway, we want to panic if that invariant is
			// violated. (N.B. this assumes there are no events with count==0.)
			span.begin, span.end = event.when, events[i+1].when
		}
	}

	return span
}
