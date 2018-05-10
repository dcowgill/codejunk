package parse_rfc3339_test

import (
	"testing"
	"time"
)

func parseRFC3339(s string) time.Time {
	t, _ := time.Parse(s, time.RFC3339)
	return t
}

func millisToTime(ms int64) time.Time {
	s := ms / 1000
	ns := 1000 * 1000 * (ms - s*1000)
	return time.Unix(ms, ns)
}

var (
	now       = time.Now()
	formatted = now.Format(time.RFC3339)
	millis    = now.UnixNano() / 1000 / 1000
)

func BenchmarkParseRFC3339(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseRFC3339(formatted)
	}
}

func BenchmarkMillisToTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		millisToTime(millis)
	}
}
