package sprintf_test

import (
	"fmt"
	"strconv"
	"testing"
)

func sprintf(col string, i int) string {
	return fmt.Sprintf("%s = $%d", col, i+1)
}

func simple(col string, i int) string {
	return col + " = $" + strconv.Itoa(i+1)
}

func BenchmarkSprintf(b *testing.B) {
	col := "thing_name"
	n := 0
	for i := 0; i < b.N; i++ {
		n += len(sprintf(col, i%20))
	}
}

func BenchmarkSimple(b *testing.B) {
	col := "thing_name"
	n := 0
	for i := 0; i < b.N; i++ {
		n += len(simple(col, i%20))
	}
}
