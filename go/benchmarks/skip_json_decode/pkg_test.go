package pkg

import (
	"encoding/json"
	"testing"
)

type Ints10 struct {
	Foo0 int
	Foo1 int
	Foo2 int
	Foo3 int
	Foo4 int
	Foo5 int
	Foo6 int
	Foo7 int
	Foo8 int
	Foo9 int
}

type Strings10 struct {
	Qux0 string
	Qux1 string
	Qux2 string
	Qux3 string
	Qux4 string
	Qux5 string
	Qux6 string
	Qux7 string
	Qux8 string
	Qux9 string
}

type Container struct {
	Foo Ints10
	Bar Ints10
	Baz Strings10
}

type Message struct {
	Container0 *Container
	Container1 *Container
	String0    string
	String1    string
	Int0       int
	Int1       int
}

var value = &Message{
	Container0: &Container{
		Foo: Ints10{
			Foo0: 0,
			Foo1: 1,
			Foo2: 4,
			Foo3: 9,
			Foo4: 16,
			Foo5: 25,
			Foo6: 36,
			Foo7: 49,
			Foo8: 64,
			Foo9: 81,
		},
		Bar: Ints10{
			Foo0: 0,
			Foo1: 2,
			Foo2: 4,
			Foo3: 6,
			Foo4: 8,
			Foo5: 10,
			Foo6: 12,
			Foo7: 14,
			Foo8: 16,
			Foo9: 18,
		},
		Baz: Strings10{
			Qux0: "0 0 0 0 0 0",
			Qux1: "1 1 1 1 1 1",
			Qux2: "2 2 2 2 2 2",
			Qux3: "3 3 3 3 3 3",
			Qux4: "4 4 4 4 4 4",
			Qux5: "5 5 5 5 5 5",
			Qux6: "6 6 6 6 6 6",
			Qux7: "7 7 7 7 7 7",
			Qux8: "8 8 8 8 8 8",
			Qux9: "9 9 9 9 9 9",
		},
	},
	Container1: &Container{
		Foo: Ints10{
			Foo0: 0,
			Foo1: 1,
			Foo2: 4,
			Foo3: 9,
			Foo4: 16,
			Foo5: 25,
			Foo6: 36,
			Foo7: 49,
			Foo8: 64,
			Foo9: 81,
		},
		Bar: Ints10{
			Foo0: 0,
			Foo1: 2,
			Foo2: 4,
			Foo3: 6,
			Foo4: 8,
			Foo5: 10,
			Foo6: 12,
			Foo7: 14,
			Foo8: 16,
			Foo9: 18,
		},
		Baz: Strings10{
			Qux0: "0 0 0 0 0 0",
			Qux1: "1 1 1 1 1 1",
			Qux2: "2 2 2 2 2 2",
			Qux3: "3 3 3 3 3 3",
			Qux4: "4 4 4 4 4 4",
			Qux5: "5 5 5 5 5 5",
			Qux6: "6 6 6 6 6 6",
			Qux7: "7 7 7 7 7 7",
			Qux8: "8 8 8 8 8 8",
			Qux9: "9 9 9 9 9 9",
		},
	},
	String0: "Hello, world!",
	String1: "Goodbye, cruel world!",
	Int0:    42,
	Int1:    1 << 30,
}

type skipDecode struct{} // Stops JSON decoding.

func (v *skipDecode) UnmarshalJSON(data []byte) error {
	return nil
}

func BenchmarkDecode(b *testing.B) {
	data, _ := json.Marshal(value)
	b.Run("FullMessage", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var msg Message
			if err := json.Unmarshal(data, &msg); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("RawMessage", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var msg map[string]json.RawMessage
			if err := json.Unmarshal(data, &msg); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("SkipDecode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var msg map[string]skipDecode
			if err := json.Unmarshal(data, &msg); err != nil {
				b.Fatal(err)
			}
		}
	})
}
