package switch_vs_map

import (
	"reflect"
	"testing"
)

var kindLookup []bool

func init() {
	kinds := []reflect.Kind{
		reflect.Bool, reflect.Int, reflect.Int8,
		reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.String,
	}
	max := 0
	for _, k := range kinds {
		if int(k) > max {
			max = int(k)
		}
	}
	kindLookup = make([]bool, max+1)
	for _, k := range kinds {
		kindLookup[int(k)] = true
	}
}

func isScalarTypeSwitch(k reflect.Kind) bool {
	switch k {
	case reflect.Bool, reflect.Int, reflect.Int8,
		reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.String:
		return true
	}
	return false
}

func isScalarTypeLookup(k reflect.Kind) bool {
	i := int(k)
	if i >= len(kindLookup) {
		return false
	}
	return kindLookup[i]
}

var tests = []reflect.Kind{
	reflect.Int,
	reflect.Float64,
	reflect.Uint64,
	reflect.Int32,
	reflect.Complex64,
	reflect.Int64,
	reflect.Uint8,
	reflect.Ptr,
	reflect.Uint32,
	reflect.Int8,
	reflect.Float64,
	reflect.Float64,
	reflect.Uintptr,
	reflect.Interface,
	reflect.Slice,
	reflect.Float32,
	reflect.Uint32,
	reflect.Slice,
	reflect.Chan,
	reflect.Uint64,
	reflect.Uint32,
	reflect.Int,
	reflect.Slice,
	reflect.Int32,
	reflect.Complex64,
}

func BenchmarkSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, k := range tests {
			isScalarTypeSwitch(k)
		}
	}
}

func BenchmarkLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, k := range tests {
			isScalarTypeLookup(k)
		}
	}
}
