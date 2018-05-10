package strcat

import "testing"

func BenchmarkCRandSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CRandSprintf()
	}
}

func BenchmarkCRandConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CRandConcat()
	}
}

func BenchmarkRandSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandSprintf()
	}
}

func BenchmarkRandConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandConcat()
	}
}

func BenchmarkRandBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandBuffer()
	}
}
