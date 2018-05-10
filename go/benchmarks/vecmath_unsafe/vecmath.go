package vecmath

import "unsafe"

func DotProduct(a, b []float64) float64 {
	var v float64
	for i, x := range a {
		v += x * b[i]
	}
	return v
}

func DotProductUnsafe(a, b []float64) float64 {
	var v float64
	bp := *(*uintptr)(unsafe.Pointer(&b))
	for _, x := range a {
		v += x * *(*float64)(unsafe.Pointer(bp))
		bp += 8
	}
	return v
}
