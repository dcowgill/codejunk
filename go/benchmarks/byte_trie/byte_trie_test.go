package byte_trie

import (
	"math/rand"
	"testing"
)

func BenchmarkByteTrieVsMap(b *testing.B) {
	// Setup.
	const (
		numSlices = 100
		sliceLen  = 100
	)
	data := make([][]byte, 100)
	for i := range data {
		v := make([]byte, sliceLen)
		for i := range v {
			v[i] = byte(rand.Intn(3)) // values must be in [0, 2]
		}
		data[i] = v
	}

	b.Run("byte trie", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var t trie
			found := 0
			for _, v := range data {
				if t.contains(v) {
					found++
				}
			}
			if found != 0 {
				b.Fatalf("found %d, expected 0", found)
			}
			// for _, v := range data {
			// 	t.insert(v)
			// }
			// found = 0
			// for _, v := range data {
			// 	if t.contains(v) {
			// 		found++
			// 	}
			// }
			// if found != len(data) {
			// 	b.Fatalf("found %d, expected %d", found, len(data))
			// }
		}
	})

	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make(map[string]bool)
			found := 0
			for _, v := range data {
				if m[string(v)] {
					found++
				}
			}
			if found != 0 {
				b.Fatalf("found %d, expected 0", found)
			}
			// for _, v := range data {
			// 	m[string(v)] = true
			// }
			// found = 0
			// for _, v := range data {
			// 	if m[string(v)] {
			// 		found++
			// 	}
			// }
			// if found != len(data) {
			// 	b.Fatalf("found %d, expected %d", found, len(data))
			// }
		}
	})
}
