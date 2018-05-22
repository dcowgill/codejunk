package timenow

import (
	"testing"
	"time"
)

func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := time.Now()
		time.Since(t)
	}
}
