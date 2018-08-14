package max_slice

func f1(xs []int) int {
	if len(xs) == 0 {
		return 0
	}
	x := xs[0]
	for _, y := range xs[1:] {
		if y > x {
			x = y
		}
	}
	return x
}

func f2(xs []int) int {
	if len(xs) == 0 {
		return 0
	}
	x := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] > x {
			x = xs[i]
		}
	}
	return x
}

func f3(xs []int) int {
	if len(xs) == 0 {
		return 0
	}
	x := xs[0]
	for i := 1; i < len(xs); i++ {
		x = max(x, xs[i])
	}
	return x
}

func f4(xs []int) int {
	if len(xs) == 0 {
		return 0
	}
	x := xs[0]
	for _, y := range xs[1:] {
		x = max(x, y)
	}
	return x
}

func f5(xs []int) int {
	return reduce(max, xs)
}

func reduce(f func(acc, x int) int, xs []int) int {
	acc := xs[0]
	for _, x := range xs[1:] {
		acc = f(acc, x)
	}
	return acc
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
