package load_location

import "testing"

func BenchmarkLoadLocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := LoadLocation()
		if v == nil {
			b.Fatal("nil")
		}
	}
}

func BenchmarkLoadLocationCached(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := LoadLocationCached()
		if v == nil {
			b.Fatal("nil")
		}
	}
}
