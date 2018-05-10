package coords_map_test

import (
	"sort"
	"testing"

	"github.com/omakasecorp/samurai/pkg/earth"
	"github.com/omakasecorp/samurai/pkg/math2"
	"github.com/omakasecorp/samurai/pkg/util"
)

var (
	k1 []int
	m1 map[int]int
	k2 []earth.Coordinates
	m2 map[earth.Coordinates]int

	mat [][]math2.Meters
)

type coords []earth.Coordinates

func (a coords) Len() int      { return len(a) }
func (a coords) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a coords) Less(i, j int) bool {
	if a[i].Lat < a[j].Lat {
		return true
	} else if a[i].Lat > a[j].Lat {
		return false
	}
	return a[i].Lon < a[j].Lon
}

func init() {
	const N = 1000

	k1 = make([]int, N)
	for i := 0; i < N; i++ {
		k1[i] = i
	}

	k2 = make([]earth.Coordinates, N)
	for i := 0; i < N; i++ {
		k2[i] = earth.DegCoords(
			math2.Degrees(+40+float64(i)*0.001),
			math2.Degrees(-73+float64(i)*0.001),
		)
	}

	util.Shuffle(sort.IntSlice(k1))
	util.Shuffle(coords(k2))

	m1 = make(map[int]int)
	for i, k := range k1 {
		m1[k] = i * i
	}

	m2 = make(map[earth.Coordinates]int)
	for i, k := range k2 {
		m2[k] = i * i
	}

	const M = 100
	mat = newDistanceMatrix(M)
	for i := 0; i < M; i++ {
		for j := 0; j < M; j++ {
			if i != j {
				mat[i][j] = k2[i].DistanceTo(k2[j])
			}
		}
	}
}

func BenchmarkIntMapLookup(b *testing.B) {
	k := 274
	for i := 0; i < b.N; i++ {
		_ = m1[k]
	}
}

func BenchmarkCoordsMapLookup(b *testing.B) {
	k := k2[274]
	for i := 0; i < b.N; i++ {
		_ = m2[k]
	}
}

func BenchmarkDistanceTo(b *testing.B) {
	p := k2[72]
	q := k2[98]
	for i := 0; i < b.N; i++ {
		p.DistanceTo(q)
	}
}

func BenchmarkMatrixLookup(b *testing.B) {
	p := 72
	q := 98
	for i := 0; i < b.N; i++ {
		_ = mat[p][q]
	}
}

// Returns two-dimensional array of math2.Meters values.
func newDistanceMatrix(n int) [][]math2.Meters {
	mem := make([]math2.Meters, n*n)
	mat := make([][]math2.Meters, n)
	for i := 0; i < n; i++ {
		mat[i], mem = mem[:n], mem[n:]
	}
	return mat
}
