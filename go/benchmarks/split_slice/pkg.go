package utils

// Returns two slices: the first contains values in xs for which pred returns
// true; the second contains the other values. N.B. modifies xs in place, and
// the two slices returned by the function are both sub-slices of xs.
func splitSlice(xs []string, pred func(string) bool) ([]string, []string) {
	n := len(xs)
	i := 0
	for i < n {
		if pred(xs[i]) {
			i++
			continue
		}
		xs[i], xs[n-1] = xs[n-1], xs[i]
		n--
	}
	return xs[:n], xs[n:]
}

// Like splitSlice but it allocates new slices.
func splitSliceAlloc(xs []string, pred func(string) bool) ([]string, []string) {
	var t, f []string
	for _, x := range xs {
		if pred(x) {
			t = append(t, x)
		} else {
			f = append(f, x)
		}
	}
	return t, f
}
