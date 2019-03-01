package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	part := flag.Int("part", 1, "which part")
	flag.Parse()
	switch *part {
	case 1:
		part1()
	case 2:
		part2()
	}
}

func part1() {
	db := readSleepHistories()
	var topGuardID, topTotal, topMode int
	for guardID, history := range db {
		total := history.totalSlept()
		if total > topTotal {
			topGuardID, topTotal = guardID, total
			topMode, _ = history.modeFreq()
		}
	}
	fmt.Println(topGuardID * topMode)
}

func part2() {
	db := readSleepHistories()
	var topGuardID, topMode, topFreq int
	for guardID, history := range db {
		mode, freq := history.modeFreq()
		if freq > topFreq {
			topGuardID, topMode, topFreq = guardID, mode, freq
		}
	}
	fmt.Println(topGuardID * topMode)
}

func readSleepHistories() map[int]*sleepHistory {
	// Read all events into memory, then sort chronologically.
	scanner := bufio.NewScanner(os.Stdin)
	var events []*event
	for scanner.Scan() {
		events = append(events, parseEvent(scanner.Text()))
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].when.Before(events[j].when)
	})
	// Assemble sleep histories for all guards.
	db := make(map[int]*sleepHistory)
	var curGuardID int
	var sleptAt int // minutes after midnight
	for _, evt := range events {
		switch evt.kind {
		case beginShift:
			curGuardID = evt.guardID
			if db[curGuardID] == nil {
				db[curGuardID] = new(sleepHistory)
			}
		case fallAsleep:
			sleptAt = evt.when.Minute()
		case wakeUp:
			db[curGuardID].addSpan(sleptAt, evt.when.Minute())
		}
	}
	return db
}

type eventKind int

const (
	beginShift eventKind = iota + 1
	fallAsleep
	wakeUp
)

type event struct {
	when    time.Time
	kind    eventKind
	guardID int
}

var (
	eventRE = regexp.MustCompile(`^\[(\d+-\d\d-\d\d \d\d:\d\d)\] (.*)$`)
	shiftRE = regexp.MustCompile(`Guard #(\d+) begins shift`)
)

func parseEvent(s string) *event {
	m := eventRE.FindStringSubmatch(s)
	if m == nil {
		panic(fmt.Sprintf("failed to parse %q", s))
	}
	ts, err := time.Parse("2006-01-02 15:04", m[1])
	if err != nil {
		panic(fmt.Sprintf("failed to parse %q: %v", s, err))
	}
	evt := event{when: ts}
	switch m[2] {
	case "falls asleep":
		evt.kind = fallAsleep
	case "wakes up":
		evt.kind = wakeUp
	default:
		m2 := shiftRE.FindStringSubmatch(m[2])
		if m2 == nil {
			panic(fmt.Sprintf("failed to parse %q", s))
		}
		evt.kind = beginShift
		evt.guardID = atoi(m2[1])
	}
	return &evt
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("strconv.Atoi(%q) failed: %v", s, err))
	}
	return n
}

type sleepHistory struct {
	spans []span
}

type span struct {
	begin, end int // minutes: [begin, end)
}

func (h *sleepHistory) addSpan(begin, end int) {
	h.spans = append(h.spans, span{begin, end})
}

func (h *sleepHistory) totalSlept() int {
	n := 0
	for _, s := range h.spans {
		n += s.end - s.begin
	}
	return n
}

func (h *sleepHistory) modeFreq() (minute, freq int) {
	hist := make([]int, 60) // histogram
	for _, s := range h.spans {
		for i := s.begin; i < s.end; i++ {
			hist[i]++
		}
	}
	var topMin, topN int
	for min, n := range hist {
		if n > topN {
			topMin, topN = min, n
		}
	}
	return topMin, topN
}
