package reflection

import "testing"

func TestSetFoo(t *testing.T) {
	a := RandFoo()
	b := RandFoo()
	SetFoo(a, b)
	if *a != *b {
		t.Fatal("a != b")
	}
}

func TestSetFooReflect(t *testing.T) {
	a := RandFoo()
	b := RandFoo()
	SetFooReflect(a, b)
	if *a != *b {
		t.Fatal("a != b")
	}
}

func BenchmarkAssignFoo(b *testing.B) {
	x := RandFoo()
	y := RandFoo()
	for i := 0; i < b.N; i++ {
		*x = *y
	}
}

func BenchmarkSetFoo(b *testing.B) {
	x := RandFoo()
	y := RandFoo()
	for i := 0; i < b.N; i++ {
		SetFoo(x, y)
	}
}

func BenchmarkSetFooReflect(b *testing.B) {
	x := RandFoo()
	y := RandFoo()
	for i := 0; i < b.N; i++ {
		SetFooReflect(x, y)
	}
}
