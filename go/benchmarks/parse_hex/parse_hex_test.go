package pkg

import (
	"math/rand"
	"testing"
)

var fns = map[string]func(string) int64{
	"parseStdlib":        parseStdlib,
	"parseHexCustom":     parseHexCustom,
	"parseHexByteLookup": parseHexByteLookup,
}

func TestParseHex(t *testing.T) {
	const numExamples = 100
	examples := make([]string, numExamples)
	for i := range examples {
		examples[i] = randHexStr()
	}
	for name, fn := range fns {
		if name == "parseStdlib" {
			continue
		}
		t.Run(name, func(t *testing.T) {
			for _, s := range examples {
				got, want := fn(s), parseStdlib(s)
				if got != want {
					t.Fatalf("fn(%q) failed: got %d, want %d", s, got, want)
				}
			}
		})
	}
}

func BenchmarkParseHex(b *testing.B) {
	const numExamples = 100
	examples := make([]string, numExamples)
	for i := range examples {
		examples[i] = randHexStr()
	}
	for name, fn := range fns {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, s := range examples {
					fn(s)
				}
			}
		})
	}
}

func randHexStr() string {
	n := 2 + rand.Intn(11) // length in [2, 12]
	b := make([]byte, n)
	for i := range b {
		x := byte(rand.Intn(16)) // [0,15]
		if x < 10 {
			b[i] = '0' + x
		} else {
			b[i] = 'a' + x - 10
		}
	}
	return string(b)
}
