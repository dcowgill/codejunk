package lib

import (
	"constraints"
	"encoding/json"
)

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

func Foldl[V any](f func(V, V) V, acc V, vs []V) V {
	for _, v := range vs {
		acc = f(acc, v)
	}
	return acc
}

func Foldr[V any](f func(V, V) V, acc V, vs []V) V {
	if len(vs) == 0 {
		return acc
	}
	return f(vs[0], Foldr(f, acc, vs[1:]))
}

func Abs[V constraints.Signed](v V) V {
	if v < 0 {
		return -v
	}
	return v
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
