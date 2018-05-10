package make_route_key

import (
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var ids []bson.ObjectId
var delay time.Duration

func init() {
	for i := 0; i < 5; i++ {
		ids = append(ids, bson.NewObjectId())
	}
	delay = 10*time.Minute + 30*time.Second
}

func BenchmarkNoBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1(ids, delay)
	}
}

func BenchmarkReallocBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f2(ids, delay)
	}
}

func BenchmarkBufferViaSyncPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f3(ids, delay)
	}
}

func BenchmarkSingleBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f4(ids, delay)
	}
}
