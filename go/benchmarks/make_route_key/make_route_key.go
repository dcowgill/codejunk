package make_route_key

import (
	"bytes"
	"strconv"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func f1(ids []bson.ObjectId, delay time.Duration) string {
	var key string
	for _, id := range ids {
		key += string(id)
	}
	key += strconv.FormatInt(int64(delay), 10)
	return key
}

func f2(ids []bson.ObjectId, delay time.Duration) string {
	// First calculate the number of bytes in the key.
	n := 0
	for _, id := range ids {
		n += len(id)
	}
	ds := strconv.FormatInt(int64(delay), 10)
	n += len(ds)
	// Allocate a single buffer for the key.
	buf := bytes.NewBuffer(make([]byte, 0, n))
	for _, id := range ids {
		buf.WriteString(string(id))
	}
	buf.WriteString(ds)
	return buf.String()
}

var bufPool = sync.Pool{
	New: func() interface{} { return bytes.NewBuffer(make([]byte, 0, 100)) },
}

func f3(ids []bson.ObjectId, delay time.Duration) string {
	// First calculate the number of bytes in the key.
	n := 0
	for _, id := range ids {
		n += len(id)
	}
	ds := strconv.FormatInt(int64(delay), 10)
	n += len(ds)
	// Allocate a single buffer for the key.
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	for _, id := range ids {
		buf.WriteString(string(id))
	}
	buf.WriteString(ds)
	s := buf.String()
	bufPool.Put(buf)
	return s
}

var oneBuf = new(bytes.Buffer)

func f4(ids []bson.ObjectId, delay time.Duration) string {
	// First calculate the number of bytes in the key.
	n := 0
	for _, id := range ids {
		n += len(id)
	}
	ds := strconv.FormatInt(int64(delay), 10)
	n += len(ds)
	// Allocate a single buffer for the key.
	buf := oneBuf
	buf.Reset()
	for _, id := range ids {
		buf.WriteString(string(id))
	}
	buf.WriteString(ds)
	return buf.String()
}
