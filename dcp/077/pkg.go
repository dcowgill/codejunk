/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Snapchat.

Given a list of possibly overlapping intervals, return a new list of intervals
where all overlapping intervals have been merged.

The input list is not necessarily ordered in any way.

For example, given [(1, 3), (5, 8), (4, 10), (20, 25)], you should return
[(1, 3), (4, 10), (20, 25)].

*/
package dcp077

import (
	"fmt"
	"sort"
)

// Represents the range [start, finish].
type span struct {
	start, finish int
}

// String implements the Stringer interface.
func (s span) String() string {
	return fmt.Sprintf("[%d, %d]", s.start, s.finish)
}

// Reports whether s and t overlap.
func (s span) overlaps(t span) bool {
	return !(s.finish < t.start || s.start > t.finish) // neither strictly before nor strictly after
}

// Returns the smallest span that covers both s and t.
func (s span) merge(t span) span {
	return span{min(s.start, t.start), max(s.finish, t.finish)}
}

// Sorts spans and merges those that overlap.
func merged(spans []span) []span {
	if len(spans) == 0 {
		return nil
	}
	spans = sorted(spans)
	merged := []span{spans[0]}
	for _, s := range spans[1:] {
		last := len(merged) - 1
		if s.overlaps(merged[last]) {
			merged[last] = s.merge(merged[last])
		} else {
			merged = append(merged, s)
		}
	}
	return merged
}

// Returns a copy of spans sorted by time.
func sorted(spans []span) []span {
	spans = copySpans(spans)
	sort.Slice(spans, func(i, j int) bool {
		s := spans[i]
		t := spans[j]
		switch {
		case s.start < t.start:
			return true
		case s.start > t.start:
			return false
		default:
			return s.finish < t.finish
		}
	})
	return spans
}

// Returns a copy of spans.
func copySpans(spans []span) []span {
	spans2 := make([]span, len(spans))
	copy(spans2, spans)
	return spans2
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
