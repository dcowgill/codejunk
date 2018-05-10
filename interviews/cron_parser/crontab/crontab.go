// Package crontab provides a parser for crontab files.
package crontab

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Job is the result of parsing a single crontab line.
type Job struct {
	Minute   []int  // minutes in the hour: 0-59
	Hour     []int  // hours in the day: 0-23
	MonthDay []int  // days of the month: 1-31
	Month    []int  // months of the year: 1-12
	Weekday  []int  // days of the week: 0-6 (Sunday to Saturday)
	Command  string // the command to execute
}

// Reports the cron schedule in job in a format suitable for testing.
// Not intended for human consumption.
func (job *Job) dump() string {
	fields := make([]string, 0, 5)
	if len(job.Minute) != 0 {
		fields = append(fields, joinInts(job.Minute, ","))
	}
	if len(job.Hour) != 0 {
		fields = append(fields, joinInts(job.Hour, ","))
	}
	if len(job.MonthDay) != 0 {
		fields = append(fields, joinInts(job.MonthDay, ","))
	}
	if len(job.Month) != 0 {
		fields = append(fields, joinInts(job.Month, ","))
	}
	if len(job.Weekday) != 0 {
		fields = append(fields, joinInts(job.Weekday, ","))
	}
	return strings.Join(fields, " ")
}

// Joins the ints in xs with sep.
func joinInts(xs []int, sep string) string {
	var s, t string
	for _, x := range xs {
		s += t
		s += strconv.Itoa(x)
		t = sep
	}
	return s
}

// Matches one or more consecutive space characters.
var ws = regexp.MustCompile(`\s+`)

// ParseJob parses a single crontab line.
func ParseJob(input string) (job *Job, err error) {
	// Note: rather than check for errors throughout, all parsing functions
	// called by this one simply panic if they encounter a problem; the panics
	// are caught here and translated into normal errors.
	fail := func(cause interface{}) error {
		return fmt.Errorf("parse failed: %s", cause)
	}
	defer func() {
		if r := recover(); r != nil {
			err = fail(r)
		}
	}()
	fields := ws.Split(strings.TrimSpace(input), 6)
	if len(fields) != 6 {
		return nil, fail(fmt.Errorf("got %d fields, want 6", len(fields)))
	}
	return &Job{
		Minute:   expand(clamp(parseField(fields[0], nil), 0, 59)),
		Hour:     expand(clamp(parseField(fields[1], nil), 0, 23)),
		MonthDay: expand(clamp(parseField(fields[2], nil), 1, 31)),
		Month:    expand(clamp(parseField(fields[3], trMonth), 1, 12)),
		Weekday:  expand(clamp(parseField(fields[4], trWeekday), 0, 6)),
		Command:  fields[5],
	}, nil
}

// Parses a field, which comprises one or more comma-separated ranges.
//
// tr is an optional translation from case-insensitive strings to ints; it is
// applied to every range during parsing (see parseRange).
func parseField(input string, tr translations) []intRange {
	fields := strings.Split(input, ",")
	ranges := make([]intRange, 0, len(fields))
	for _, s := range fields {
		if s != "" {
			ranges = append(ranges, parseRange(s, tr))
		}
	}
	return ranges
}

// Parses a range. Grammar, after tr has been applied:
//
//	field	= values [step]
//	values	= "*" | range
//	range	= INT "-" INT
//	step	= "/" INT
//
func parseRange(input string, tr translations) intRange {
	// Parses s as a base 10 int, after applying translations.
	// Panics if the parse fails or if result is under minval.
	parseInt := func(s string, minval int) int {
		var n int64
		if t, ok := tr.lookup(s); ok {
			s = t
		}
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil || int(n) < minval {
			panic(fmt.Errorf("invalid range %q: expected integer >= %d, got %q",
				input, minval, s))
		}
		return int(n)
	}
	// After splitting on a slash, there must be 1 or 2 fields.
	parts := strings.Split(input, "/")
	if len(parts) > 2 {
		panic(fmt.Errorf("invalid range: %q", input))
	}
	// Parse the 'step' non-terminal, if present.
	r := intRange{}
	if len(parts) == 2 {
		r.step = parseInt(parts[1], 1)
	}
	// Parse the 'values' non-terminal.
	values := parts[0]
	if values == "*" {
		// An asterisk: include all values.
		r.low, r.high = 0, 9999
	} else if dash := strings.Index(values, "-"); dash >= 1 {
		// A dash: parse as "low-high".
		r.low = parseInt(values[:dash], 0)
		r.high = parseInt(values[dash+1:], 0)
	} else {
		// There are two possibilities: "X" or "X/Y".
		// In the first case, low = high = X.
		// In the second case, a range is implied: "X-maxval/Y".
		r.low = parseInt(values, 0)
		r.high = r.low
		if r.step != 0 {
			r.high = 9999
		}
	}
	return r
}

// A range of values in a crontab line.
// N.B. step may be zero, in which case it must be treated as 1.
type intRange struct{ low, high, step int }

// Ensures each range in rs fits into the range [low, high]. Also handles
// wrap-around ranges (i.e. where low>high) by splitting them in two.
func clamp(rs []intRange, low, high int) []intRange {
	result := make([]intRange, 0, len(rs))
	for _, r := range rs {
		l, h := max(low, r.low), min(high, r.high)
		if l > h { // wrap-around
			result = append(result, intRange{low: l, high: high, step: r.step},
				intRange{low: low, high: h, step: r.step})
		} else {
			result = append(result, intRange{low: l, high: h, step: r.step})
		}
	}
	return result
}

// Returns the union of the ints in all ranges.
func expand(rs []intRange) []int {
	var set intSet
	for _, r := range rs {
		for i := r.low; i <= r.high; i += max(r.step, 1) {
			set.add(i)
		}
	}
	return set.slice()
}

// Set of ints. Only supports elements in the range [0, 63], but that's good
// enough for our widest unit: minutes (0-59).
type intSet int64

// Adds n to the set x.
func (x *intSet) add(n int) { *x |= (1 << uint(n)) }

// Reports whether n is in the set x.
func (x *intSet) has(n int) bool { return *x&(1<<uint(n)) != 0 }

// Reports the number of elements in the set.
func (x *intSet) count() int {
	n := 0
	y := *x
	for y != 0 {
		y &= (y - 1)
		n++
	}
	return n
}

// Returns all set members.
func (x *intSet) slice() []int {
	a := make([]int, 0, x.count())
	for i := 0; i < 64; i++ {
		if x.has(i) {
			a = append(a, i)
		}
	}
	return a
}

// Reports the lesser of a and b.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Reports the greater of a and b.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Maps one set of strings to another. Gets applied to field ranges before
// integer parsing takes place, e.g. so that "Feb" can be parsed as "2".
type translations map[string]string

// Equivalent to tr[lowercase(s)], but safe when tr is nil.
func (tr translations) lookup(s string) (string, bool) {
	if tr != nil {
		t, ok := tr[strings.ToLower(s)]
		return t, ok
	}
	return "", false
}

var (
	trMonth = translations{
		"jan": "1",
		"feb": "2",
		"mar": "3",
		"apr": "4",
		"may": "5",
		"jun": "6",
		"jul": "7",
		"aug": "8",
		"sep": "9",
		"oct": "10",
		"nov": "11",
		"dec": "12",
	}
	trWeekday = translations{
		"sun": "0",
		"mon": "1",
		"tue": "2",
		"wed": "3",
		"thu": "4",
		"fri": "5",
		"sat": "6",
		"7":   "0", // by convention, both 0 and 7 mean Sunday
	}
)
