package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	part := flag.Int("part", 1, "which part")
	flag.Parse()
	switch *part {
	case 1:
		part1()
	case 2:
		part2()
	}
}

func part1() {
	fmt.Println(len(reactAll(readInput())))
}

func part2() {
	original := readInput()
	buf := make([]byte, len(original))
	min := len(original)
	for i := byte(1); i <= 26; i++ {
		copy(buf, original)
		n := len(reactAll(removeBytes(buf, i, i+26)))
		if n < min {
			min = n
		}
	}
	fmt.Println(min)
}

// Drains stdin and applies normalizeLetter to the bytes.
// Trims non-letter bytes from the end.
func readInput() []byte {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	for i := range bytes {
		bytes[i] = normalizeLetter(bytes[i])
	}
	i := len(bytes) - 1
	for ; i >= 0 && bytes[i] == 0; i-- {
	}
	return bytes[:i+1]
}

// Reacts until quiescent. Modifies src in-place. Returns final slice.
func reactAll(src []byte) []byte {
	dst := make([]byte, 0, len(src))
	for {
		dst = react(src, dst)
		if len(dst) == len(src) {
			return dst
		}
		src, dst = dst, src[:0]
	}
}

// Perform a single-pass reaction of neighboring bytes.
func react(src, dst []byte) []byte {
	for i := 0; i < len(src); i++ {
		if i == len(src)-1 {
			dst = append(dst, src[i])
			break
		}
		x, y := src[i], src[i+1]
		if delta(x, y) != 26 {
			dst = append(dst, x)
		} else {
			i++
		}
	}
	return dst
}

// Converts a-z to 1-26 and A-Z to 27-52.
func normalizeLetter(b byte) byte {
	switch {
	case b >= 'a' && b <= 'z':
		return b - 'a' + 1
	case b >= 'A' && b <= 'Z':
		return b - 'A' + 27
	}
	return 0 // invalid
}

// Returns abs(x - y).
func delta(x, y byte) byte {
	if x > y {
		return x - y
	}
	return y - x
}

// Removes the specified bytes from a, in-place.
func removeBytes(a []byte, remove ...byte) []byte {
	j := 0
nextByte:
	for _, b1 := range a {
		for _, b2 := range remove {
			if b1 == b2 {
				continue nextByte
			}
		}
		a[j] = b1
		j++
	}
	return a[:j]
}
