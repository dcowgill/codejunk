package lib

import (
	"constraints"
	"encoding/json"
)

func Map[X any, Y any](f func(X) Y, xs []X) []Y {
	ys := make([]Y, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	a := make([]K, 0, len(m))
	for k := range m {
		a = append(a, k)
	}
	return a
}

func MapValues[K comparable, V any](m map[K]V) []V {
	a := make([]V, 0, len(m))
	for _, v := range m {
		a = append(a, v)
	}
	return a
}

func Reduce[X any, Y any](f func(Y, X) Y, acc Y, xs []X) Y {
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}

func Foldl[V any](f func(V, V) V, vs []V) V {
	switch len(vs) {
	case 0:
		var zero V
		return zero
	case 1:
		return vs[0]
	}
	acc := vs[0]
	for _, v := range vs[1:] {
		acc = f(acc, v)
	}
	return acc
}

func Foldr[V any](f func(V, V) V, vs []V) V {
	switch len(vs) {
	case 0:
		var zero V
		return zero
	case 1:
		return vs[0]
	}
	return f(vs[0], Foldr(f, vs[1:]))
}

func Abs[V constraints.Signed](v V) V {
	if v < 0 {
		return -v
	}
	return v
}

func Min[V constraints.Ordered](a, b V) V {
	if a < b {
		return a
	}
	return b
}

func Max[V constraints.Ordered](a, b V) V {
	if a > b {
		return a
	}
	return b
}

func Greatest[V constraints.Ordered](vs []V) (answer V) {
	if len(vs) == 0 {
		return
	}
	answer = vs[0]
	for _, v := range vs[1:] {
		if v > answer {
			answer = v
		}
	}
	return
}

func Least[V constraints.Ordered](vs []V) (answer V) {
	if len(vs) == 0 {
		return
	}
	answer = vs[0]
	for _, v := range vs[1:] {
		if v < answer {
			answer = v
		}
	}
	return
}

func Sum[V constraints.Integer](vs []V) V {
	var answer V = 0
	for _, v := range vs {
		answer += v
	}
	return answer
}

func Product[V constraints.Integer](vs []V) V {
	var answer V = 1
	for _, v := range vs {
		answer *= v
	}
	return answer
}

func Make2DArray[V any](rows, cols int) [][]V {
	mem := make([]V, rows*cols)
	mat := make([][]V, rows)
	for i := 0; i < rows; i++ {
		mat[i], mem = mem[:cols], mem[cols:]
	}
	return mat
}

func Dumps(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(data)
}
