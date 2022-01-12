package d17

import (
	"adventofcode2021/lib"
)

const (
	xmin = 288
	xmax = 330
	ymin = -96
	ymax = -50
)

func Run() {
	lib.Run(17, part1, part2)
}

func part1() int64 {
	sim := func(v int) int {
		y := 0
		h := 0
		for {
			y += v
			h = lib.Max(h, y)
			if v < 0 {
				if y < ymin {
					return 0
				}
				if y <= ymax {
					return h
				}
			}
			v--
		}
		panic("unreachable")
	}
	h := 0
	for v := 0; v < 1000; v++ {
		h = lib.Max(h, sim(v))
	}
	return int64(h)
}

func part2() int64 {
	// Find the x-axis velocities that work.
	var vxs []int
	for vx := 1; vx <= xmax; vx++ {
		if tryX(vx, xmin, xmax) {
			vxs = append(vxs, vx)
		}
	}
	// Find the y-axis velocities that work.
	var vys []int
	for vy := ymin; vy <= 1000; vy++ {
		if tryY(vy, ymin, ymax) {
			vys = append(vys, vy)
		}
	}
	// Test each (vx, vy) pair in the vectors' cross-product.
	n := 0
	for _, vx := range vxs {
		for _, vy := range vys {
			if tryXY(vx, vy, xmin, xmax, ymin, ymax) {
				n++
			}
		}
	}
	return int64(n)
}

// If we launch the probe with velocity vx along the x-axis,
// does it ever end a step inside [xmin, xmax]?
func tryX(vx, xmin, xmax int) bool {
	x := 0
	d := 0 - sign(vx)
	for vx != 0 {
		x += vx
		vx += d
		if x > xmax {
			return false
		}
		if x >= xmin {
			return true
		}
	}
	return false
}

// If we launch the probe with velocity vy along the y-axis,
// does it ever end a step inside [ymin, ymax]?
func tryY(vy, ymin, ymax int) bool {
	y := 0
	for {
		y += vy
		if vy < 0 {
			if y < ymin {
				return false
			}
			if y <= ymax {
				return true
			}
		}
		vy -= 1
	}
	return false
}

// If we launch the probe with velocities vx and vy along the x-axis and
// y-axis, respectively, does it ever a step inside the target area?
func tryXY(vx, vy, xmin, xmax, ymin, ymax int) bool {
	x := 0
	y := 0
	d := 0 - sign(vx)
	for {
		x += vx
		y += vy
		if vy < 0 {
			if y < ymin {
				return false
			}
			if x > xmax {
				return false
			}
			if y <= ymax && x >= xmin {
				return true
			}
		}
		if vx != 0 {
			vx += d
		}
		vy--
	}
	panic("unreachable")
}

func sign(v int) int {
	if v < 0 {
		return -1
	}
	return 1
}
