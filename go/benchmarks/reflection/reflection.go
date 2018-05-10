package reflection

import (
	"math/rand"
	"reflect"
	"strconv"
)

type Foo struct {
	F01 string
	F02 int
	F03 float64
	F04 string
	F05 int
	F06 float64
	F07 string
	F08 int
	F09 float64
	F10 string
	F11 int
	F12 float64
	F13 string
	F14 int
	F15 float64
}

func SetFoo(src, dst *Foo) {
	*src = *dst
}

func SetFooReflect(src *Foo, dst interface{}) {
	reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(src).Elem())
}

func RandFoo() *Foo {
	return &Foo{
		F01: randString(),
		F02: rand.Int(),
		F03: rand.Float64(),
		F04: randString(),
		F05: rand.Int(),
		F06: rand.Float64(),
		F07: randString(),
		F08: rand.Int(),
		F09: rand.Float64(),
		F10: randString(),
		F11: rand.Int(),
		F12: rand.Float64(),
		F13: randString(),
		F14: rand.Int(),
		F15: rand.Float64(),
	}
}

func randString() string {
	s := ""
	for i := 0; i < 10; i++ {
		s += strconv.Itoa(rand.Int())
	}
	return s
}
