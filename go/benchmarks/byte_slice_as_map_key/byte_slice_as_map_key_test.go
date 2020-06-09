package byte_slice_as_map_key

import (
	"math/rand"
	"testing"
)

func BenchmarkMapLookup(b *testing.B) {
	const (
		numStrings = 100
		numBytes   = 100
	)
	var (
		byteSlices [][]byte
		strings    []string
	)
	for i := 0; i < numStrings; i++ {
		a := make([]byte, numBytes)
		for j := range a {
			a[j] = 'a' + byte(rand.Intn(26))

		}
		byteSlices = append(byteSlices, a)
		strings = append(strings, string(a))
	}

	b.Run("byteslice", func(b *testing.B) {
		m := make(map[string]struct{})
		for i := 0; i < b.N; i++ {
			for _, k := range byteSlices {
				m[string(k)] = struct{}{}
			}
		}
		// fmt.Printf("DEBUG: len(m) = %d\n", len(m))
	})

	b.Run("string", func(b *testing.B) {
		m := make(map[string]struct{})
		for i := 0; i < b.N; i++ {
			for _, k := range strings {
				m[k] = struct{}{}
			}
		}
		// fmt.Printf("DEBUG: len(m) = %d\n", len(m))
	})
}
