package errstack

import (
	"errors"
	"testing"
)

func cause() error { return errors.New("base error") }

func f0_1() error { return wrap1(cause(), "msg0_1") }
func f1_1() error { return wrap1(f0_1(), "msg1_1") }
func f2_1() error { return wrap1(f1_1(), "msg2_1") }
func f3_1() error { return wrap1(f2_1(), "msg3_1") }
func f4_1() error { return wrap1(f3_1(), "msg4_1") }
func f5_1() error { return wrap1(f4_1(), "msg5_1") }
func f6_1() error { return wrap1(f5_1(), "msg6_1") }
func f7_1() error { return wrap1(f6_1(), "msg7_1") }
func f8_1() error { return wrap1(f7_1(), "msg8_1") }
func f9_1() error { return wrap1(f8_1(), "msg9_1") }
func tp_1() error { return wrap1(f9_1(), "msgtop") }

func f0_2() error { return wrap2(cause(), "msg0_2") }
func f1_2() error { return wrap2(f0_2(), "msg1_2") }
func f2_2() error { return wrap2(f1_2(), "msg2_2") }
func f3_2() error { return wrap2(f2_2(), "msg3_2") }
func f4_2() error { return wrap2(f3_2(), "msg4_2") }
func f5_2() error { return wrap2(f4_2(), "msg5_2") }
func f6_2() error { return wrap2(f5_2(), "msg6_2") }
func f7_2() error { return wrap2(f6_2(), "msg7_2") }
func f8_2() error { return wrap2(f7_2(), "msg8_2") }
func f9_2() error { return wrap2(f8_2(), "msg9_2") }
func tp_2() error { return wrap2(f9_2(), "msgtop") }

var numErrs int

func BenchmarkWrap1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := tp_1(); err != nil {
			numErrs++
		}
	}
}

func BenchmarkWrap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := tp_2(); err != nil {
			numErrs++
		}
	}
}

var numCallers int

func BenchmarkCallers1(b *testing.B) {
	err := tp_1()
	for i := 0; i < b.N; i++ {
		numCallers += len(getCallers(err))
	}
}

func BenchmarkCallers2(b *testing.B) {
	err := tp_2()
	for i := 0; i < b.N; i++ {
		numCallers += len(getCallers(err))
	}
}
