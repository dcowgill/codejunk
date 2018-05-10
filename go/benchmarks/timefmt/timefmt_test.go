package timefmt

import (
	"testing"
	"time"
)

func TestStuff(t *testing.T) {
	v := time.Now()
	a := Format_YYYY_MM(v)
	b := Printf_YYYY_MM(v)
	if a != b {
		t.Fatalf("%q != %q", a, b)
	}
}

func TestStuff2(t *testing.T) {
	v := time.Now()
	e := v.Format("2006-01-02")
	fns := []func(time.Time) string{
		Format_YYYY_MM_DD,
		Printf_YYYY_MM_DD_v1,
		Printf_YYYY_MM_DD_v2,
	}
	for i, fn := range fns {
		if x := fn(v); x != e {
			t.Fatalf("i=%d x=%q e=%q", i, x, e)
		}
	}
}

func BenchmarkFormat_YYYY_MM(b *testing.B) {
	v := time.Now()
	for i := 0; i < b.N; i++ {
		Format_YYYY_MM(v)
	}
}

func BenchmarkPrintf_YYYY_MM(b *testing.B) {
	v := time.Now()
	for i := 0; i < b.N; i++ {
		Printf_YYYY_MM(v)
	}
}

func BenchmarkFormat_YYYY_MM_DD(b *testing.B) {
	v := time.Now()
	for i := 0; i < b.N; i++ {
		Format_YYYY_MM_DD(v)
	}
}

func BenchmarkPrintf_YYYY_MM_DD_v1(b *testing.B) {
	v := time.Now()
	for i := 0; i < b.N; i++ {
		Printf_YYYY_MM_DD_v1(v)
	}
}

func BenchmarkPrintf_YYYY_MM_DD_v2(b *testing.B) {
	v := time.Now()
	for i := 0; i < b.N; i++ {
		Printf_YYYY_MM_DD_v2(v)
	}
}
