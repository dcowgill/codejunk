package makemap

type Map map[string][]int

func NoSize(keys []string) {
	m := make(Map)
	for _, k := range keys {
		m[k] = make([]int, 10)
	}
}

func WithSize(keys []string) {
	m := make(Map, len(keys))
	for _, k := range keys {
		m[k] = make([]int, 10)
	}
}
