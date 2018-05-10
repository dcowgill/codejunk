package appendbytetostring

import (
	"bytes"
	"testing"
)

func AppendString(s string, b byte) string {
	return s + string(b)
}

func AllocByteSlice(s string, b byte) string {
	bs := make([]byte, len(s)+1)
	copy(bs, s)
	bs[len(bs)-1] = b
	return string(bs)
}

func BytesBuffer(s string, b byte) string {
	buf := bytes.NewBufferString(s)
	buf.WriteByte(b)
	return buf.String()
}

var s = "abcdefgh"
var c = byte('i')

func BenchmarkApendString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendString(s, c)
	}
}

func BenchmarkAllocByteSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllocByteSlice(s, c)
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesBuffer(s, c)
	}
}
