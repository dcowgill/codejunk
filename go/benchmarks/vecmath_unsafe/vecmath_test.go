package vecmath

import "testing"

var tests = []struct {
	a, b []float64
	v    float64
}{
	{[]float64{}, []float64{}, 0},
	{[]float64{2, 2, 2}, []float64{4, 5, 6}, 8 + 10 + 12},
	{[]float64{1, 2, 3, 4, 5}, []float64{5, 6, 7, 8, 9}, 115},
	{
		[]float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		[]float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		46,
	},
}

func TestDotProduct(t *testing.T) {
	for _, tt := range tests {
		v := DotProduct(tt.a, tt.b)
		if v != tt.v {
			t.Errorf("DotProduct(%v, %v) returned %v, want %v", tt.a, tt.b, v, tt.v)
		}
	}
}

func TestDotProductUnsafe(t *testing.T) {
	for _, tt := range tests {
		v := DotProductUnsafe(tt.a, tt.b)
		if v != tt.v {
			t.Errorf("DotProduct(%v, %v) returned %v, want %v", tt.a, tt.b, v, tt.v)
		}
	}
}

func BenchmarkDotProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			DotProduct(tt.a, tt.b)
		}
	}
}

func BenchmarkDotProductUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			DotProductUnsafe(tt.a, tt.b)
		}
	}
}
