package dcp043

import (
	"testing"
)

func TestStack(t *testing.T) {
	const N = 10
	var s stack

	// Push 1..N; max should be the pushed value.
	for i := 1; i <= N; i++ {
		s.push(i)
		if x := s.max(); x != i {
			t.Fatalf("s.max() returned %d, want %d\n", x, i)
		}
		if x := s.size(); x != i {
			t.Fatalf("s.size() returned %d, want %d\n", x, i)
		}
	}

	// Push 1..N; max remain N.
	for i := 1; i <= N; i++ {
		s.push(i)
		if x := s.max(); x != N {
			t.Fatalf("s.max() returned %d, want %d\n", x, N)
		}
		if x := s.size(); x != i+N {
			t.Fatalf("s.size() returned %d, want %d\n", x, i)
		}
	}

	// Pop N values, which should be N..i; max should remain N.
	for i := N; i > 0; i-- {
		if x := s.max(); x != N {
			t.Fatalf("s.max() returned %d, want %d\n", x, N)
		}
		if x := s.pop(); x != i {
			t.Fatalf("s.pop() returned %d, want %d\n", x, i)
		}
	}

	// Pop N values, which should be N..i; max should also be N..i.
	for i := N; i > 0; i-- {
		if x := s.max(); x != i {
			t.Fatalf("s.max() returned %d, want %d\n", x, i)
		}
		if x := s.pop(); x != i {
			t.Fatalf("s.pop() returned %d, want %d\n", x, i)
		}
	}

	// Ensure the stack is empty.
	if s.size() != 0 {
		t.Fatalf("expected an empty stack")
	}
}
